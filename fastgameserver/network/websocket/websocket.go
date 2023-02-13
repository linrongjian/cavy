package websocket

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	wsReadLimit       = 64 * 1024
	wsReadBufferSize  = 4 * 1024
	wsWriteBufferSize = 4 * 1024
	keeperAwakeTime   = 10 * time.Second // 心跳时间
	pingTime          = 20 * time.Second // 发送ping 时间
)

var (
	upgrader = websocket.Upgrader{ReadBufferSize: wsReadBufferSize, WriteBufferSize: wsWriteBufferSize}

	// 连接管理
	connManager sync.Map

	// 是否启用检查心跳
	isKeeperAwake = false
)
