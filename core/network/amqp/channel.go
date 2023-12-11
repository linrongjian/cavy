package amqp

import (
	"fmt"

	"github.com/linrongjian/cavy/core/logger"

	"github.com/streadway/amqp"
)

type channel struct {
	ch    *amqp.Channel
	queue amqp.Queue
	kind  channelType
}

func (c *channel) Subscribe(key string) error {
	if c.kind == TopicChannelType {
		err := c.ch.QueueBind(c.queue.Name, key, c.kind.Name(), false, nil)
		if err != nil {
			logger.Errorf("Topic Subscribe err:%v", err)
			return err
		}
		return nil
	} else if c.kind == FanoutChannelType {
		logger.Errorf("Fanou Subscribe")
		return nil
	} else if c.kind == WorkerChannelType {
		logger.Errorf("Worker Subscribe")
		return nil
	}
	return fmt.Errorf("Subscribe Kind type:%d err", c.kind)
}

// Unsubscribe 取消订阅
func (c *channel) Unsubscribe(key string) error {
	logger.Debugf("Unsubscribe %s", key)
	if c.kind == TopicChannelType {
		err := c.ch.QueueUnbind(c.queue.Name, key, c.kind.Name(), nil)
		if err != nil {
			logger.Errorf("Topic Unsubscribe err:%v", err)
			return err
		}
		return nil
	} else if c.kind == FanoutChannelType {
		logger.Errorf("Fanou Unsubscribe")
		return nil
	} else if c.kind == WorkerChannelType {
		logger.Errorf("Worker Unsubscribe")
		return nil
	}
	return fmt.Errorf("Unsubscribe Kind type:%d err", c.kind)

}

// Receive 接受
func (c *channel) Receive(reader func(value Delivery)) error {
	logger.Debug("Receive")
	if c.kind == WorkerChannelType {
		err := c.ch.Qos(1, 0, false)
		if err != nil {
			logger.Errorf("Workder Receive Qos err:%v", err)
			return err
		}
	} else if c.kind == FanoutChannelType {
		err := c.ch.QueueBind(c.queue.Name, "", c.kind.Name(), false, nil)
		if err != nil {
			logger.Errorf("Fanout Subscribe err:%v", err)
			return err
		}
	} else if c.kind != TopicChannelType {
		logger.Errorf("Receive Kind type:%d err", c.kind)
		return fmt.Errorf("Receive Kind type:%d err", c.kind)
	}

	msgs, err := c.ch.Consume(c.queue.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Errorf("Topic Receive err:%v", err)
		return err
	}
	go func() {
		for value := range msgs {
			reader(value)
		}
	}()

	return nil
}

// Publish 广播
func (c *channel) Publish(key string, msg []byte) error {
	// logger.Debugf("Publish %s msgLen %d", key, len(msg))
	amqpMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}
	if c.kind == TopicChannelType {
		err := c.ch.Publish(c.kind.Name(), key, false, false, amqpMsg)
		if err != nil {
			logger.Errorf("Topic Publish err:%v", err)
			return err
		}
	} else if c.kind == FanoutChannelType {
		err := c.ch.Publish(c.kind.Name(), "", false, false, amqpMsg)
		if err != nil {
			logger.Errorf("Fanout Publish err:%v", err)
			return err
		}
	} else if c.kind == WorkerChannelType {
		amqpMsg.DeliveryMode = amqp.Persistent
		err := c.ch.Publish(c.kind.Name(), c.queue.Name, false, false, amqpMsg)
		if err != nil {
			logger.Errorf("Worker Publish err:%v", err)
			return err
		}
	} else {
		logger.Errorf("Publish kind type:%v err", c.kind)
		return fmt.Errorf("Publish kind type:%v err", c.kind)
	}

	return nil
}

func (c *channel) Close() {
	c.ch.QueueDelete(c.queue.Name, false, false, false)
	c.ch.ExchangeDelete(c.kind.Name(), false, false)
	c.ch.Close()
}
