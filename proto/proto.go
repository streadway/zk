// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

// GENERATED - DO NOT EDIT

package proto

import "fmt"

type Id struct {
	Scheme string
	Id     string
}

func (m *Id) String() string { return fmt.Sprintf("%+v", *m) }
func (m *Id) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Scheme, err = in.ReadString(); err != nil {
		return err
	}
	if m.Id, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *Id) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Scheme); err != nil {
		return err
	}
	if err = o.WriteString(m.Id); err != nil {
		return err
	}
	return nil
}

type ACL struct {
	Perms int32
	Id    Id
}

func (m *ACL) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ACL) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Perms, err = in.ReadInt32(); err != nil {
		return err
	}
	if err = (&(m.Id)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *ACL) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Perms); err != nil {
		return err
	}
	if err = m.Id.Encode(o); err != nil {
		return err
	}
	return nil
}

type Stat struct {
	Czxid          int64
	Mzxid          int64
	Ctime          int64
	Mtime          int64
	Version        int32
	Cversion       int32
	Aversion       int32
	EphemeralOwner int64
	DataLength     int32
	NumChildren    int32
	Pzxid          int64
}

func (m *Stat) String() string { return fmt.Sprintf("%+v", *m) }
func (m *Stat) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Czxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Mzxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Ctime, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Mtime, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Cversion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Aversion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.EphemeralOwner, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.DataLength, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.NumChildren, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Pzxid, err = in.ReadInt64(); err != nil {
		return err
	}
	return nil
}
func (m *Stat) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt64(m.Czxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Mzxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Ctime); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Mtime); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Cversion); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Aversion); err != nil {
		return err
	}
	if err = o.WriteInt64(m.EphemeralOwner); err != nil {
		return err
	}
	if err = o.WriteInt32(m.DataLength); err != nil {
		return err
	}
	if err = o.WriteInt32(m.NumChildren); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Pzxid); err != nil {
		return err
	}
	return nil
}

type StatPersisted struct {
	Czxid          int64
	Mzxid          int64
	Ctime          int64
	Mtime          int64
	Version        int32
	Cversion       int32
	Aversion       int32
	EphemeralOwner int64
	Pzxid          int64
}

func (m *StatPersisted) String() string { return fmt.Sprintf("%+v", *m) }
func (m *StatPersisted) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Czxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Mzxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Ctime, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Mtime, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Cversion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Aversion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.EphemeralOwner, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Pzxid, err = in.ReadInt64(); err != nil {
		return err
	}
	return nil
}
func (m *StatPersisted) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt64(m.Czxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Mzxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Ctime); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Mtime); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Cversion); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Aversion); err != nil {
		return err
	}
	if err = o.WriteInt64(m.EphemeralOwner); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Pzxid); err != nil {
		return err
	}
	return nil
}

type StatPersistedV1 struct {
	Czxid          int64
	Mzxid          int64
	Ctime          int64
	Mtime          int64
	Version        int32
	Cversion       int32
	Aversion       int32
	EphemeralOwner int64
}

func (m *StatPersistedV1) String() string { return fmt.Sprintf("%+v", *m) }
func (m *StatPersistedV1) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Czxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Mzxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Ctime, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Mtime, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Cversion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Aversion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.EphemeralOwner, err = in.ReadInt64(); err != nil {
		return err
	}
	return nil
}
func (m *StatPersistedV1) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt64(m.Czxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Mzxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Ctime); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Mtime); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Cversion); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Aversion); err != nil {
		return err
	}
	if err = o.WriteInt64(m.EphemeralOwner); err != nil {
		return err
	}
	return nil
}

type ConnectRequest struct {
	ProtocolVersion int32
	LastZxidSeen    int64
	TimeOut         int32
	SessionId       int64
	Passwd          []byte
}

func (m *ConnectRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ConnectRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.ProtocolVersion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.LastZxidSeen, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.TimeOut, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.SessionId, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Passwd, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *ConnectRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.ProtocolVersion); err != nil {
		return err
	}
	if err = o.WriteInt64(m.LastZxidSeen); err != nil {
		return err
	}
	if err = o.WriteInt32(m.TimeOut); err != nil {
		return err
	}
	if err = o.WriteInt64(m.SessionId); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Passwd); err != nil {
		return err
	}
	return nil
}

type ConnectResponse struct {
	ProtocolVersion int32
	TimeOut         int32
	SessionId       int64
	Passwd          []byte
}

func (m *ConnectResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ConnectResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.ProtocolVersion, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.TimeOut, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.SessionId, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Passwd, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *ConnectResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.ProtocolVersion); err != nil {
		return err
	}
	if err = o.WriteInt32(m.TimeOut); err != nil {
		return err
	}
	if err = o.WriteInt64(m.SessionId); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Passwd); err != nil {
		return err
	}
	return nil
}

type SetWatches struct {
	RelativeZxid int64
	DataWatches  []string
	ExistWatches []string
	ChildWatches []string
}

func (m *SetWatches) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetWatches) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.RelativeZxid, err = in.ReadInt64(); err != nil {
		return err
	}
	lenDataWatches, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.DataWatches = make([]string, lenDataWatches)
	for i := 0; i < len(m.DataWatches); i++ {
		if m.DataWatches[i], err = in.ReadString(); err != nil {
			return err
		}
	}
	lenExistWatches, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.ExistWatches = make([]string, lenExistWatches)
	for i := 0; i < len(m.ExistWatches); i++ {
		if m.ExistWatches[i], err = in.ReadString(); err != nil {
			return err
		}
	}
	lenChildWatches, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.ChildWatches = make([]string, lenChildWatches)
	for i := 0; i < len(m.ChildWatches); i++ {
		if m.ChildWatches[i], err = in.ReadString(); err != nil {
			return err
		}
	}
	return nil
}
func (m *SetWatches) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt64(m.RelativeZxid); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.DataWatches))); err != nil {
		return err
	}
	for i := 0; i < len(m.DataWatches); i++ {
		if err = o.WriteString(m.DataWatches[i]); err != nil {
			return err
		}
	}
	if err = o.WriteInt32(int32(len(m.ExistWatches))); err != nil {
		return err
	}
	for i := 0; i < len(m.ExistWatches); i++ {
		if err = o.WriteString(m.ExistWatches[i]); err != nil {
			return err
		}
	}
	if err = o.WriteInt32(int32(len(m.ChildWatches))); err != nil {
		return err
	}
	for i := 0; i < len(m.ChildWatches); i++ {
		if err = o.WriteString(m.ChildWatches[i]); err != nil {
			return err
		}
	}
	return nil
}

type RequestHeader struct {
	Xid  int32
	Type int32
}

func (m *RequestHeader) String() string { return fmt.Sprintf("%+v", *m) }
func (m *RequestHeader) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Xid, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *RequestHeader) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Xid); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	return nil
}

type MultiHeader struct {
	Type int32
	Done bool
	Err  int32
}

func (m *MultiHeader) String() string { return fmt.Sprintf("%+v", *m) }
func (m *MultiHeader) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Done, err = in.ReadBool(); err != nil {
		return err
	}
	if m.Err, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *MultiHeader) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	if err = o.WriteBool(m.Done); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Err); err != nil {
		return err
	}
	return nil
}

type AuthPacket struct {
	Type   int32
	Scheme string
	Auth   []byte
}

func (m *AuthPacket) String() string { return fmt.Sprintf("%+v", *m) }
func (m *AuthPacket) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Scheme, err = in.ReadString(); err != nil {
		return err
	}
	if m.Auth, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *AuthPacket) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	if err = o.WriteString(m.Scheme); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Auth); err != nil {
		return err
	}
	return nil
}

type ReplyHeader struct {
	Xid  int32
	Zxid int64
	Err  int32
}

func (m *ReplyHeader) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ReplyHeader) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Xid, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Zxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Err, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *ReplyHeader) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Xid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Zxid); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Err); err != nil {
		return err
	}
	return nil
}

type GetDataRequest struct {
	Path  string
	Watch bool
}

func (m *GetDataRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetDataRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Watch, err = in.ReadBool(); err != nil {
		return err
	}
	return nil
}
func (m *GetDataRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBool(m.Watch); err != nil {
		return err
	}
	return nil
}

type SetDataRequest struct {
	Path    string
	Data    []byte
	Version int32
}

func (m *SetDataRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetDataRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *SetDataRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type SetDataResponse struct {
	Stat Stat
}

func (m *SetDataResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetDataResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if err = (&(m.Stat)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *SetDataResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = m.Stat.Encode(o); err != nil {
		return err
	}
	return nil
}

type GetSASLRequest struct {
	Token []byte
}

func (m *GetSASLRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetSASLRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Token, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *GetSASLRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteBuffer(m.Token); err != nil {
		return err
	}
	return nil
}

type SetSASLRequest struct {
	Token []byte
}

func (m *SetSASLRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetSASLRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Token, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *SetSASLRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteBuffer(m.Token); err != nil {
		return err
	}
	return nil
}

type SetSASLResponse struct {
	Token []byte
}

func (m *SetSASLResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetSASLResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Token, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *SetSASLResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteBuffer(m.Token); err != nil {
		return err
	}
	return nil
}

type CreateRequest struct {
	Path  string
	Data  []byte
	Acl   []ACL
	Flags int32
}

func (m *CreateRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CreateRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	lenAcl, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Acl = make([]ACL, lenAcl)
	for i := 0; i < len(m.Acl); i++ {
		if err = (&(m.Acl[i])).Decode(in); err != nil {
			return err
		}
	}
	if m.Flags, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *CreateRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.Acl))); err != nil {
		return err
	}
	for i := 0; i < len(m.Acl); i++ {
		if err = m.Acl[i].Encode(o); err != nil {
			return err
		}
	}
	if err = o.WriteInt32(m.Flags); err != nil {
		return err
	}
	return nil
}

type DeleteRequest struct {
	Path    string
	Version int32
}

func (m *DeleteRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *DeleteRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *DeleteRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type GetChildrenRequest struct {
	Path  string
	Watch bool
}

func (m *GetChildrenRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetChildrenRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Watch, err = in.ReadBool(); err != nil {
		return err
	}
	return nil
}
func (m *GetChildrenRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBool(m.Watch); err != nil {
		return err
	}
	return nil
}

type GetChildren2Request struct {
	Path  string
	Watch bool
}

func (m *GetChildren2Request) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetChildren2Request) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Watch, err = in.ReadBool(); err != nil {
		return err
	}
	return nil
}
func (m *GetChildren2Request) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBool(m.Watch); err != nil {
		return err
	}
	return nil
}

type CheckVersionRequest struct {
	Path    string
	Version int32
}

func (m *CheckVersionRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CheckVersionRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *CheckVersionRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type GetMaxChildrenRequest struct {
	Path string
}

func (m *GetMaxChildrenRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetMaxChildrenRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *GetMaxChildrenRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type GetMaxChildrenResponse struct {
	Max int32
}

func (m *GetMaxChildrenResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetMaxChildrenResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Max, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *GetMaxChildrenResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Max); err != nil {
		return err
	}
	return nil
}

type SetMaxChildrenRequest struct {
	Path string
	Max  int32
}

func (m *SetMaxChildrenRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetMaxChildrenRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Max, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *SetMaxChildrenRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Max); err != nil {
		return err
	}
	return nil
}

type SyncRequest struct {
	Path string
}

func (m *SyncRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SyncRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *SyncRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type SyncResponse struct {
	Path string
}

func (m *SyncResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SyncResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *SyncResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type GetACLRequest struct {
	Path string
}

func (m *GetACLRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetACLRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *GetACLRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type SetACLRequest struct {
	Path    string
	Acl     []ACL
	Version int32
}

func (m *SetACLRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetACLRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	lenAcl, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Acl = make([]ACL, lenAcl)
	for i := 0; i < len(m.Acl); i++ {
		if err = (&(m.Acl[i])).Decode(in); err != nil {
			return err
		}
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *SetACLRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.Acl))); err != nil {
		return err
	}
	for i := 0; i < len(m.Acl); i++ {
		if err = m.Acl[i].Encode(o); err != nil {
			return err
		}
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type SetACLResponse struct {
	Stat Stat
}

func (m *SetACLResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetACLResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if err = (&(m.Stat)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *SetACLResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = m.Stat.Encode(o); err != nil {
		return err
	}
	return nil
}

type WatcherEvent struct {
	Type  int32
	State int32
	Path  string
}

func (m *WatcherEvent) String() string { return fmt.Sprintf("%+v", *m) }
func (m *WatcherEvent) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.State, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *WatcherEvent) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	if err = o.WriteInt32(m.State); err != nil {
		return err
	}
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type ErrorResponse struct {
	Err int32
}

func (m *ErrorResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ErrorResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Err, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *ErrorResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Err); err != nil {
		return err
	}
	return nil
}

type CreateResponse struct {
	Path string
}

func (m *CreateResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CreateResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *CreateResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type ExistsRequest struct {
	Path  string
	Watch bool
}

func (m *ExistsRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ExistsRequest) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Watch, err = in.ReadBool(); err != nil {
		return err
	}
	return nil
}
func (m *ExistsRequest) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBool(m.Watch); err != nil {
		return err
	}
	return nil
}

type ExistsResponse struct {
	Stat Stat
}

func (m *ExistsResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ExistsResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if err = (&(m.Stat)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *ExistsResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = m.Stat.Encode(o); err != nil {
		return err
	}
	return nil
}

type GetDataResponse struct {
	Data []byte
	Stat Stat
}

func (m *GetDataResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetDataResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	if err = (&(m.Stat)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *GetDataResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = m.Stat.Encode(o); err != nil {
		return err
	}
	return nil
}

type GetChildrenResponse struct {
	Children []string
}

func (m *GetChildrenResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetChildrenResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	lenChildren, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Children = make([]string, lenChildren)
	for i := 0; i < len(m.Children); i++ {
		if m.Children[i], err = in.ReadString(); err != nil {
			return err
		}
	}
	return nil
}
func (m *GetChildrenResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(int32(len(m.Children))); err != nil {
		return err
	}
	for i := 0; i < len(m.Children); i++ {
		if err = o.WriteString(m.Children[i]); err != nil {
			return err
		}
	}
	return nil
}

type GetChildren2Response struct {
	Children []string
	Stat     Stat
}

func (m *GetChildren2Response) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetChildren2Response) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	lenChildren, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Children = make([]string, lenChildren)
	for i := 0; i < len(m.Children); i++ {
		if m.Children[i], err = in.ReadString(); err != nil {
			return err
		}
	}
	if err = (&(m.Stat)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *GetChildren2Response) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(int32(len(m.Children))); err != nil {
		return err
	}
	for i := 0; i < len(m.Children); i++ {
		if err = o.WriteString(m.Children[i]); err != nil {
			return err
		}
	}
	if err = m.Stat.Encode(o); err != nil {
		return err
	}
	return nil
}

type GetACLResponse struct {
	Acl  []ACL
	Stat Stat
}

func (m *GetACLResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (m *GetACLResponse) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	lenAcl, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Acl = make([]ACL, lenAcl)
	for i := 0; i < len(m.Acl); i++ {
		if err = (&(m.Acl[i])).Decode(in); err != nil {
			return err
		}
	}
	if err = (&(m.Stat)).Decode(in); err != nil {
		return err
	}
	return nil
}
func (m *GetACLResponse) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(int32(len(m.Acl))); err != nil {
		return err
	}
	for i := 0; i < len(m.Acl); i++ {
		if err = m.Acl[i].Encode(o); err != nil {
			return err
		}
	}
	if err = m.Stat.Encode(o); err != nil {
		return err
	}
	return nil
}

type LearnerInfo struct {
	Serverid        int64
	ProtocolVersion int32
}

func (m *LearnerInfo) String() string { return fmt.Sprintf("%+v", *m) }
func (m *LearnerInfo) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Serverid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.ProtocolVersion, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *LearnerInfo) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt64(m.Serverid); err != nil {
		return err
	}
	if err = o.WriteInt32(m.ProtocolVersion); err != nil {
		return err
	}
	return nil
}

type QuorumPacket struct {
	Type     int32
	Zxid     int64
	Data     []byte
	Authinfo []Id
}

func (m *QuorumPacket) String() string { return fmt.Sprintf("%+v", *m) }
func (m *QuorumPacket) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Zxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	lenAuthinfo, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Authinfo = make([]Id, lenAuthinfo)
	for i := 0; i < len(m.Authinfo); i++ {
		if err = (&(m.Authinfo[i])).Decode(in); err != nil {
			return err
		}
	}
	return nil
}
func (m *QuorumPacket) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Zxid); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.Authinfo))); err != nil {
		return err
	}
	for i := 0; i < len(m.Authinfo); i++ {
		if err = m.Authinfo[i].Encode(o); err != nil {
			return err
		}
	}
	return nil
}

type FileHeader struct {
	Magic   int32
	Version int32
	Dbid    int64
}

func (m *FileHeader) String() string { return fmt.Sprintf("%+v", *m) }
func (m *FileHeader) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Magic, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Dbid, err = in.ReadInt64(); err != nil {
		return err
	}
	return nil
}
func (m *FileHeader) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Magic); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Dbid); err != nil {
		return err
	}
	return nil
}

type TxnHeader struct {
	ClientId int64
	Cxid     int32
	Zxid     int64
	Time     int64
	Type     int32
}

func (m *TxnHeader) String() string { return fmt.Sprintf("%+v", *m) }
func (m *TxnHeader) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.ClientId, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Cxid, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Zxid, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Time, err = in.ReadInt64(); err != nil {
		return err
	}
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *TxnHeader) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt64(m.ClientId); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Cxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Zxid); err != nil {
		return err
	}
	if err = o.WriteInt64(m.Time); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	return nil
}

type CreateTxnV0 struct {
	Path      string
	Data      []byte
	Acl       []ACL
	Ephemeral bool
}

func (m *CreateTxnV0) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CreateTxnV0) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	lenAcl, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Acl = make([]ACL, lenAcl)
	for i := 0; i < len(m.Acl); i++ {
		if err = (&(m.Acl[i])).Decode(in); err != nil {
			return err
		}
	}
	if m.Ephemeral, err = in.ReadBool(); err != nil {
		return err
	}
	return nil
}
func (m *CreateTxnV0) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.Acl))); err != nil {
		return err
	}
	for i := 0; i < len(m.Acl); i++ {
		if err = m.Acl[i].Encode(o); err != nil {
			return err
		}
	}
	if err = o.WriteBool(m.Ephemeral); err != nil {
		return err
	}
	return nil
}

type CreateTxn struct {
	Path           string
	Data           []byte
	Acl            []ACL
	Ephemeral      bool
	ParentCVersion int32
}

func (m *CreateTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CreateTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	lenAcl, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Acl = make([]ACL, lenAcl)
	for i := 0; i < len(m.Acl); i++ {
		if err = (&(m.Acl[i])).Decode(in); err != nil {
			return err
		}
	}
	if m.Ephemeral, err = in.ReadBool(); err != nil {
		return err
	}
	if m.ParentCVersion, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *CreateTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.Acl))); err != nil {
		return err
	}
	for i := 0; i < len(m.Acl); i++ {
		if err = m.Acl[i].Encode(o); err != nil {
			return err
		}
	}
	if err = o.WriteBool(m.Ephemeral); err != nil {
		return err
	}
	if err = o.WriteInt32(m.ParentCVersion); err != nil {
		return err
	}
	return nil
}

type DeleteTxn struct {
	Path string
}

func (m *DeleteTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *DeleteTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	return nil
}
func (m *DeleteTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	return nil
}

type SetDataTxn struct {
	Path    string
	Data    []byte
	Version int32
}

func (m *SetDataTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetDataTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *SetDataTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type CheckVersionTxn struct {
	Path    string
	Version int32
}

func (m *CheckVersionTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CheckVersionTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *CheckVersionTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type SetACLTxn struct {
	Path    string
	Acl     []ACL
	Version int32
}

func (m *SetACLTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetACLTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	lenAcl, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Acl = make([]ACL, lenAcl)
	for i := 0; i < len(m.Acl); i++ {
		if err = (&(m.Acl[i])).Decode(in); err != nil {
			return err
		}
	}
	if m.Version, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *SetACLTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(int32(len(m.Acl))); err != nil {
		return err
	}
	for i := 0; i < len(m.Acl); i++ {
		if err = m.Acl[i].Encode(o); err != nil {
			return err
		}
	}
	if err = o.WriteInt32(m.Version); err != nil {
		return err
	}
	return nil
}

type SetMaxChildrenTxn struct {
	Path string
	Max  int32
}

func (m *SetMaxChildrenTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *SetMaxChildrenTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Path, err = in.ReadString(); err != nil {
		return err
	}
	if m.Max, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *SetMaxChildrenTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteString(m.Path); err != nil {
		return err
	}
	if err = o.WriteInt32(m.Max); err != nil {
		return err
	}
	return nil
}

type CreateSessionTxn struct {
	TimeOut int32
}

func (m *CreateSessionTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *CreateSessionTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.TimeOut, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *CreateSessionTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.TimeOut); err != nil {
		return err
	}
	return nil
}

type ErrorTxn struct {
	Err int32
}

func (m *ErrorTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *ErrorTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Err, err = in.ReadInt32(); err != nil {
		return err
	}
	return nil
}
func (m *ErrorTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Err); err != nil {
		return err
	}
	return nil
}

type Txn struct {
	Type int32
	Data []byte
}

func (m *Txn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *Txn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	if m.Type, err = in.ReadInt32(); err != nil {
		return err
	}
	if m.Data, err = in.ReadBuffer(); err != nil {
		return err
	}
	return nil
}
func (m *Txn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(m.Type); err != nil {
		return err
	}
	if err = o.WriteBuffer(m.Data); err != nil {
		return err
	}
	return nil
}

type MultiTxn struct {
	Txns []Txn
}

func (m *MultiTxn) String() string { return fmt.Sprintf("%+v", *m) }
func (m *MultiTxn) Decode(in Input) error {
	if m == nil {
		return nil
	}
	var err error
	lenTxns, err := in.ReadInt32()
	if err != nil {
		return err
	}
	m.Txns = make([]Txn, lenTxns)
	for i := 0; i < len(m.Txns); i++ {
		if err = (&(m.Txns[i])).Decode(in); err != nil {
			return err
		}
	}
	return nil
}
func (m *MultiTxn) Encode(o Output) error {
	if m == nil {
		return nil
	}
	var err error
	if err = o.WriteInt32(int32(len(m.Txns))); err != nil {
		return err
	}
	for i := 0; i < len(m.Txns); i++ {
		if err = m.Txns[i].Encode(o); err != nil {
			return err
		}
	}
	return nil
}
