// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package zk

import (
	"errors"
	"fmt"
	"io"
)

var log = func(f string, a ...interface{}) { fmt.Printf(f+"\n", a...) }
var logln = fmt.Println

type errCode int32

var (
	ErrExists     = errors.New("zk: node already exists")   // Create
	ErrInvalidAcl = errors.New("zk: ACL is not supported")  // Create
	ErrVersion    = errors.New("zk: node version differs")  // Set/Delete
	ErrNoNode     = errors.New("zk: no node found at path") // Exists/Get

	ErrExpired    = errors.New("zk: session expired")         // All
	ErrConnection = errors.New("zk: connection lost")         // All
	ErrAuth       = errors.New("zk: connection unauthorized") // All

	ErrProtocol = errors.New("zk: protocol error")

	errMap = map[errCode]error{
		errNodeExists: ErrExists,
		errBadVersion: ErrVersion,
		errNoNode:     ErrNoNode,
		errInvalidAcl: ErrInvalidAcl,
	}
)

func ioError(err error) error {
	switch err {
	case io.EOF:
		return ErrConnection
	}
	return err
}

func (e errCode) toError() error {
	err := errMap[e]
	if err == nil && e != 0 {
		panic(fmt.Errorf("unknown error: %d", e))
	}
	return err
}

const (
	// API errors
	//errApiError                = errCode(-100)
	errNoNode errCode = -101
	//errNoAuth                  = errCode(-102)
	errBadVersion = errCode(-103)
	//errNoChildrenForEphemerals = errCode(-108)
	errNodeExists = errCode(-110)
	//errNotEmpty                = errCode(-111)
	errSessionExpired = errCode(-112)
	//errInvalidCallback         = errCode(-113)
	errInvalidAcl = errCode(-114)
	//errAuthFailed              = errCode(-115)
	//errClosing                 = errCode(-116)
	//errNothing                 = errCode(-117)
	//errSessionMoved            = errCode(-118)
)
