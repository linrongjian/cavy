package logconsumer

import (
	"cavy/core/network/amqp/rmqconsumer"
	"fmt"
	"os"
)

type LogConsumer interface {
	Options() Options
}

type Options struct {
	RmqConsumer rmqconsumer.RmqConsumer
}

type logConsumer struct {
	opts Options
}

type Option func(*Options)

func NewLogConsumer(opts ...Option) (LogConsumer, error) {
	options := Options{}
	logconsumer := &logConsumer{
		opts: options,
	}
	for _, o := range opts {
		o(&options)
	}

	var err error
	uri := os.Getenv("LOG_RABBITMQ_URI")
	exchange := os.Getenv("LOG_RABBITMQ_EXCHANGE")
	queue := os.Getenv("LOG_RABBITMQ_QUEUE")
	if uri == "" || exchange == "" || queue == "" {
		return nil, fmt.Errorf("log consumer no env")
	}

	conf := &rmqconsumer.Config{
		Uri:          uri,
		Exchange:     exchange,
		ExchangeType: "topic",
		Queue:        queue,
		Tag:          "simple-consumer",
		RoutingKey:   "test-key"}

	logconsumer.opts.RmqConsumer, err = rmqconsumer.NewRmqConsumer(rmqconsumer.WithConfig(conf))
	if err != nil {
		return nil, fmt.Errorf("new rmqconsumer: %s", err)
	}
	return logconsumer, err
}

func (a *logConsumer) Options() Options {
	return a.opts
}
