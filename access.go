// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package zk

import (
	"github.com/streadway/zk/proto"
)

// Permission is a bitmask of permissions that apply to a node for a scheme's identity.
type Permission int32

const (
	PermRead Permission = 1 << iota
	PermWrite
	PermCreate
	PermDelete
	PermAdmin

	PermAll Permission = (1 << iota) - 1
)

// Access represents permissions bitmask for an identity under a scheme
type Access struct {
	Perms    Permission // Bitmask of Permissions
	Scheme   string     // one of "world" "auth" "digest" "host" "ip"
	Identity string     // Scheme specific identity like 127.0.0.1/32
}

// ACL is an Access Control List used in Create, SetAcl and GetAcl
type ACL []Access

// Commonly used ACLs for nodes
var (
	AclOpen     = ACL{Access{PermAll, "world", "anyone"}}
	AclReadOnly = ACL{Access{PermRead, "world", "anyone"}}
)

func toProtoACLs(acl ACL) []proto.ACL {
	out := make([]proto.ACL, len(acl))
	for i, a := range acl {
		out[i] = proto.ACL{
			Perms: int32(a.Perms),
			Id: proto.Id{
				Scheme: a.Scheme,
				Id:     a.Identity,
			},
		}
	}
	return out
}

func fromProtoACLs(acls []proto.ACL) ACL {
	out := make(ACL, len(acls))
	for i, a := range acls {
		out[i] = Access{
			Perms:    Permission(a.Perms),
			Scheme:   a.Id.Scheme,
			Identity: a.Id.Id,
		}
	}
	return out
}
