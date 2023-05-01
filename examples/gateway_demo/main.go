package main

import (
	"eventgo/component/gateway"
	"eventgo/core/logger"
)

func main() {
	g := gateway.NewGateServer()
	if g.Init() != nil {
		logger.Error("gate init err")
	}
	g.Run()
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
