# Go Zookeeper client

Implements the Zookeeper protocol in Go.

# Status

Early Alpha - WIP

Docs: http://godoc.org/github.com/streadway/zk

Code: http://github.com/streadway/zk

# Contributing

Fork, write tests, fail tests, fix stuff, pass tests, `go fmt`, `go vet`, open
a pull request.

# TODO

  * Reconnection or other kind of connection management
  * SetAuth during connection establishment
  * SASL (blocked on a kerberos domain setup to test)
  * Multi transactions (blocked on finding a good API)
  * Acceptance tests for flakey connections
  * Acceptance tests for server partitions

# License

BSD 2 clause - see LICENSE for more details.
