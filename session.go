// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package zk

import (
	"bytes"
	"fmt"
	"github.com/streadway/zk/proto" // don't export types from here
	"io"
	"math/rand"
	"net"
	"time"
)

// CreateType specifies the type of the node to Create.
type CreateType int32

const (
	CreatePersistent CreateType = 0 // Node remains after session expires
	CreateEphemeral  CreateType = 1 // Node deletes when the creator session expires
	CreateSequence   CreateType = 2 // Sequence number is appended to the node path
)

const (
	// Default session timeout
	Timeout = 4 * time.Second

	// Default address when Config.Addrs is empty
	Address = "127.0.0.1:2181"
)

type Stat struct {
	proto.Stat // FIXME(streadway) don't expose proto package
}

// packet encapsulates a recv that can be processed as an event or rpc reply
type packet struct {
	*proto.ReplyHeader
	proto.Input
}

// Config captures the connection and session state for dialing or re-dialing a
// Zookeeper server.
type Config struct {
	// Host:Port pairs of a subset of servers in the ensemble
	Addrs []string

	// net.Conn factory - default is DialRandom
	Dial func(conf *Config) (net.Conn, error)

	Id       int64  // Id is the session identity useful for server correlation
	Password []byte // Password authenticates this connection to the session id

	Zxid    int64         // Zxid is the last transaction the client has seen
	Timeout time.Duration // Duration until a disconnect expires

	DataWatches   Watches
	ChildWatches  Watches
	ExistsWatches Watches
}

type Session struct {
	*Config

	conn net.Conn

	recvs chan *packet
	sends chan *txn

	events chan Event

	expire chan bool
	closed chan bool

	err error
}

// DialRandom dials a normally distributed random selection from the Addrs with
// a connection timeout such that every address could be attempted before the
// session is expired.
func DialRandom(conf *Config) (net.Conn, error) {
	return net.DialTimeout("tcp",
		conf.Addrs[rand.Intn(len(conf.Addrs))],
		conf.Timeout/time.Duration(len(conf.Addrs)),
	)
}

// Dial connects then either transfers or establishes a new session to a
// Zookeeper ensemble.
func Dial(conf *Config) (*Session, error) {
	var err error

	zk := &Session{
		Config: conf,

		recvs: make(chan *packet),
		sends: make(chan *txn),

		events: make(chan Event),

		// buffered so we can initiate a shutdown only once
		expire: make(chan bool, 1),
		closed: make(chan bool),
	}

	if zk.Config.Timeout == 0 {
		zk.Config.Timeout = Timeout
	}

	if zk.Dial == nil {
		zk.Dial = DialRandom
	}

	if len(conf.Addrs) == 0 {
		zk.Addrs = []string{Address}
	}

	zk.conn, err = zk.Dial(zk.Config)
	if err != nil {
		return nil, err
	}

	if err = zk.connect(); err != nil {
		return nil, err
	}

	errs := make(chan error, 1)
	go zk.receiver(errs)
	go zk.notifier()
	go zk.muxer(errs)

	// Restore session state
	if err := zk.setWatches(); err != nil {
		return nil, err
	}

	return zk, nil
}

// special xids - consider moving to proto
const (
	xidWatch      = -1
	xidPing       = -2
	xidAuth       = -4
	xidSetWatches = -8
)

// protocol operation - consider moving to proto
type op int32

const (
	//opNotify       op = 0
	opCreate  op = 1
	opDelete  op = 2
	opExists  op = 3
	opGetData op = 4
	opSetData op = 5
	opGetAcl  op = 6
	opSetAcl  op = 7
	//opGetChildren op = 8 // not used
	opSync         op = 9
	opPing         op = 11
	opGetChildren2 op = 12
	// FIXME(streadway) implement multi transactions
	//opCheck        op = 13
	//opMulti        op = 14
	opClose op = -11
	// FIXME(streadway) implement auth on setup
	//opSetAuth      op = 100
	opSetWatches op = 101
	//opSasl   op = 102 // maybe implemented
)

type txn struct {
	op  op
	xid int32
	req proto.Encodable
	rep proto.Decodable

	err error
	fin chan bool
	ack chan bool
}

func (tx txn) encode(w *Session) error {
	head := &proto.RequestHeader{
		Type: int32(tx.op),
		Xid:  int32(tx.xid),
	}
	if tx.req == nil {
		return w.send(head)
	}
	return w.send(head, tx.req)
}

func (tx txn) decode(in proto.Input) error {
	if tx.rep != nil {
		return tx.rep.Decode(in)
	}
	return nil
}

// pops and merges the set of watches for an event type
func (zk *Session) watches(watch Event) []chan<- Event {
	switch watch.Type {
	default:
		panic(fmt.Errorf("unhandled watch type: %d", watch.Type))

	case Created, Changed:
		return append(
			zk.ExistsWatches.pop(watch.Path),
			zk.DataWatches.pop(watch.Path)...,
		)

	case Deleted:
		return append(append(
			zk.ExistsWatches.pop(watch.Path), // maybe when create is not observed while disconnected
			zk.DataWatches.pop(watch.Path)...),
			zk.ChildWatches.pop(watch.Path)...,
		)

	case Child:
		return zk.ChildWatches.pop(watch.Path)
	}

	return nil
}

func (zk Session) readFrame() (proto.Input, error) {
	in := proto.BinaryReader{zk.conn}
	len, err := in.ReadInt32()
	if len < 0 {
		return nil, ErrProtocol
	}
	if err != nil {
		return nil, ioError(err)
	}

	buf := make([]byte, len)
	if _, err = io.ReadFull(zk.conn, buf); err != nil {
		return nil, ioError(err)
	}

	return proto.BinaryReader{bytes.NewBuffer(buf)}, nil
}

func (zk Session) readOne(msg proto.Decodable) (proto.Input, error) {
	in, err := zk.readFrame()
	if err != nil {
		return nil, err
	}

	if err := msg.Decode(in); err != nil {
		return nil, ErrProtocol
	}

	return in, nil
}

func (zk Session) readReply() (*packet, error) {
	rep := &proto.ReplyHeader{}
	in, err := zk.readOne(rep)
	if err != nil {
		return nil, err
	}

	return &packet{rep, in}, nil
}

// reply reader
func (zk Session) receiver(errs chan error) {
	for {
		pkt, err := zk.readReply()
		if err != nil {
			errs <- err
			return
		}
		zk.recvs <- pkt
	}
}

// buffered total ordering of all watch events
func (zk Session) notifier() {
	var events []Event // ref to grown array
	var buffer []Event // index into array

	for {
		for len(buffer) > 0 {
			cur := buffer[0]
			for _, watch := range zk.watches(cur) {
				select {
				case next := <-zk.events:
					buffer = append(buffer, next)
				case <-zk.closed:
					// undelivered watches
					return
				case watch <- cur:
					// delivered
				}
			}
			buffer = buffer[1:]
		}

		select {
		case ev := <-zk.events:
			buffer = append(events, ev)

		case <-zk.closed:
			return
		}
	}
}

func (zk Session) muxer(errs chan error) {
	var (
		err     error
		xid     int32
		pending = make(map[int32]*txn)
		health  = time.NewTicker(zk.Timeout / 3)
	)

	defer health.Stop()

	for {
		select {
		case now := <-health.C:
			zk.send(&proto.RequestHeader{Type: int32(opPing), Xid: int32(xidPing)})
			zk.conn.SetDeadline(now.Add(zk.Timeout))

		case req := <-zk.sends:
			xid++
			if req.xid == 0 {
				req.xid = xid
			}
			if pending[req.xid] != nil {
				panic("duplicate send: please report a bug with a test case")
			}
			pending[req.xid] = req
			if err = req.encode(&zk); err != nil {
				err = ErrConnection
				goto error
			}

		case pkt := <-zk.recvs:
			switch pkt.Xid {
			case xidPing:
				zk.conn.SetDeadline(time.Now().Add(zk.Timeout))

			case xidWatch:
				ev := &proto.WatcherEvent{}
				if err = ev.Decode(pkt); err != nil {
					err = ErrProtocol
					goto error
				}
				zk.events <- Event{
					Type: EventType(ev.Type),
					Path: ev.Path,
				}

			default:
				if pkt.Zxid < zk.Config.Zxid {
					panic("out of order delivery")
				}
				zk.Config.Zxid = pkt.Zxid

				tx := pending[pkt.Xid]
				if tx == nil {
					err = ErrProtocol
					goto error
				}
				delete(pending, pkt.Xid)

				if pkt.Err != 0 {
					tx.err = errCode(pkt.Err).toError()
					tx.fin <- true
					<-tx.ack
					continue
				}

				if err = tx.decode(pkt); err != nil {
					err = ErrProtocol
					tx.err = err
					tx.fin <- true
					<-tx.ack
					goto error
				}

				tx.fin <- true
				<-tx.ack
			}

		case err = <-errs:
			goto error

		case <-zk.expire:
			err = ErrExpired
			goto error
		}
	}

error:
	// cleanup
	zk.err = err
	for _, tx := range pending {
		tx.err = err
		tx.fin <- true
		<-tx.ack
	}
	// reader should close the conn
	close(zk.closed)
}

// connect makes a synchronous request without packet headers to establish or
// re-establish a zookeeper session.
func (zk *Session) connect() error {
	// TODO the server expects a lone bool at the end of the request in this
	// packet to indicate the session should stay connected under a server
	// partition.  Normally, read/write sessions are disconnected during a
	// server partition.

	// ConnectRequest doesn't use a packet, doesn't have an Xid or opCode and
	// uses hijacks previously established sessions.
	req := &proto.ConnectRequest{
		LastZxidSeen: zk.Config.Zxid,
		SessionId:    zk.Config.Id,
		TimeOut:      int32(zk.Config.Timeout / time.Millisecond),
		Passwd:       zk.Config.Password,
	}

	if err := zk.send(req); err != nil {
		return err
	}

	res := &proto.ConnectResponse{}
	if _, err := zk.readOne(res); err != nil {
		return err
	}

	// Same signal from java client that the session wasn't established.
	if res.TimeOut <= 0 {
		return ErrExpired
	}

	zk.Config.Id = res.SessionId
	zk.Config.Timeout = time.Duration(res.TimeOut) * time.Millisecond
	zk.Config.Password = res.Passwd

	return nil
}

// setWatches will initiate watch events on the server if the latest state of
// the node at the path is after this session's zxid.
func (zk *Session) setWatches() error {
	req := &proto.SetWatches{
		RelativeZxid: zk.Zxid,
		DataWatches:  zk.DataWatches.paths(),
		ExistWatches: zk.ExistsWatches.paths(),
		ChildWatches: zk.ChildWatches.paths(),
	}

	if len(req.DataWatches) == 0 &&
		len(req.ExistWatches) == 0 &&
		len(req.ChildWatches) == 0 {
		return nil
	}

	done, err := zk.rpc(&txn{op: opSetWatches, xid: xidSetWatches, req: req})
	done <- true
	return err
}

func (zk *Session) send(messages ...proto.Encodable) error {
	buf := new(bytes.Buffer)
	pkt := proto.BinaryWriter{buf}

	for _, m := range messages {
		if err := m.Encode(pkt); err != nil {
			return err
		}
	}

	return proto.BinaryWriter{zk.conn}.WriteBuffer(buf.Bytes())
}

// rpc sends a request header, request, then handles possible responses.  If
// the reply header returns an error, that will be translated to an error that
// can be compared like ErrNoNode.
func (zk *Session) rpc(tx *txn) (chan bool, error) {
	tx.fin = make(chan bool, 1)
	tx.ack = make(chan bool, 1)

	select {
	case <-zk.closed:
		return tx.ack, zk.err
	case zk.sends <- tx:
		<-tx.fin
	}
	return tx.ack, tx.err
}

// Close tells the server it should expire this session and closes the
// connection.
func (zk *Session) Close() error {
	done, err := zk.rpc(&txn{op: opClose})
	done <- true

	// either our err is expired, or we initiate a muxer shutdown
	select {
	case zk.expire <- true:
	default:
	}

	return err
}

// Delete removes the node if the version matches.
func (zk *Session) Delete(path string, version int32) error {
	req := &proto.DeleteRequest{Path: path, Version: version}
	done, err := zk.rpc(&txn{op: opDelete, req: req})
	done <- true
	return err
}

// Create makes a node if it doesn't already exist, setting the data and ACL.
//
// Create returns the path created which will differ from the path requested for Sequence nodes.
//
// Create returns ErrExists if the node already exists.
func (zk *Session) Create(path string, data []byte, types CreateType, acl ACL) (string, error) {
	req := &proto.CreateRequest{
		Path:  path,
		Data:  data,
		Acl:   toProtoACLs(acl),
		Flags: int32(types),
	}
	rep := &proto.CreateResponse{}
	done, err := zk.rpc(&txn{op: opCreate, req: req, rep: rep})
	done <- true
	return rep.Path, err
}

// Get returns the data and node statistics for the path.  If watch is
// provided, it will be fired once when the data changes at this path.
//
// Get returns ErrNoNode and does not set a watch when the node does not exist.
func (zk *Session) Get(path string, watch chan<- Event) ([]byte, Stat, error) {
	req := &proto.GetDataRequest{
		Path:  path,
		Watch: (watch != nil),
	}
	rep := &proto.GetDataResponse{}

	done, err := zk.rpc(&txn{op: opGetData, req: req, rep: rep})
	defer func() { done <- true }()

	if err != nil {
		return nil, Stat{}, err
	}

	if watch != nil {
		zk.DataWatches.add(path, watch)
	}

	return rep.Data, Stat{rep.Stat}, nil
}

// Set changes the data and updates the stat attributes at the provided path.
//
// Set returns ErrNoNode and does not set a watch when the node does not exist.
func (zk *Session) Set(path string, data []byte, version int32) (Stat, error) {
	req := &proto.SetDataRequest{
		Path:    path,
		Data:    data,
		Version: version,
	}
	rep := &proto.SetDataResponse{}

	done, err := zk.rpc(&txn{op: opSetData, req: req, rep: rep})
	done <- true

	if err != nil {
		return Stat{}, err
	}

	return Stat{rep.Stat}, nil
}

// Children returns the paths of the children and node statistics for the path
// given.  If watch is provided, it will be fired once when the number of
// children at this path changes.
//
// Children returns ErrNoNode and does not set a watch when the node does not
// exist.
func (zk *Session) Children(path string, watch chan<- Event) ([]string, Stat, error) {
	req := &proto.GetChildren2Request{Path: path, Watch: (watch != nil)}
	rep := &proto.GetChildren2Response{}

	done, err := zk.rpc(&txn{op: opGetChildren2, req: req, rep: rep})
	defer func() { done <- true }()

	if err != nil {
		return nil, Stat{}, err
	}

	if watch != nil {
		zk.ChildWatches.add(path, watch)
	}

	return rep.Children, Stat{rep.Stat}, nil
}

// Sync reaches consensus and commits the data of the node at the path to disk.
func (zk *Session) Sync(path string) (string, error) {
	req := &proto.SyncRequest{Path: path}
	rep := &proto.SyncResponse{}

	done, err := zk.rpc(&txn{op: opSync, req: req, rep: rep})
	done <- true
	return rep.Path, err
}

// Exists returns the stat of a node if it exists and optionally sets a watch
// that will fire when node changes.
func (zk *Session) Exists(path string, watch chan<- Event) (bool, Stat, error) {
	req := &proto.ExistsRequest{Path: path, Watch: (watch != nil)}
	rep := &proto.ExistsResponse{}

	done, err := zk.rpc(&txn{op: opExists, req: req, rep: rep})
	defer func() { done <- true }()

	if watch != nil {
		zk.ExistsWatches.add(path, watch)
	}

	switch err {
	default:
		return false, Stat{}, err
	case nil:
		return true, Stat{rep.Stat}, nil
	case ErrNoNode:
		return false, Stat{}, nil
	}
}

// GetAcl returns the last seen access control list and the stat for the given path.
func (zk *Session) GetAcl(path string) (ACL, Stat, error) {
	req := &proto.GetACLRequest{Path: path}
	rep := &proto.GetACLResponse{}

	done, err := zk.rpc(&txn{op: opGetAcl, req: req, rep: rep})
	done <- true

	return fromProtoACLs(rep.Acl), Stat{rep.Stat}, err
}

// SetAcl replaces the access control list for the given path and version.
func (zk *Session) SetAcl(path string, acl ACL, version int32) (Stat, error) {
	req := &proto.SetACLRequest{
		Path:    path,
		Acl:     toProtoACLs(acl),
		Version: version,
	}
	rep := &proto.SetACLResponse{}

	done, err := zk.rpc(&txn{op: opSetAcl, req: req, rep: rep})
	done <- true

	return Stat{rep.Stat}, err
}
