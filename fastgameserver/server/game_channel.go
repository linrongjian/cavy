package server

import (
	"time"
)

var (
	DefaultDialTimeout = time.Second * 5
)

type GameChannel interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...ConnectorOption) (Connector, error)
	Listen(addr string, opts ...AcceptorOption) (Acceptor, error)
	String() string
}

type Acceptor interface {
	Addr() string
	Close() error
	Accept(func(GameChan)) error
}

type Connector interface {
	GameChan
}

type GameMessage struct {
	Header map[string]string
	Body   []byte
}

type GameChan interface {
	Recv(*GameMessage) error
	Send(*GameMessage) error
	Close() error
	Local() string
	Remote() string
}

type Option func(*Options)
type ConnectorOption func(*ConnectorOptions)
type AcceptorOption func(*AcceptorOptions)
