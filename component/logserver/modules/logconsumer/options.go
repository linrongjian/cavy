package logconsumer

import "github.com/streadway/amqp"

type Options struct {
	Conn *amqp.Connection
}
