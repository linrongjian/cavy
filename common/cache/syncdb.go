package cache

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/linrongjian/cavy/common/stateless"
)

type Syncdb struct {
	PlayerID int64 `json:"playerid"`
	Player   bool  `json:"player"`
	Hero     bool  `json:"hero"`
	Props    bool  `json:"props"`
	Friend   bool  `json:"friend"`
}

// AddPlayerSyncDB 添加db同步信息
func AddPlayerSyncDB(playerid int64, sync *Syncdb) error {
	con := pool.Get()
	defer con.Close()

	data, err := json.Marshal(sync)
	if err != nil {
		return err
	}

	_, err = con.Do("HSET", stateless.SyncDatabase, playerid, data)
	if err != nil {
		return err
	}

	return nil
}

// GetPlayerSyncDB 获取db同步信息
func GetPlayerSyncDB(playerid int64) *Syncdb {
	con := pool.Get()
	defer con.Close()

	sync := &Syncdb{}
	data, err := redis.Bytes(con.Do("HGET", stateless.SyncDatabase, playerid))
	if err != nil {
		return sync
	}

	err = json.Unmarshal(data, sync)
	if err != nil {
		return sync
	}

	return sync
}

// DelPlayerSyncDB 删除db同步信息
func DelPlayerSyncDB(playerid int64) error {
	con := pool.Get()
	defer con.Close()

	_, err := con.Do("HDEL", stateless.SyncDatabase, playerid)
	if err != nil {
		return err
	}

	return nil
}

// GetNeedSyncPlayer 获取需同步db的用户
func GetNeedSyncPlayer() []int64 {
	con := pool.Get()
	defer con.Close()

	playerids, err := redis.Int64s(con.Do("HKEYS", stateless.SyncDatabase))
	if err != nil {
		return nil
	}

	return playerids
}
