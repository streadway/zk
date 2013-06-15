// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func emit(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

type field struct {
	typ  string
	name string
}

var structName string
var structFields []field

type step func(*scanner.Scanner, rune) step

func fail(s *scanner.Scanner, tok rune) step {
	s.Error(s, fmt.Sprintf("unhandled: %s %s", scanner.TokenString(tok), s.TokenText()))
	return fail
}

func goVector(s *scanner.Scanner) string {
	typ := "[]byte"

	for {
		switch tok := s.Scan(); tok {
		case '<':
		case '>':
			return typ
		case scanner.EOF:
			fail(s, tok)
			panic("EOF in vector")
		case scanner.Ident:
			typ = "[]" + goType(s)
		}
	}
}

func goType(s *scanner.Scanner) string {
	switch t := s.TokenText(); t {
	case "ustring":
		return "string"
	case "long":
		return "int64"
	case "int":
		return "int32"
	case "buffer":
		return "[]byte"
	case "boolean":
		return "bool"
	case "vector":
		return goVector(s)
	default:
		return t
	}
}

func goName(s *scanner.Scanner) string {
	if s.Scan() != scanner.Ident {
		panic("expecting ident for field name")
	}
	return strings.Title(s.TokenText())
}

func emitField(s *scanner.Scanner, tok rune) step {
	f := field{goType(s), goName(s)}
	structFields = append(structFields, f)

	emit("%s %s", f.name, f.typ)

	return fields
}

func endStruct(s *scanner.Scanner, tok rune) step {
	emit("}")
	emit("")

	emit(`func (m *%s) String() string { return fmt.Sprintf("%%+v", *m) }`, structName)
	emit("func (m *%s) Decode(in Input) error {", structName)
	emit("if m == nil { return nil }")
	emit("var err error")
	for _, f := range structFields {
		switch f.typ {
		case "int32", "int64", "bool":
			emit("if m.%s, err = in.Read%s(); err != nil { return err }", f.name, strings.Title(f.typ))
		case "string":
			emit("if m.%s, err = in.ReadString(); err != nil { return err }", f.name)
		case "[]byte":
			emit("if m.%s, err = in.ReadBuffer(); err != nil { return err }", f.name)
		default:
			if len(f.typ) >= 2 && f.typ[:2] == "[]" {
				l := "len" + f.name
				emit("%s, err := in.ReadInt32()", l)
				emit("if err != nil { return err }")
				emit("m.%s = make(%s, %s)", f.name, f.typ, l)
				emit("for i := 0; i < len(m.%s); i++ {", f.name)
				if f.typ[2:] == "string" {
					emit("if m.%s[i], err = in.ReadString(); err != nil { return err }", f.name)
				} else {
					emit("if err = (&(m.%s[i])).Decode(in); err != nil { return err }", f.name)
				}
				emit("}")
			} else {
				emit("if err = (&(m.%s)).Decode(in); err != nil { return err }", f.name)
			}
		}
	}
	emit("return nil")
	emit("}")

	emit("func (m *%s) Encode(o Output) error {", structName)
	emit("if m == nil { return nil }")
	emit("var err error")
	for _, f := range structFields {
		switch f.typ {
		case "int32", "int64", "bool":
			emit("if err = o.Write%s(m.%s); err != nil { return err }", strings.Title(f.typ), f.name)
		case "string":
			emit("if err = o.WriteString(m.%s); err != nil { return err }", f.name)
		case "[]byte":
			emit("if err = o.WriteBuffer(m.%s); err != nil { return err }", f.name)
		default:
			if len(f.typ) >= 2 && f.typ[:2] == "[]" {
				emit("if err = o.WriteInt32(int32(len(m.%s))); err != nil { return err }", f.name)
				emit("for i := 0; i < len(m.%s); i++ {", f.name)
				if f.typ[2:] == "string" {
					emit("if err = o.WriteString(m.%s[i]); err != nil { return err }", f.name)
				} else {
					emit("if err = m.%s[i].Encode(o); err != nil { return err }", f.name)
				}
				emit("}")
			} else {
				emit("if err = m.%s.Encode(o); err != nil { return err }", f.name)
			}
		}
	}
	emit("return nil")
	emit("}")

	structFields = structFields[:0]

	return class
}

func fields(s *scanner.Scanner, tok rune) step {
	switch tok {
	case '.':
		return fields
	case '{':
		return fields
	case '}':
		return endStruct(s, tok)
	case ';':
		return fields
	case scanner.Ident:
		if s.Peek() == '.' {
			return fields
		}
		return emitField(s, tok)
	}
	return fail
}

func startStruct(s *scanner.Scanner, tok rune) step {
	structName = s.TokenText()
	emit("type %s struct {", structName)
	return fields
}

func class(s *scanner.Scanner, tok rune) step {
	if s.TokenText() != "class" {
		return class
	}
	return startStruct
}

func start(s *scanner.Scanner, tok rune) step {
	emit(`
// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

// GENERATED - DO NOT EDIT

package proto

import "fmt"

`)

	return class
}

func err(s *scanner.Scanner, msg string) {
	s.ErrorCount++
	fmt.Println("err: ", msg, " at: ", s.Pos())
}

func main() {
	s := new(scanner.Scanner)
	s.Init(os.Stdin)
	s.Error = err

	parse := start
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		parse = parse(s, tok)
	}
}
