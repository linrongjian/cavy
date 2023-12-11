package cavy

import (
	"cavy/component/logserver/logconsumer"
	"cavy/component/logserver/logproducer"
	"cavy/core/network/amqp/rmqconsumer"
)

var (
	Logconsumer logconsumer.LogConsumer
)

// type Options struct {
// 	Logconsumer logconsumer.LogConsumer
// }

// type Option func(*Options)

// type Cavy interface {
// 	Options() Options
// }

// type cavy struct {
// opts Options
// }

// func NewCavy(opts ...Option) (Cavy, error) {
// 	options := Options{}
// 	c := &cavy{
// 		opts: options,
// 	}
// 	for _, o := range opts {
// 		o(&options)
// 	}
// 	return c, nil
// }

// func (c *cavy) Options() Options {
// 	return c.opts
// }

func NewLogconsumer(h rmqconsumer.Handle) (logconsumer.LogConsumer, error) {
	logconsumer, err := logconsumer.NewLogConsumer()
	if err != nil {
		return nil, err
	}
	logconsumer.Options().RmqConsumer.Init(rmqconsumer.WithHandle(h))
	return logconsumer, nil
}

func NewLogProducer() (logproducer.LogProducer, error) {
	logproducer, err := logproducer.NewLogProducer()
	if err != nil {
		return nil, err
	}
	return logproducer, nil
}
