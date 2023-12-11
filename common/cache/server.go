package cache

import (
	"encoding/json"

	"github.com/linrongjian/cavy/common/stateless"

	"github.com/gomodule/redigo/redis"
)

// Server 服务器信息
type Server struct {
	ID      string `json:"ID"`      //ID
	Version string `json:"Version"` //版本
	Addr    string `json:"Addr"`    //地址
	Name    string `json:"Name"`    //服务器名称
}

// WriteServerInfo 写服务器信息
func WriteServerInfo(server *Server) error {
	conn := pool.Get()
	defer conn.Close()

	buf, err := json.Marshal(server)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSET", stateless.ServerHashKey(server.Name), server.ID, buf)
	if err != nil {
		return err
	}

	return nil
}

// ReadServerInfo 读取服务器信息
func ReadServerInfo(name string, id string) *Server {
	conn := pool.Get()
	defer conn.Close()

	buf, err := redis.Bytes(conn.Do("HGET", stateless.ServerHashKey(name), id))
	if err != nil {
		return nil
	}

	server := &Server{}
	err = json.Unmarshal(buf, server)
	if err != nil {
		return nil
	}

	return server
}

// RemoveServerInfo 删除服务器信息
func RemoveServerInfo(name string, id string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", stateless.ServerHashKey(name), id)
	if err != nil {
		return err
	}

	return nil
}
