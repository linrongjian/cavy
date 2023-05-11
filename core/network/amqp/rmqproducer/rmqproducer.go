package rmqproducer

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type RmqProducer interface {
	Init(...Option) error
	Options() Options
	Publish(body []byte) error
}

type rmqproducer struct {
	opts       Options
	conn       *amqp.Connection
	channel    *amqp.Channel
	unackedMsg sync.Map
	publishTag uint64
	once       sync.Once
}

type Option func(*Options)

func NewRmqProducer(opts ...Option) (RmqProducer, error) {
	o := Options{}
	p := &rmqproducer{
		opts: o,
	}
	if err := p.Init(opts...); err != nil {
		return nil, fmt.Errorf("consumer init: %s", err)
	}
	return p, nil
}

func (p *rmqproducer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&p.opts)
	}

	var err error

	p.once.Do(func() {

		log.Printf("dialing %q", p.opts.Config.Uri)
		p.conn, err = amqp.Dial(p.opts.Config.Uri)
		if err != nil {
			err = fmt.Errorf("dial: %s", err)
			return
		}

		go func() {
			fmt.Printf("closing: %s", <-p.conn.NotifyClose(make(chan *amqp.Error)))
		}()

		log.Printf("got Connection, getting Channel")
		p.channel, err = p.conn.Channel()
		if err != nil {
			err = fmt.Errorf("channel: %s", err)
			return
		}

		log.Printf("got Channel, declaring Exchange (%q)", p.opts.Config.Exchange)
		if err = p.channel.ExchangeDeclare(
			p.opts.Config.Exchange, // name of the exchange
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

		if p.opts.Config.Reliable {
			log.Printf("enabling publishing confirms.")
			if err = p.channel.Confirm(false); err != nil {
				err = fmt.Errorf("channel could not be put into confirm mode: %s", err)
			}
			c := p.channel.NotifyPublish(make(chan amqp.Confirmation))
			go p.confirm(c)
		}

		p.publishTag = 1
	})

	return err
}

func (c *rmqproducer) Options() Options {
	return c.opts
}

func (p *rmqproducer) Publish(body []byte) error {
	// log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)

	if err := p.channel.Publish(
		p.opts.Config.Exchange,   // publish to an exchange
		p.opts.Config.RoutingKey, // routing to 0 or more queues
		false,                    // mandatory
		false,                    // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("exchange Publish: %s", err)
	}

	p.unackedMsg.Store(p.publishTag, body)
	p.publishTag++

	return nil
}

func (p *rmqproducer) confirm(confirms <-chan amqp.Confirmation) {
	log.Printf("waiting for confirmation of publishing")
	for confirmed := range confirms {
		if confirmed.Ack {
			p.unackedMsg.Delete(confirmed.DeliveryTag)
		} else {
			log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
		}
	}
}
