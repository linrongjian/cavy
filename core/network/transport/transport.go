package transport

import (
	"time"
)

var (
	DefaultDialTimeout = time.Second * 5
)

type Transport interface {
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

type Message struct {
	Header map[string]string
	Body   []byte
}

type Channel interface {
	Recv(*Message) error

	Send(*Message) error

	Close() error

	Local() string

	Remote() string
}

type Option func(*Options)

type ConnectorOption func(*ConnectorOptions)

type AcceptorOption func(*AcceptorOptions)
