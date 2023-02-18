package wswrap

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// NewConn 生成连接
func NewConn(w http.ResponseWriter, r *http.Request) (*GWsConn, error) {
	upgrader.CheckOrigin = checkOrigin
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	ws.SetReadLimit(wsReadLimit)

	id := uuid.NewV4().String()
	conn := &GWsConn{
		id:           id,
		Ws:           ws,
		wsLock:       &sync.Mutex{},
		UserData:     nil,
		lastRecvTime: time.Now(),
		lastPingTime: time.Now(),
		Isreconnect:  false,
		wg:           &sync.WaitGroup{},
	}
	conn.wg.Add(1)
	connManager.Store(id, conn)
	ws.SetCloseHandler(conn.onClose)

	ws.SetPongHandler(func(message string) error {
		//logger.Info("Pong", id, message)
		return nil
	})
	ws.SetPingHandler(func(message string) error {
		//logger.Info("Ping", id, message)
		err := ws.WriteControl(websocket.PongMessage, []byte(message), time.Now().Add(time.Second))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})

	return conn, nil
}

// 连接
func connect(scheme string, host string, path string) (*websocket.Conn, error) {
	u := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

	header := http.Header{}

	if scheme == "wss" {
		websocket.DefaultDialer.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func checkOrigin(r *http.Request) bool {
	return true
}
