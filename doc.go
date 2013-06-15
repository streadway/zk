// Copyright (c) 2013, Sean Treadway, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/streadway/zk

/*
Native client for the Zookeeper protocol.

Stability

Alpha.  This package is still under development, the APIs may change.

Sessions

A Session represents an active connection to a Zookeeper server and an identity
in a Zookeeper ensemble.  Sessions can be moved between connections within an
ensemble by using Dial with a previous Session's config.

	// Move or reconnect a session
	next, err := Dial(prev.Config)

Watches from the previous connection will be moved to the next connection.
After a successful Dial, previous sessions should be considered disconnected.

Watches

The Zookeeper ensemble ensures that watches are totally ordered against all
responses.  Since this library presents a synchronous interface to a Zookeeper
server, it maintains order of watches relative to other watches.

A change in data or children is observable before the watch fires.  This means
that when a watch channel receives an event, the change can be observed with a
separate call.

Errors

Each API method can return one of many standard errors.  These errors
communicate the state of the session in the ensemble.  It is up to the
application to respond to these errors specifically to the application's needs.

		ErrConnection --> Disconnected
		ErrExpired -----> Expired
		ErrAuth --------> Unauthenticated

When disconnected, the outstanding requests may or may not have been committed
to the ensemble.  Commands are only committed when no error is returned so it
is up to the application to determine the strategy for repeating the command.
Applications should Dial a new connection with the same session config to retry
the command that caused the disconnect.

When expired, a new session must be established.  Any ephemeral nodes will be
deleted and must be re-created by the new session in a new connection.

Dial may return ErrAuth.   ErrAuth indicates the session is expired and the
connection closed.  The application must use different authentication
credentials and Dial again.

*/
package zk
