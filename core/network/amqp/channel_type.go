package amqp

type channelType int

const (
	// TopicChannelType topic交换机
	TopicChannelType = channelType(1)
	// FanoutChannelType fanout 交换机
	FanoutChannelType = channelType(2)
	// WorkerChannelType Worker交换机
	WorkerChannelType = channelType(3)
)

var (
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
