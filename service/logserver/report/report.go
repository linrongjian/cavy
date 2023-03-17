package report

import (
	"fastserver/core/logger"
	"fastserver/core/network/protocols/mqwrap"
	"fastserver/core/protocol/pb"

	"google.golang.org/protobuf/proto"
)

type LogReport struct {
}

func NewLogReport() *LogReport {
	r := new(LogReport)

	mqChannel, err := mqwrap.NewChannel(mqwrap.TopicChannelType, "")
	if err != nil {
		logger.Errorf("NewMQChannel err:%v", err)
		return nil
	}

	err = mqChannel.Subscribe("test")
	if err != nil {
		logger.Errorf("Subscribe id err:%v", err)
		return nil
	}

	err = mqChannel.Receive(r.rmqRecv)
	if err != nil {
		logger.Errorf("Receive err:%v", err)
		return nil
	}
	return r
}

func (r *LogReport) rmqRecv(value mqwrap.Delivery) {
	logger.Infof("onQmqRecv", "onRecvPush")

	recvMsg := &pb.Message{}
	err := proto.Unmarshal(value.Body, recvMsg)
	if err != nil {
		logger.Errorf("Unmarshal err:%v", err)
		return
	}
}
