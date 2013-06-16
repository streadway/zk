// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package zk

import (
	"reflect"
	"testing"
	"time"
)

func testSessionConfig(t *testing.T, conf *Config) *Session {
	c, err := Dial(conf)
	if err != nil {
		panic("can't connect: " + err.Error())
	}
	return c
}

func testSession(t *testing.T) *Session {
	return testSessionConfig(t, &Config{Addrs: []string{"127.0.0.1:2181"}})
}

func withSession(t *testing.T, f func(*Session)) {
	c := testSession(t)
	defer c.Close()
	f(c)
}

func TestGetDataNoExist(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/alwaysmissing"
		_, _, err := c.Get(path, nil)
		if err != ErrNoNode {
			t.Fatalf("expected node to be missing at %q, got err: %v", path, err)
		}
	})
}

func TestGetDataExist(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"

		c.Set(path, []byte("foobar"), -1)

		v, stat, err := c.Get(path, nil)
		if err != nil {
			t.Fatalf("expected no err at %q, got: %v", path, err)
		}

		if ex, s := "foobar", string(v); ex != s {
			t.Fatalf("expected content: %q, got: %q", ex, s)
		}

		if stat.Czxid == 0 {
			t.Fatalf("expected stat to have Czxid, got: %+v", stat)
		}
	})
}

func TestCreateEphemeral(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_be_ephemeral"
		defer c.Delete(path, -1)

		if path, err := c.Create(path, []byte("foobar"), CreateEphemeral, AclOpen); err != nil {
			t.Fatalf("expected to create ephemeral node at %q, got: %v", path, err)
		}

		withSession(t, func(c2 *Session) {
			if _, stat, _ := c2.Get(path, nil); stat.EphemeralOwner == 0 {
				t.Fatalf("expected ephemeral node have an owner, it did not", path)
			}
		})
	})
}

func TestCreateConflict(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"

		if path, err := c.Create(path, []byte("foobar"), 0, AclOpen); err != ErrExists {
			t.Fatalf("expected not to create ephemeral node at %q, got: %q", path, err)
		}
	})
}

func TestCreateEphemeralSequence(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"

		newPath, err := c.Create(path,
			[]byte("foobar"),
			CreateEphemeral|CreateSequence,
			AclOpen,
		)

		if err != nil {
			t.Fatalf("expected to create an ephemeral sequence node at %q, got: %q", path, err)
		}

		if newPath == path {
			t.Fatal("expected a new path to be made for a sequence node, got the same")
		}
	})
}

func TestChildrenNoExist(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/alwaysmissing"

		_, _, err := c.Children(path, nil)

		if err != ErrNoNode {
			t.Fatalf("expected parent node to be missing at %q, got err: %v", path, err)
		}
	})
}

func TestChildren(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test"

		paths, stat, err := c.Children(path, nil)

		if err != nil {
			t.Fatalf("expected to list children, got: %v", err)
		}

		if len(paths) != int(stat.NumChildren) {
			t.Fatal("expected number of children to match stat value, got %d != %d", len(paths), stat.NumChildren)
		}

		if len(paths) == 0 {
			t.Fatal("expected to receive at least one child, did not receive any: %+v", stat)
		}
	})
}

func TestSet(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"

		stat, err := c.Set(path, []byte("foobar"), -1)

		if err != nil {
			t.Fatalf("expected to set node value, got: %v", err)
		}

		if stat.Version == 0 {
			t.Fatalf("expected increment version, got: %+v", stat)
		}
	})
}

func TestCloseOnly(t *testing.T) {
	withSession(t, func(c *Session) {
		if err := c.Close(); err != nil {
			t.Fatalf("expected to close a fresh session, got: %v", err)
		}
	})
}

func TestCloseEphemeral(t *testing.T) {
	path := "/test/create_and_expire"

	withSession(t, func(c *Session) {
		if _, err := c.Create(path, []byte("foobar"), CreateEphemeral, AclOpen); err != nil {
			t.Fatalf("expected to create a path to delete, got: %v", err)
		}
	})

	withSession(t, func(c *Session) {
		if found, _, err := c.Exists(path, nil); found || err != nil {
			t.Fatalf("expected close to remove ephemeral nodes, got: %v %q", found, err)
		}
	})
}

func TestDelete(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/create_and_delete"

		if _, err := c.Create(path, []byte("foobar"), CreateEphemeral, AclOpen); err != nil {
			t.Fatalf("expected to create a path to delete, got: %v", err)
		}

		if err := c.Delete(path, -1); err != nil {
			t.Fatalf("expected to delete, got: %v", err)
		}
	})
}

func TestSync(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"

		if _, err := c.Sync(path); err != nil {
			t.Fatalf("expected sync not to error, got: %q", err)
		}
	})
}

func TestExists(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"

		found, stat, err := c.Exists(path, nil)
		if err != nil {
			t.Fatalf("expected exists not to error, got: %q", err)
		}

		if !found {
			t.Fatalf("expected to find node at %q", path)
		}

		if stat.Czxid == 0 {
			t.Fatalf("expected a meaningful stat, got: %+v", stat)
		}
	})
}

func TestNotExists(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/alwaysmissing"

		found, stat, err := c.Exists(path, nil)
		if err != nil {
			t.Fatalf("expected exists not to error, got: %q", err)
		}

		if found {
			t.Fatalf("expected to not find a missing node at %q", path)
		}

		empty := Stat{}
		if empty != stat {
			t.Fatalf("expected a missing stat when the node is missing, got: %+v", stat)
		}
	})
}

func TestWatchExists(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/sometimes_missing"
		watch := make(chan Event, 1)

		c.Delete(path, -1)

		found, _, err := c.Exists(path, watch)
		if err != nil {
			t.Fatalf("expected to set a watch for existence, got: %q", err)
		}
		if found {
			t.Fatalf("expected no node at existence watch")
		}

		withSession(t, func(c2 *Session) {
			if _, err := c2.Create(path, []byte{}, 0, AclOpen); err != nil {
				t.Fatalf("expected to create %q, got: %q", path, err)
			}

			select {
			case ev := <-watch:
				if ev.Path != path {
					t.Fatalf("expected watch path to match node, got: %+v", ev)
				}
			case <-time.After(100000 * time.Millisecond):
				t.Fatalf("timeout on watch at: %q", path)
			}
		})
	})
}

func TestWatchGet(t *testing.T) {
	withSession(t, func(c *Session) {
		path := "/test/should_contain_foobar"
		watch := make(chan Event, 1)

		_, _, err := c.Get(path, watch)
		if err != nil {
			t.Fatalf("expected to set a watch for data, got: %q", err)
		}

		withSession(t, func(c2 *Session) {
			if _, err := c2.Set(path, []byte("foobar"), -1); err != nil {
				t.Fatalf("expected to set %q, got: %q", path, err)
			}

			select {
			case ev := <-watch:
				if ev.Path != path {
					t.Fatalf("expected watch path to match node, got: %+v", ev)
				}
			case <-time.After(100000 * time.Millisecond):
				t.Fatalf("timeout on watch at: %q", path)
			}
		})
	})
}

func TestGetAcl(t *testing.T) {
	path := "/test/should_change_acl"
	withSession(t, func(c *Session) {
		defer c.Delete(path, -1)

		_, _, err := c.GetAcl(path)
		if err != ErrNoNode {
			t.Fatalf("expected missing node, got: %v", err)
		}

		c.Create(path, []byte{}, CreatePersistent, AclOpen)

		acl, stat, _ := c.GetAcl(path)
		if !reflect.DeepEqual(acl, AclOpen) {
			t.Fatalf("expected AclOpen for GetAcl")
		}

		if stat.Mzxid <= 0 {
			t.Fatalf("expected stat to contain the modified transaction id, got: %#v", stat)
		}
	})
}

func TestSetAcl(t *testing.T) {
	path := "/test/should_change_acl"
	withSession(t, func(c *Session) {
		defer c.Delete(path, -1)

		c.Create(path, []byte{}, CreatePersistent, AclOpen)

		stat, err := c.SetAcl(path, AclReadOnly, -1)
		if err != nil {
			t.Fatalf("expected SetAcl to succeed, got: %v", err)
		}
		if stat.Aversion <= 0 {
			t.Fatalf("expected SetAcl to bump the version, got: %#v", stat)
		}

		acl, _, _ := c.GetAcl(path)
		if !reflect.DeepEqual(acl, AclReadOnly) {
			t.Fatalf("expected AclReadOnly for GetAcl")
		}
	})
}
