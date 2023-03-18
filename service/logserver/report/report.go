package report

import (
	"bufio"
	"fastserver/core/logger"
	"fastserver/core/network/protocols/mqwrap"
	"os"
	"strings"
)

type LogReport struct {
	mqChannel *mqwrap.MqChannel
}

func NewLogReport() *LogReport {
	r := new(LogReport)

	mqChannel, err := mqwrap.NewChannel(mqwrap.TopicChannelType, "game-event-dev")
	if err != nil {
		logger.Errorf("NewMQChannel err:%v", err)
		return nil
	}
	r.mqChannel = mqChannel

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

func (r *LogReport) Complete() {
	logger.Infof("onQmqRecv", "complete")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if input == "1" {
			r.mqChannel.Publish("test", []byte("sadfsaf"))
		}
		// inputs := strings.Split(input, " ")
		// for _, num := range inputs {
		// 	n, _ := strconv.Atoi(num)
		// 	fmt.Printf("%d ", n)
		// }
	}

}

func (r *LogReport) rmqRecv(value mqwrap.Delivery) {
	logger.Infof("onQmqRecv", string(value.Body))

	// recvMsg := &pb.Message{}
	// err := proto.Unmarshal(value.Body, recvMsg)
	// if err != nil {
	// 	logger.Errorf("Unmarshal err:%v", err)
	// 	return
	// }
}
