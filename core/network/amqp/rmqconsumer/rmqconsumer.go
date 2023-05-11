package rmqconsumer

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type RmqConsumer interface {
	Init(...Option) error
	Options() Options
	Shutdown() error
}

type Handle func(body []byte)

type rmqconsumer struct {
	opts    Options
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error

	once sync.Once
}

type Option func(*Options)

func NewRmqConsumer(opts ...Option) (RmqConsumer, error) {
	o := Options{}
	c := &rmqconsumer{
		opts: o,
	}
	if err := c.Init(opts...); err != nil {
		return nil, fmt.Errorf("consumer init: %s", err)
	}
	return c, nil
}

func (c *rmqconsumer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&c.opts)
	}

	var err error

	c.once.Do(func() {

		log.Printf("dialing %q", c.opts.Config.Uri)
		c.conn, err = amqp.Dial(c.opts.Config.Uri)
		if err != nil {
			err = fmt.Errorf("dial: %s", err)
			return
		}

		go func() {
			fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		}()

		log.Printf("got Connection, getting Channel")
		c.channel, err = c.conn.Channel()
		if err != nil {
			err = fmt.Errorf("channel: %s", err)
			return
		}

		log.Printf("got Channel, declaring Exchange (%q)", c.opts.Config.Exchange)
		if err = c.channel.ExchangeDeclare(
			c.opts.Config.Exchange, // name of the exchange
			"direct",               // type
			true,                   // durable
			false,                  // delete when complete
			false,                  // internal
			false,                  // noWait
			nil,                    // arguments
		); err != nil {
			err = fmt.Errorf("exchange Declare: %s", err)
			return
		}

		log.Printf("declared Exchange, declaring Queue %q", c.opts.Config.Queue)
		var queue amqp.Queue
		queue, err = c.channel.QueueDeclare(
			c.opts.Config.Queue, // name of the queue
			true,                // durable
			false,               // delete when unused
			false,               // exclusive
			false,               // noWait
			nil,                 // arguments
		)
		if err != nil {
			err = fmt.Errorf("queue Declare: %s", err)
			return
		}

		log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
			queue.Name, queue.Messages, queue.Consumers, c.opts.Config.RoutingKey)

		if err = c.channel.QueueBind(
			queue.Name,               // name of the queue
			c.opts.Config.RoutingKey, // bindingKey
			c.opts.Config.Exchange,   // sourceExchange
			false,                    // noWait
			nil,                      // arguments
		); err != nil {
			err = fmt.Errorf("queue bind: %s", err)
			return
		}

		log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.opts.Config.Tag)

		var deliveries <-chan amqp.Delivery
		deliveries, err = c.channel.Consume(
			queue.Name,        // name
			c.opts.Config.Tag, // consumerTag,
			false,             // noAck
			false,             // exclusive
			false,             // noLocal
			false,             // noWait
			nil,               // arguments
		)
		if err != nil {
			err = fmt.Errorf("queue Consume: %s", err)
			return
		}

		go c.handle(deliveries)
	})

	return err
}

func (c *rmqconsumer) Options() Options {
	return c.opts
}

func (c *rmqconsumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.opts.Config.Tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

func (c *rmqconsumer) handle(deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		// log.Printf(
		// 	"got %dB delivery: [%v] %q",
		// 	len(d.Body),
		// 	d.DeliveryTag,
		// 	d.Body,
		// )
		d.Ack(false)
		if c.opts.Handle != nil {
			c.opts.Handle(d.Body)
		}
	}
	log.Printf("handle: deliveries channel closed")
	c.done <- nil
}
