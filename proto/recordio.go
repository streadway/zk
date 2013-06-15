// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package proto

import (
	"encoding/binary"
	"io"
)

type Encodable interface {
	Encode(Output) error
}

type Output interface {
	WriteBool(bool) error
	WriteInt32(int32) error
	WriteInt64(int64) error
	WriteString(string) error
	WriteBuffer([]byte) error
}

type BinaryWriter struct {
	io.Writer
}

func (w BinaryWriter) WriteBool(val bool) error {
	var err error
	if val {
		_, err = w.Write([]byte{1})
	} else {
		_, err = w.Write([]byte{0})
	}
	return err
}

func (w BinaryWriter) WriteInt32(val int32) error {
	return binary.Write(w, binary.BigEndian, val)
}

func (w BinaryWriter) WriteInt64(val int64) error {
	return binary.Write(w, binary.BigEndian, val)
}

func (w BinaryWriter) WriteString(val string) error {
	return w.WriteBuffer([]byte(val))
}

func (w BinaryWriter) WriteBuffer(val []byte) error {
	err := w.WriteInt32(int32(len(val)))
	if err != nil {
		return err
	}
	_, err = w.Write(val)
	return err
}

type Decodable interface {
	Decode(Input) error
}

type Input interface {
	ReadBool() (bool, error)
	ReadInt32() (int32, error)
	ReadInt64() (int64, error)
	ReadString() (string, error)
	ReadBuffer() ([]byte, error)
}

type BinaryReader struct {
	io.Reader
}

func (r BinaryReader) ReadBool() (bool, error) {
	buf := []byte{0}
	_, err := io.ReadFull(r, buf)
	if buf[0] == 0 {
		return true, err
	}
	return false, err
}

func (r BinaryReader) ReadInt32() (int32, error) {
	var val int32
	err := binary.Read(r, binary.BigEndian, &val)
	return val, err
}

func (r BinaryReader) ReadInt64() (int64, error) {
	var val int64
	err := binary.Read(r, binary.BigEndian, &val)
	return val, err
}

func (r BinaryReader) ReadString() (string, error) {
	val, err := r.ReadBuffer()
	return string(val), err
}

func (r BinaryReader) ReadBuffer() ([]byte, error) {
	l, err := r.ReadInt32()
	if err != nil {
		return nil, err
	}
	if l > 0 {
		val := make([]byte, l)
		_, err = io.ReadFull(r, val)
		return val, err
	}
	return nil, nil
}
