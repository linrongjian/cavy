package journal

import (
	"bufio"
	"cavy/core/logger"
	"cavy/core/network/protocols/mqwrap"
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

	err = mqChannel.Subscribe("test1")
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

func (r *LogReport) Init() {
	logger.Infof("onQmqRecv", "complete")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if input == "1" {
			// for {
			// 	time.Sleep(10 * time.Nanosecond)
			// 	for i := 0; i < 1000; i++ {
			r.mqChannel.Publish("test", []byte("sadfsaf"))
			// 	}
			// }
		}
		if input == "2" {
			// for {
			// 	time.Sleep(10 * time.Nanosecond)
			// 	for i := 0; i < 1000; i++ {
			r.mqChannel.Publish("test1", []byte("sadfsaf"))
			// 	}
			// }
		}
		// inputs := strings.Split(input, " ")
		// for _, num := range inputs {
		// 	n, _ := strconv.Atoi(num)
		// 	fmt.Printf("%d ", n)
		// }
	}

}

func (r *LogReport) rmqRecv(value mqwrap.Delivery) {
	logger.Infof("onQmqRecv", value.RoutingKey, string(value.Body))

	// recvMsg := &pb.Message{}
	// err := proto.Unmarshal(value.Body, recvMsg)
	// if err != nil {
	// 	logger.Errorf("Unmarshal err:%v", err)
	// 	return
	// }
}
