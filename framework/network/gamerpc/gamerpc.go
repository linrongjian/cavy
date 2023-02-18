package gamerpc

import (
	"time"
)

var (
	DefaultDialTimeout = time.Second * 5
)

type GameRpc interface {
	Init(...Option) error

	Options() Options

	Dial(addr string, opts ...ConnectorOption) (Connector, error)

	Listen(addr string, opts ...AcceptorOption) (Acceptor, error)

	String() string
}

type Acceptor interface {
	Addr() string

	Close() error

	Accept(func(Channel)) error
}

type Connector interface {
	Channel
}

type GMessage struct {
	Header map[string]string
	Body   []byte
}

type Channel interface {
	Recv(*GMessage) error

	Send(*GMessage) error

	Close() error

	Local() string

	Remote() string
}

type Option func(*Options)

type ConnectorOption func(*ConnectorOptions)

type AcceptorOption func(*AcceptorOptions)
