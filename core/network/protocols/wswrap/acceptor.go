package wswrap

import (
	"CavyGo/core/logger"
	"CavyGo/core/network/transport"
	"CavyGo/core/protocol/pb"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type GWsConn struct {
	id           string //连接ID
	Ws           *websocket.Conn
	wg           *sync.WaitGroup
	wsLock       *sync.Mutex //ws 并发写锁
	UserData     interface{} //给上层使用，可以继连接绑定数据
	lastRecvTime time.Time
	lastPingTime time.Time
	Isreconnect  bool   //是否重连
	addr         string //连接地址
	path         string //请求路径
	scheme       string
}

// GetID 获得ID
func (c *GWsConn) GetID() string {
	return c.id
}

// SendText 发送文本数据
func (c *GWsConn) SendText(data []byte) error {
	if c.Ws == nil {
		msg := fmt.Sprintf("SendText Ws is nil")
		return fmt.Errorf(msg)
	}
	c.wsLock.Lock()
	err := c.Ws.WriteMessage(websocket.TextMessage, data)
	c.wsLock.Unlock()
	return err
}

// SendKick 发送踢下线消息
func (c *GWsConn) SendKick(kickcode int) error {
	if c.Ws == nil {
		msg := fmt.Sprintf("SendText Ws is nil")
		return fmt.Errorf(msg)
	}
	c.wsLock.Lock()
	err := c.Ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("kick_%d", kickcode)))
	c.wsLock.Unlock()
	return err
}

func (c *GWsConn) onClose(code int, msg string) error {
	if _, ok := connManager.Load(c.id); ok {
		connManager.Delete(c.id)
		c.wg.Done()
	}
	return nil
}

// WaitWebSocket 等待
func (c *GWsConn) WaitWebSocket(recv func(c *GWsConn, cmd pb.Cmd, data []byte)) {
	if recv == nil {
		return
	}

	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {

			logger.Errorf("read message err:", err)
			if err := c.Close(); err != nil {
				return
			}
			break
		}
		c.lastRecvTime = time.Now()

		recvMsg := &pb.Message{}
		err = proto.Unmarshal(message, recvMsg)
		if err != nil {
			continue
		}

		if recvMsg.GetCmd() == pb.Cmd_Ping {
			heartMsg := &pb.HeartBeat{}
			err = proto.Unmarshal(recvMsg.GetData(), heartMsg)
			if err != nil {
				continue
			}
			nowTime := int32(time.Now().Unix())
			heartMsg.NowTime = &nowTime
			buf, err := proto.Marshal(heartMsg)
			if err != nil {
				continue
			}
			c.sendPong(buf)
		} else if recvMsg.GetCmd() == pb.Cmd_Pong {
			c.lastPingTime = time.Now()
		} else if recv != nil {
			recv(c, recvMsg.GetCmd(), recvMsg.GetData())
		}
	}
	c.wg.Wait()
}

func (c *GWsConn) Rev(*transport.Message) error {
	return nil
}

func (c *GWsConn) Send(m *transport.Message) error {

	if c.Ws == nil {
		msg := fmt.Sprintf("Send Ws is nil")
		return fmt.Errorf(msg)
	}
	c.wsLock.Lock()
	err := c.Ws.WriteMessage(websocket.BinaryMessage, m.Body)
	c.wsLock.Unlock()

	return err
}

func (c *GWsConn) Close() error {
	var err error
	if !c.Isreconnect {
		err = c.Ws.Close()
	} else {
		err = nil
	}

	if _, ok := connManager.Load(c.id); ok {
		connManager.Delete(c.id)
		c.wg.Done()
	}

	return err
}

func (c *GWsConn) Local() string {
	return ""
}

func (c *GWsConn) Remote() string {
	return ""
}

// SendProto 发送proto
func (c *GWsConn) SendProto(m proto.Message) error {
	buf, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	return c.SendBytes(buf)
}

// SendBytes 发送二进制数据†
func (c *GWsConn) SendBytes(data []byte) error {
	if c.Ws == nil {
		msg := fmt.Sprintf("Send Ws is nil")
		return fmt.Errorf(msg)
	}
	c.wsLock.Lock()
	err := c.Ws.WriteMessage(websocket.BinaryMessage, data)
	c.wsLock.Unlock()

	return err
}

// sendPong 发送pong
func (c *GWsConn) sendPong(data []byte) {
	pongCmd := pb.Cmd_Pong

	pongMsg := &pb.PushMessage{
		Id:   proto.String(uuid.NewV4().String()),
		Cmd:  &pongCmd,
		Data: data,
	}

	buf, err := proto.Marshal(pongMsg)
	if err != nil {
		return
	}
	if err := c.SendBytes(buf); err != nil {
		return
	}
}
