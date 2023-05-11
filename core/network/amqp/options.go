package amqp

import "github.com/streadway/amqp"

type Options struct {
	Url  string
	Conn *amqp.Connection
}

type ChannelOptions struct {
	RabbitmqURL string
}

type ConsumerOptions struct {
	Uri          string
	exchange     string
	exchangeType string
	que          string
	Tag          string
	Key          string
	Handle       ConsumerHandle
	Conn         *amqp.Connection
	Channel      *amqp.Channel
}

type ConsumerOption func(*ConsumerOptions)

type Config struct {
}
