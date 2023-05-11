package logproducer

import (
	"cavy/core/network/amqp/rmqproducer"
	"fmt"
	"os"
)

type LogProducer interface {
	Options() Options
}

type Options struct {
	RmqProducer rmqproducer.RmqProducer
}

type logProducer struct {
	opts Options
}

type Option func(*Options)

func NewLogProducer(opts ...Option) (LogProducer, error) {
	options := Options{}
	logproducer := &logProducer{
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

	conf := &rmqproducer.Config{
		Uri:          uri,
		Exchange:     exchange,
		ExchangeType: "topic",
		RoutingKey:   "test-key",
		Reliable:     true,
	}

	logproducer.opts.RmqProducer, err = rmqproducer.NewRmqProducer(rmqproducer.WithConfig(conf))
	if err != nil {
		return nil, fmt.Errorf("new rmqconsumer: %s", err)
	}
	return logproducer, err
}

func (p *logProducer) Options() Options {
	return p.opts
}
