package amqp

type Amqp interface {
	Init(opts ...Option)
	Options() Options
}

type Channel interface {
	Init(opts ...Option)
	Options() Options
	Publish(msg []byte)
	Consume(msg []byte)
}

type Connection interface {
	Init(opts ...Option)
	Options() Options
	Reconnect() error
}

type Consumer interface {
	Init(...ConsumerOption) error
	Options() ConsumerOptions
	Shutdown() error
}

type ConsumerHandle func(delivery Delivery, done chan error)
