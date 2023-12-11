package cache

import (
	"cavy/common/mlog"

	"github.com/gomodule/redigo/redis"
)

// HashCache 哈希表缓存
type HashCache struct {
	name string
	log  *mlog.Logger
}

// NewHashCache 创建哈希表缓存对象
func NewHashCache(name string) *HashCache {
	return &HashCache{
		name: name,
		log:  mlog.NewLogger(mlog.Fields{"HashCacheName": name}),
	}
}

// Exists 是否存在
func (h *HashCache) Exists(keys ...interface{}) map[interface{}]bool {
	conn := pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	for _, k := range keys {
		conn.Send("HEXISTS", h.name, k)
	}
	exists, err := redis.Ints(conn.Do("EXEC"))
	if err != nil {
		h.log.Errorf("HEXISTS err:%v", err)
	}

	respMap := map[interface{}]bool{}
	for i, v := range keys {
		if exists[i] > 0 {
			respMap[v] = true
		} else {
			respMap[v] = false
		}
	}

	return respMap
}

// Load 加载数据
func (h *HashCache) Load(keys ...interface{}) (map[interface{}][]byte, error) {
	conn := pool.Get()
	defer conn.Close()
	respMap := map[interface{}][]byte{}
	conn.Send("MULTI")
	for _, v := range keys {
		conn.Send("HGET", h.name, v)
	}

	data, err := redis.ByteSlices(conn.Do("EXEC"))
	if err != nil {
		h.log.Errorf("HGET err:%v", err)
		return respMap, err
	}

	for i, v := range keys {
		respMap[v] = data[i]
	}

	return respMap, nil
}

// Save 保存数据
func (h *HashCache) Save(data map[interface{}][]byte) error {
	conn := pool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	for k, v := range data {
		conn.Send("HSET", h.name, k, v)
	}
	_, err := conn.Do("EXEC")
	if err != nil {
		h.log.Errorf("HSET err:%v", err)
		return err
	}
	return nil
}

// Remove 删除数据
func (h *HashCache) Remove(ids []int64) error {
	if len(ids) == 0 {
		return nil
	}
	conn := pool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	for _, v := range ids {
		conn.Send("HDEL", h.name, v)
	}
	_, err := conn.Do("EXEC")
	if err != nil {
		h.log.Errorf("HDEL err:%v", err)
		return err
	}

	return nil
}
