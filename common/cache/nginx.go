package cache

import (
	"github.com/linrongjian/cavy/common/stateless"

	"github.com/gomodule/redigo/redis"
)

// WriteNginxAddr 写nginx地址
func WriteNginxAddr(addr string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("LPUSH", stateless.NginxListKey, addr)
	if err != nil {
		return err
	}

	return nil
}

func WriteNginxAddrs(addrs []string) error {
	conn := pool.Get()
	defer conn.Close()

	args := redis.Args{}
	args = args.Add(stateless.NginxListKey)
	for _, v := range addrs {
		args = args.Add(v)
	}

	_, err := conn.Do("LPUSH", args...)
	if err != nil {
		return err
	}
	return nil
}

// ReadNginxAddrs 读取nginx地址
func ReadNginxAddrs() ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	addrs, err := redis.Strings(conn.Do("LRANGE", stateless.NginxListKey, 0, -1))
	if err != nil {
		return nil, err
	}
	return addrs, nil
}

// RemoveNginxAddr 删除nginx地址
func RemoveNginxAddr(addr string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("LREM", stateless.NginxListKey, 0, addr)
	if err != nil {
		return err
	}

	return nil
}
