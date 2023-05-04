package gateway

import (
	"eventgo/core/network/protocols/wswrap"
	"eventgo/core/protocol/pb"

	"github.com/golang/protobuf/proto"
)

// onReceiveHandler 收到数据处理 (返回错误就会断开连接)
func receiveHandle(c *wswrap.GWsConn, cmd pb.Cmd, data []byte) {
	player := c.UserData.(*Player)
	if player == nil {
		return
	}

	if cmd == pb.Cmd_MsgSubscribe {
		var sub pb.Subscribe
		err := proto.Unmarshal(data, &sub)
		if err != nil {
			return
		}

		for _, v := range sub.Channel {
			player.Subscribe(v)
		}
	} else if cmd == pb.Cmd_MsgUnsubscribe {
		var unsub pb.Unsubscribe
		err := proto.Unmarshal(data, &unsub)
		if err != nil {
			return
		}

		for _, v := range unsub.Channel {
			player.Unsubscribe(v)
		}
	}
}
