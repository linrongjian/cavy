package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/linrongjian/cavy/common/api"
	"github.com/linrongjian/cavy/common/hook"
	"github.com/linrongjian/cavy/common/mlog"
)

// baseCache 基础缓存数据结构
type baseCache struct {
	Log *mlog.Logger
}

var (
	addr string      //redis地址
	pool *redis.Pool //redis 连接池
)

func init() {
	hook.AddHook(func(s *api.Conf) {
		if !s.ConnRedis {
			return
		}
		addr = s.RedisServer.Addr
		if addr == "" {
			addr = "127.0.0.1:6379"
		}
		log := mlog.NewLogger(mlog.Fields{"Redis": addr, "DataBase": s.RedisServer.DataBase})
		log.Info("Connect Redis")
		//建立redis连接池
		pool = newPool(addr, s.RedisServer.DataBase)
		if !CheckRedisConnect() {
			msg := fmt.Sprintf("Connect Redis %s failed!", addr)
			panic(msg)
		}
	})
}

// Startup 启动 正常情况下不需要调用
func Startup(server string, database int) bool {
	if pool != nil {
		return true
	}

	addr = server
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	//建立redis连接池
	pool = newPool(addr, database)
	if !CheckRedisConnect() {
		msg := fmt.Sprintf("Connect Redis %s failed!", addr)
		panic(msg)
	}
	return true
}

func Connect(host string, port int, database int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	if pool == nil {
		pool = newPool(addr, database)
	}
	rsp := fmt.Sprintf("RedisConnect: %s", addr)
	if !CheckRedisConnect() {

		return errors.New("Err" + rsp)
	}
	fmt.Print(rsp)
	return nil
}

func newPool(addr string, database int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return conn, err
			}
			if database != 0 {
				conn.Do("SELECT", database)
			}
			return conn, nil
		},
	}
}

// CheckRedisConnect 测试连接redis
func CheckRedisConnect() bool {
	con := pool.Get()
	defer con.Close()
	if con.Err() != nil {
		return false
	}
	return true
}

// Test ..
func Test() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", "AAAAA", "Name", "SSSSS")
	if err != nil {
		return
	}
}

// GetRedisPool 获得redis连接对象池
func GetRedisPool() *redis.Pool {
	return pool
}
