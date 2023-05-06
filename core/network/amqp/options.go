package amqp

import "github.com/streadway/amqp"

type Options struct {
	Url  string
	Conn *amqp.Connection
}

type ChannelOptions struct {
	RabbitmqURL string
}
