package gateway

import (
	"cavy/core/logger"
	"cavy/core/network/protocols/httpwrap"
	"cavy/core/network/protocols/wswrap"
	"cavy/core/protocol/pb"
	"math/rand"
	"strconv"

	"github.com/golang/protobuf/proto"
)

var (
	count = 1
)

// onConnectHandle 连接
func onConnectHandle(c *httpwrap.Context) {
	//logger.Info("connectHandle enter")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(UnknownError)),
	}
	defer c.WriteRsp(&httpRsp)

	playerId := strconv.Itoa(rand.Int()) //c.Query.Get("playerid")
	if playerId == "" {
		httpRsp.Result = proto.Int32(int32(ErrParamNil))
		httpRsp.Msg = proto.String("playerid请求参数为空")
		return
	}

	logger.Info("ws connect:", count)
	count++

	// session := c.Query.Get("session")
	// if session == "" {
	// 	httpRsp.Result = proto.Int32(int32(ErrParamNil))
	// 	httpRsp.Msg = proto.String("session请求参数为空")
	// 	return
	// }

	// logger.Debugf("ws connect: playerId:%v", playerId)

	// old := rds.Do(rds.CmdGet, rkey.KeyUserToken+playerId)

	// if old.Err() == nil && old.String() != session {
	// 	httpRsp.Result = proto.Int32(int32(ErrParam))
	// 	httpRsp.Msg = proto.String("session请求参数错误")
	// 	return
	// }

	// 升级为websocket
	ws, err := wswrap.NewConn(c.W, c.GetHTTPRequest())
	if err != nil {
		httpRsp.Result = proto.Int32(int32(ErrConnect))
		httpRsp.Msg = proto.String("连接失败")
		return
	}

	defer func() {
		// zlgame.DeletePlayer(playerId)
		ws.Close()
		ws.UserData = nil
		// err = notifyOnlineHandle(c, playerId, false)
		if err != nil {
		}
	}()

	// player := NewPlayer(playerId, session, ws)
	// if player == nil {
	// 	httpRsp.Result = proto.Int32(int32(ErrNewPlayer))
	// 	httpRsp.Msg = proto.String("创建玩家失败")
	// 	return
	// }
	// ws.UserData = player

	// err = notifyOnlineHandle(c, playerId, true)
	// if err != nil {
	// }

	// zlgame.PrintOnline()

	ws.WaitWebSocket(receiveHandle)

	httpRsp.Result = proto.Int32(int32(SUCCESS))
	return
}
