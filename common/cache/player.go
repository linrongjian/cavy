package cache

import (
	"github.com/linrongjian/cavy/common/mlog"
	"github.com/linrongjian/cavy/common/stateless"

	"github.com/gomodule/redigo/redis"
)

// GetOnline 获取在线状态
func GetOnline(playerid int64) bool {
	log := mlog.NewLogger(nil)

	conn := pool.Get()
	defer conn.Close()

	online, err := redis.Bool(conn.Do("SISMEMBER", stateless.Online, playerid))
	if err != nil {
		log.Errorf("SADD err:%v", err)
		return false
	}

	return online
}

// SetOnline 设置在线
func SetOnline(playerid int64) {
	log := mlog.NewLogger(nil)

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", stateless.Online, playerid)
	if err != nil {
		log.Errorf("SADD err:%v", err)
	}

	return
}

// SetOffline 设置不在线
func SetOffline(playerid int64) {
	log := mlog.NewLogger(nil)

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SREM", stateless.Online, playerid)
	if err != nil {
		log.Errorf("SADD err:%v", err)
	}

	return
}

// GetPlayersOnline 批量获取用户在线状态
func GetPlayersOnline(playerids []int64) (onlines map[int64]bool) {
	log := mlog.NewLogger(nil)

	conn := pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	for _, v := range playerids {
		conn.Send("SISMEMBER", stateless.Online, v)
	}
	values, err := redis.Ints(conn.Do("EXEC"))
	if err != nil {
		log.Errorf("exec err:%v", err)
		return nil
	}

	onlines = map[int64]bool{}
	for i := range values {
		if values[i] == 1 {
			onlines[playerids[i]] = true
		} else {
			onlines[playerids[i]] = false
		}
	}

	return onlines
}

// 获得在线数量
func GetOnlineCount() int {
	conn := pool.Get()
	defer conn.Close()
	count, err := redis.Int(conn.Do("SCARD", stateless.Online))

	if err != nil {
		return 0
	}
	return count
}

// 获得全部在线用户
func GetAllOnlinePlayer() map[int64]bool {
	conn := pool.Get()
	defer conn.Close()

	vals, err := redis.Int64s(conn.Do("SMEMBERS", stateless.Online))
	if err != nil {
		return nil
	}

	players := map[int64]bool{}
	for _, v := range vals {
		players[v] = true
	}

	return players
}

// GetOnlinePlayersWithCount 获得指定数目的在线用户
func GetOnlinePlayersWithCount(count int32) []int64 {
	conn := pool.Get()
	defer conn.Close()

	playerids, err := redis.Int64s(conn.Do("SRANDMEMBER", stateless.Online, count))
	if err != nil {
		return nil
	}

	return playerids
}

// ClearOnline 清除在线表
func ClearOnline() {
	log := mlog.NewLogger(nil)

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", stateless.Online)
	if err != nil {
		log.Errorf("SADD err:%v", err)
		return
	}

	return
}

// GetAllPlayerIds 获取全服玩家id
func GetAllPlayerIds() []int64 {
	log := mlog.NewLogger(nil)

	playerids := []int64{}

	conn := pool.Get()
	defer conn.Close()

	data, err := redis.Int64s(conn.Do("HKEYS", stateless.PlayerKey))
	if err != nil {
		log.Errorf("HKEYS err:%v", err)
		return playerids
	}

	playerids = data

	return playerids
}
