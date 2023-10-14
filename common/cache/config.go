package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/linrongjian/cavy/common/mlog"
	"github.com/linrongjian/cavy/common/stateless"
)

func ReadConfig(key string) string {
	conn := pool.Get()
	defer conn.Close()

	log := mlog.NewLogger(mlog.Fields{
		"Key":   stateless.ConfigKey,
		"Field": key,
	})
	cfg, err := redis.String(conn.Do("HGET", stateless.ConfigKey, key))
	if err != nil {
		log.Error("HGET ", err)
		return ""
	}
	return cfg
}

func WriteConfig(key string, config string) error {
	conn := pool.Get()
	defer conn.Close()

	log := mlog.NewLogger(mlog.Fields{
		"Key":   stateless.ConfigKey,
		"Field": key,
	})
	_, err := conn.Do("HSET", stateless.ConfigKey, key, config)
	if err != nil {
		log.Error("HSET ", err)
		return err
	}

	return nil
}
