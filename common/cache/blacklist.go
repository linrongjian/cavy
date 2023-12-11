package cache

import (
	"fmt"
	"time"

	"cavy/common/stateless"

	"github.com/gomodule/redigo/redis"
)

// AddRequestNum 增加请求次数
func AddRequestNum(playerid int64, ttl int64, num int32) error {
	con := pool.Get()
	defer con.Close()

	key := fmt.Sprintf("%s%d", stateless.RequestNum, playerid)

	_, err := con.Do("SETEX", key, ttl, num)
	if err != nil {
		return err
	}

	return nil
}

// GetRequestNum 获取请求次数
func GetRequestNum(playerid int64) int32 {
	con := pool.Get()
	defer con.Close()

	key := fmt.Sprintf("%s%d", stateless.RequestNum, playerid)

	num, err := redis.Int(con.Do("GET", key))
	if err != nil {
		return 0
	}

	return int32(num)
}

// AddBlackList 加入黑名单
func AddBlackList(playerid int64, ttl int64) error {
	con := pool.Get()
	defer con.Close()

	key := fmt.Sprintf("%s%d", stateless.BlackList, playerid)

	_, err := con.Do("SETEX", key, ttl, 1)
	if err != nil {
		return err
	}

	_, err = con.Do("HSET", stateless.BlackListBak, playerid, time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
}

// IsInBlackList 是否在黑名单
func IsInBlackList(playerid int64) bool {
	con := pool.Get()
	defer con.Close()

	key := fmt.Sprintf("%s%d", stateless.BlackList, playerid)

	exist, err := redis.Bool(con.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exist
}
