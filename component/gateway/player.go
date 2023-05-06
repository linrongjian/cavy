package gateway

import (
	"CavyGo/core/app"
	"CavyGo/core/logger"
	"CavyGo/core/network/protocols/mqwrap"
	"CavyGo/core/network/protocols/wswrap"
	"CavyGo/core/protocol/pb"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	KickAlreadyLogin = 3000 //已经登录
	KickBlackList    = 3001 //黑名单被踢
	KickSealPlayer   = 3002 //封号被踢
)

// Player 用户接口
type Player struct {
	id        string
	ws        *wswrap.GWsConn
	mqChannel *mqwrap.MqChannel
	wg        *sync.WaitGroup
	session   string
}

var (
	conns       sync.Map
	lock        = &sync.Mutex{}
	countOnline int32
)

// NewPlayer 创建用户
func NewPlayer(playerID string, session string, conn *wswrap.GWsConn) *Player {
	lock.Lock()
	defer lock.Unlock()

	// 相同session不踢
	if p, has := conns.Load(playerID); has && p.(*Player).session != session {
		logger.Infof("kick, close ws playerid:%s %s", playerID, p.(*Player).GetConnectID())
		if err := p.(*Player).SendKick(KickAlreadyLogin); err != nil {
			logger.Error("Send kick err:", err.Error())
		}

		p.(*Player).ws.Close()
		p.(*Player).wg.Wait()
	}

	// 这里要wait一下，等playerdelete完
	startts := time.Now().Unix()
	for {
		if _, has := conns.Load(playerID); !has {
			break
		}

		endts := time.Now().Unix()
		if endts-startts > 5 {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	player := &Player{
		id:      playerID,
		ws:      conn,
		wg:      &sync.WaitGroup{},
		session: session,
	}
	player.wg.Add(1)

	mqChannel, err := mqwrap.NewChannel(mqwrap.TopicChannelType, "")
	if err != nil {
		logger.Errorf("NewMQChannel err:%v", err)
		return nil
	}
	player.mqChannel = mqChannel

	err = mqChannel.Subscribe(playerID)
	if err != nil {
		logger.Errorf("Subscribe id err:%v", err)
		return nil
	}

	// 注意，这里是单服通知
	subKey := fmt.Sprintf("allplayer-%v", app.Conf.ServerID)
	// log.Infof("subkey:%v", subKey)
	err = mqChannel.Subscribe(subKey)
	if err != nil {
		logger.Errorf("Subscribe allplayer err:%v", err)
		return nil
	}

	// // 注意，这里是全服通知
	// err = mqChannel.Subscribe("allplayer")
	// if err != nil {
	// 	log.Errorf("Subscribe allplayer err:%v", err)
	// 	return nil
	// }

	err = mqChannel.Receive(player.onRecvPush)
	if err != nil {
		logger.Errorf("Receive err:%v", err)
		return nil
	}

	conns.Store(playerID, player)

	return player
}

// GetPlayer 获得用户
func GetPlayer(id string) *Player {
	player, has := conns.Load(id)
	if !has {
		return nil
	}
	return player.(*Player)
}

// FindPlayer 查找用户
func FindPlayer(playerID string) *Player {
	if p, exists := conns.Load(playerID); exists {
		return p.(*Player)
	}
	return nil
}

// DeletePlayer 删除用户
func DeletePlayer(playerID string) {
	player := FindPlayer(playerID)
	if player != nil {
		player.wg.Done()
		player.mqChannel.Close()
		player.mqChannel = nil
	}

	conns.Delete(playerID)
}

// 用户接收推送函数
func (p *Player) onRecvPush(value mqwrap.Delivery) {

	logger.Info("func", "onRecvPush")

	recvMsg := &pb.Message{}
	err := proto.Unmarshal(value.Body, recvMsg)
	if err != nil {
		logger.Errorf("Unmarshal err:%v", err)
		return
	}

	if *recvMsg.Cmd == pb.Cmd_KickBlackList {
		if err := p.SendKick(KickBlackList); err != nil {
			logger.Error("Send kick err:", err.Error())
		}
		p.ws.Close()
		return
	} else if *recvMsg.Cmd == pb.Cmd_KickSealPlayer {
		if err := p.SendKick(KickSealPlayer); err != nil {
			logger.Error("Send kick err:", err.Error())
		}
		p.ws.Close()
		return
	}

	// id := uuid.NewV4().String()
	// log.Infof("Player:%s Recv Push ID:%s Cmd:%d DataLen:%d", p.id, id, int(*recvMsg.Cmd), len(recvMsg.Data))
	//msgAckMap[id] = value

	pushMsg := &pb.PushMessage{
		Id:   proto.String(""),
		Cmd:  recvMsg.Cmd,
		Data: recvMsg.Data,
	}
	p.ws.SendProto(pushMsg)
}

// GetPlayerID 获取ID
func (p *Player) GetPlayerID() string {
	return p.id
}

// Close 关闭
func (p *Player) Close() {
	if p.ws != nil {
		p.ws.Close()
		p.ws.Isreconnect = true
		atomic.AddInt32(&countOnline, -1)
	}
	p.wg.Wait()
}

// SendKick 关闭客户端
func (p *Player) SendKick(kickcode int) error {
	if p.ws != nil {
		return p.ws.SendKick(kickcode)
	}
	return nil
}

// GetConnectID 获得连接Id
func (p *Player) GetConnectID() string {
	return p.ws.GetID()
}

// Subscribe 订阅频道
func (p *Player) Subscribe(key string) {
	err := p.mqChannel.Subscribe(key)
	if err != nil {
	}
}

// Unsubscribe 取消订阅频道
func (p *Player) Unsubscribe(key string) {
	err := p.mqChannel.Unsubscribe(key)
	if err != nil {
	}
}

// ClearOnline 清除在线用户
func ClearOnline() {
	//conn := db.Connect()
	//defer db.Disconnect()
	//
	//serveridstr := fmt.Sprintf("%d", servercfg.ServerID)
	//
	//conn.Send("MULTI")
	//conn.Send("DEL", rconst.SetStatisticsOnlineNewPlayerPrefix+serveridstr)
	//conn.Send("DEL", rconst.SetStatisticsOnlineOldPlayerPrefix+serveridstr)
	//conn.Send("DEL", rconst.SetOnlinePrefix+serveridstr)
	//_, err := conn.Do("EXEC")
	//if err != nil {
	//	logger.Error("exec err, err:", err.Error())
	//	return
	//}
}

// ClearOnline 清除在线用户
func PrintOnline() {
	atomic.AddInt32(&countOnline, 1)
	logger.Info("在线连接数:", countOnline)
}
