package mqwrap

import (
	"errors"
	"fmt"
	"servergo/core/logger"

	"github.com/streadway/amqp"
)

// var mqChannel *MqChannel

type channelType int

// MqChannel 通道
type MqChannel struct {
	ch    *amqp.Channel
	queue amqp.Queue
	kind  channelType
}

// Delivery captures the fields for a previously delivered message resident in
// a queue to be delivered by the server to a consumer from Channel.Consume or
// Channel.Get.
type Delivery = amqp.Delivery

const (
	// TopicChannelType topic交换机
	TopicChannelType = channelType(1)
	// FanoutChannelType fanout 交换机
	FanoutChannelType = channelType(2)
	// WorkerChannelType Worker交换机
	WorkerChannelType = channelType(3)
)

var (
	mqConnect           *amqp.Connection
	channelTypeValueMap = map[channelType]string{
		TopicChannelType:  "topic",
		FanoutChannelType: "fanout",
		WorkerChannelType: "worker",
	}
	channelTypeNameMap = map[channelType]string{
		TopicChannelType:  "amq.topic",
		FanoutChannelType: "amq.fanout",
		WorkerChannelType: "",
	}
)

func Startup() {
	logger.Infof("connect mq addr:%s account:%s password:%s", Opts.Host, Opts.Account, Opts.Password)
	url := fmt.Sprintf("amqp://%s:%s@%s", Opts.Account, Opts.Password, Opts.Host)
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Errorf("Conn MQ URL:%s err:%v", url, err)
	}
	mqConnect = conn
	// mqChannel, _ = NewChannel(TopicChannelType, "")
}

func (c channelType) String() string {
	if v, ok := channelTypeValueMap[c]; ok {
		return v
	}
	return ""
}

func (c channelType) Name() string {
	if v, ok := channelTypeNameMap[c]; ok {
		return v
	}
	return ""
}

// NewChannel 生成通道
func NewChannel(kind channelType, name string) (*MqChannel, error) {
	channel := &MqChannel{}
	channel.kind = kind

	if mqConnect == nil {
		logger.Errorf("MQ Connect is Null")
		return nil, fmt.Errorf("MQ Connect is Null")
	}

	ch, err := mqConnect.Channel()
	if err != nil {
		logger.Errorf("NewChannel err:%v", err)
		return nil, err
	}
	channel.ch = ch

	if kind == TopicChannelType {
		// err = ch.ExchangeDeclare(kind.Name(), kind.String(), true, false, false, false, nil)
		if err != nil {
			logger.Errorf("ExchangeDeclare err:%v", err)
			return nil, err
		}
		// queue, err := ch.QueueDeclare("", false, false, true, false, nil)
		agrs := amqp.Table{"x-message-ttl": 10000}
		queue, err := ch.QueueDeclare(name, false, true, false, false, agrs)
		if err != nil {
			logger.Errorf("QueueDeclare err:%v", err)
			return nil, err
		}
		channel.queue = queue
	} else if kind == FanoutChannelType {
		err := ch.ExchangeDeclare(kind.Name(), kind.String(), true, false, false, false, nil)
		if err != nil {
			logger.Errorf("ExchangeDeclare err:%v", err)
			return nil, err
		}
		queue, err := ch.QueueDeclare("", false, false, true, false, nil)
		if err != nil {
			logger.Errorf("QueueDeclare err:%v", err)
			return nil, err
		}
		channel.queue = queue
	} else if kind == WorkerChannelType {
		agrs := amqp.Table{"x-message-ttl": 10000}
		queue, err := ch.QueueDeclare(name, false, true, false, false, agrs)
		if err != nil {
			logger.Errorf("QueueDeclare err:%v", err)
			return nil, err
		}
		channel.queue = queue
	} else {
		return nil, errors.New("x")
	}

	return channel, nil
}

// func ChannelBind() error {

// 	ch, err := mqConnect.Channel()
// 	if err != nil {
// 		logger.Errorf("ChannelBind Err:%v", err)
// 		return err
// 	}
// 	err = ch.ExchangeBind("xw", "", "xw_exchange_fanout", false, nil)
// 	if nil != err {

// 	}

// 	return nil
// }

// Close 关闭
func (c *MqChannel) Close() {
	c.ch.QueueDelete(c.queue.Name, false, false, false)
	c.ch.ExchangeDelete(c.kind.Name(), false, false)
	c.ch.Close()
}

// Publish 广播
func (c *MqChannel) Publish(key string, msg []byte) error {
	logger.Debugf("Publish %s msgLen %d", key, len(msg))
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

// Subscribe 订阅
func (c *MqChannel) Subscribe(key string) error {
	logger.Infof("Subscribe %s", key)
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
func (c *MqChannel) Unsubscribe(key string) error {
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
func (c *MqChannel) Receive(reader func(value Delivery)) error {
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

// KickOutPlayer 踢掉玩家
// func (c *Channel) KickOutPlayer(userID string, code goddess.MessageCode) error {
// 	msg := []byte("1")

// 	if mqChannel == nil {
// 		ch, err := NewChannel(TopicChannelType, "")
// 		if err != nil {
// 			log.Errorf("KickOutPlayer newChannel TopicChannelType err:%v", err)
// 			return err
// 		}
// 		mqChannel = ch
// 	}

// 	messageCode := goddess.MessageCode(code)
// 	pushMsg := &goddess.Message{
// 		Cmd:  &messageCode,
// 		Data: []byte(msg),
// 	}
// 	buf, err := proto.Marshal(pushMsg)
// 	if err != nil {
// 		log.Errorf("KickOutPlayer Marshal pushMsg:%v err:%v", pushMsg, err)
// 		return err
// 	}

// 	err = mqChannel.Publish(userID, buf)
// 	if err != nil {
// 		log.Errorf("KickOutPlayer publish userID:%v err:%v", userID, err)
// 		mqChannel.Close()
// 		mqChannel = nil
// 		return err
// 	}

// 	return nil
// }
