// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

package zk

import (
	"sync"
)

// EventType is one of 4 node events emitted by the server
type EventType int32

const (
	Created EventType = 1 // From Exists, Get
	Deleted EventType = 2 // From Exists, Get
	Changed EventType = 3 // From Exists, Get
	Child   EventType = 4 // From Children
)

// Event contains the type of change, and path of that change.  It will be
// delivered once per watch.
type Event struct {
	Type EventType
	Path string
}

// Watches synchronizes access to an internal map of paths to event chans.
// Event chans receive a single Event and are removed from this map when the
// watch at that path is triggered.
type Watches struct {
	m   sync.Mutex
	reg map[string][]chan<- Event
}

// Add registers an event chan at a path.
func (w *Watches) add(path string, ch chan<- Event) {
	w.m.Lock()
	defer w.m.Unlock()

	if w.reg == nil {
		w.reg = make(map[string][]chan<- Event)
	}

	w.reg[path] = append(w.reg[path], ch)
}

// Pop removes and returns all event chans for a given path.  Returns nil if
// the path has no watches.
func (w *Watches) pop(path string) []chan<- Event {
	w.m.Lock()
	defer w.m.Unlock()

	chans := w.reg[path]
	delete(w.reg, path)
	return chans
}

// Paths returns a slice of unique paths being watched.
func (w *Watches) paths() []string {
	w.m.Lock()
	defer w.m.Unlock()

	res := make([]string, 0, len(w.reg))
	for path, _ := range w.reg {
		res = append(res, path)
	}
	return res
}
