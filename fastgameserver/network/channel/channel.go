package channel

import (
	"time"
)

var (
	DefaultDialTimeout = time.Second * 5
)

type GChannel interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...ConnectorOption) (Connector, error)
	Listen(addr string, opts ...AcceptorOption) (Acceptor, error)
	String() string
}

type Acceptor interface {
	Addr() string
	Close() error
	Accept(func(GChan)) error
}

type Connector interface {
	GChan
}

type GMessage struct {
	Header map[string]string
	Body   []byte
}

type GChan interface {
	Recv(*GMessage) error
	Send(*GMessage) error
	Close() error
	Local() string
	Remote() string
}

type Option func(*Options)
type ConnectorOption func(*ConnectorOptions)
type AcceptorOption func(*AcceptorOptions)
