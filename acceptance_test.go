// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package zk

import (
	"testing"
	"time"
)

func acceptance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping acceptance test when -short flag is used")
	}
}

// testSessionPool creates a buffered channel and fills it with results from
// testSession(t)
func testSessionPool(t *testing.T, count int) chan *Session {
	clients := make(chan *Session, count)
	go func() {
		for i := 0; i < count; i++ {
			clients <- testSession(t)
		}
	}()
	return clients
}

// spins until an increment sticks
func testIncrementSpin(t *testing.T, pool chan *Session, count int, path string) {
	c := <-pool
	defer func() { pool <- c }()

	for inc := 0; inc < count; inc++ {
	Retry:
		data, stat, err := c.Get(path, nil)
		if err != nil {
			t.Fatalf("expected get to always succeed, got: %q", err)
		}
		data[0] = data[0] + 1

		_, err = c.Set(path, data, stat.Version)
		if err == ErrVersion {
			goto Retry
		}
		if err != nil {
			t.Fatalf("expected increments to succeed or fail due to conflict, got: %q", err)
		}
	}
}

func TestIncrementRace(t *testing.T) {
	acceptance(t)

	withSession(t, func(c *Session) {
		path := "/test/acceptance_sequence"
		pool := testSessionPool(t, 10)

		c.Delete(path, -1)

		if _, err := c.Create(path, []byte{0}, 0, AclOpen); err != nil && err != ErrExists {
			t.Fatalf("expected no error, or ErrExists, got: %q", err)
		}

		for i := 0; i < cap(pool); i++ {
			go testIncrementSpin(t, pool, cap(pool), path)
		}
		for i := 0; i < cap(pool); i++ {
			(<-pool).Close()
		}

		data, _, err := c.Get(path, nil)
		if err != nil {
			t.Fatalf("expected to get the final value, got: %v", err)
		}
		if int(data[0]) != cap(pool)*cap(pool) {
			t.Fatalf("expected to increment %d times, got %d", cap(pool)*cap(pool), data[0])
		}

		c.Delete(path, -1)
	})
}

func TestDialExpiredSession(t *testing.T) {
	acceptance(t)
	withSession(t, func(c *Session) {
		c.Close()

		c2, err := Dial(c.Config)
		if err == nil {
			c2.Close()
		}
		if err != ErrExpired {
			t.Fatalf("expected to dial an expired session, got: %q", err)
		}
	})
}

func TestReconnectSetsWatches(t *testing.T) {
	acceptance(t)

	path := "/test/should_contain_foobar"
	watch := make(chan Event)

	withSession(t, func(c *Session) {
		_, _, err := c.Get(path, watch)
		if err != nil {
			t.Fatalf("expected to get %q, got: %q", path, err)
		}

		c.conn.Close()

		_, err = c.Set(path, []byte("foobar"), -1)
		for err == ErrConnection {
			c = testSessionConfig(t, c.Config)
			_, err = c.Set(path, []byte("foobar"), -1)
		}
		if err != nil {
			t.Fatalf("expected to set %q, got: %q", path, err)
		}

		select {
		case changed := <-watch:
			if changed.Path != path {
				t.Fatalf("expected watched path to be %q, got: %q", path, changed.Path)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout on watch to be triggered after reconnect")
		}
	})
}

func TestUnbufferedWatchIsReentrant(t *testing.T) {
	acceptance(t)

	withSession(t, func(c *Session) {
		var watches = []struct {
			path string
			ch   chan Event
		}{
			{"/test/should_exist_soon0", make(chan Event)},
			{"/test/should_exist_soon1", make(chan Event)},
			{"/test/should_exist_soon2", make(chan Event)},
		}

		clean := func() {
			for _, w := range watches {
				c.Delete(w.path, -1)
			}
		}

		clean()

		for _, w := range watches {
			c.Exists(w.path, w.ch)
		}

		for _, w := range watches {
			c.Create(w.path, []byte{}, CreatePersistent, AclOpen)
		}

		select {
		case e := <-watches[0].ch:
			c.Exists(e.Path, watches[0].ch)
		case e := <-watches[1].ch:
			c.Exists(e.Path, watches[1].ch)
		case e := <-watches[2].ch:
			c.Exists(e.Path, watches[2].ch)
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("timeout waiting for creation")
		}

		clean()
	})
}

func TestRepeatedUnbufferedWatch(t *testing.T) {
	acceptance(t)

	path := "/test/should_contain_foobar"
	watch := make(chan Event)
	const count = 10
	withSession(t, func(c *Session) {
		c.Get(path, watch)

		for i := 0; i < count; i++ {
			go c.Set(path, []byte("foobar"), -1)
			<-watch
			c.Get(path, watch)
		}
	})
}

func TestKeepalive(t *testing.T) {
	acceptance(t)

	path := "/test/should_contain_foobar"
	withSession(t, func(c *Session) {
		time.Sleep(c.Timeout + time.Second)
		_, _, err := c.Get(path, nil)
		if err != nil {
			t.Fatalf("expected to keep a session alive, got: %q", err)
		}
	})
}

func TestConcurrent(t *testing.T) {
	acceptance(t)

	path := "/test/should_have_concurrent_reader_writer"
	withSession(t, func(c *Session) {
		sets := make(chan int, 100)
		done := make(chan bool)

		c.Create(path, []byte{0}, CreateEphemeral, AclOpen)
		defer c.Delete(path, -1)

		go func() {
			defer close(sets)
			for i := 1; i <= cap(sets); i++ {
				if _, err := c.Set(path, []byte{byte(i)}, -1); err != nil {
					t.Fatalf("expected to set path, got: %q", err)
				}
			}
		}()

		go func() {
			defer close(done)
			for _ = range sets {
				if _, _, err := c.Get(path, nil); err != nil {
					t.Fatalf("expected to get path, got: %q", err)
				}
			}
		}()

		<-done

		body, _, err := c.Get(path, nil)
		if err != nil {
			t.Fatalf("expected to finally get path, got: %q", err)
		}
		if len(body) == 0 || body[0] != byte(cap(sets)) {
			t.Fatalf("expected to set all values, got %v", body)
		}
	})
}

func TestSessionRedial(t *testing.T) {
	acceptance(t)

	path := "/test/should_be_ephemeral"
	withSession(t, func(c *Session) {
		var (
			err   error
			tries = 5

			// try for 2 session timeout periods to create 5 sessions.
			ticks = time.Tick(c.Config.Timeout * 2 / time.Duration(tries))
		)

		if _, err := c.Create(path, []byte{42}, CreateEphemeral, AclOpen); err != nil {
			t.Fatalf("expected to create ephemeral node, got: %q", err)
		}

		for ; tries > 0; tries-- {
			<-ticks
			if c, err = Dial(c.Config); err != nil {
				t.Fatalf("expected to redial, got: %q", err)
			}
		}

		data, _, err := c.Get(path, nil)
		if err != nil {
			t.Fatalf("expected to get ephemeral node after session timeout, got: %q", err)
		}

		if len(data) == 0 || data[0] != 42 {
			t.Fatalf("expected to get node data after session timeout, got: %q", data)
		}

		if err := c.Delete(path, -1); err != nil {
			t.Fatalf("expected to cleanup ephemeral node, got: %q", err)
		}
	})
}
