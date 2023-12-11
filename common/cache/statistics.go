package cache

import (
	"encoding/json"
	"time"

	"github.com/linrongjian/cavy/common/mlog"
	"github.com/linrongjian/cavy/common/stateless"

	"github.com/gomodule/redigo/redis"
)

// 统计对象信息
type stStatistics struct {
	NewPlayer      map[int64]bool `json:"NewPlayer"`      //新用户
	OldPlayer      map[int64]bool `json:"OldPlayer"`      //老用户
	RechargePlayer map[int64]bool `json:"RechargePlayer"` //充值用户
	RechargeCount  int64          `json:"RechargeCount"`  //付费值
	LastTime       time.Time      `json:"LastTime"`       //最后统计时间
}

// 统计对象
type Statistics struct {
	*stStatistics
	log *mlog.Logger
}

// NewStatistics 创建统计对象
func NewStatistics() *Statistics {
	obj := &Statistics{
		stStatistics: &stStatistics{
			NewPlayer:      map[int64]bool{},
			OldPlayer:      map[int64]bool{},
			RechargePlayer: map[int64]bool{},
		},
		log: mlog.NewLogger(nil),
	}

	if obj.has() {
		obj.load()
	}

	// 创建的时候检测
	obj.ExpireResetData()

	return obj
}

// 是否存在
func (s *Statistics) has() bool {
	conn := pool.Get()
	defer conn.Close()

	has, err := redis.Bool(conn.Do("EXISTS", stateless.StatisticsKey))
	if err != nil {
		s.log.Errorf("EXISTS %s err:%v", stateless.StatisticsKey, err)
		return false
	}

	return has
}

// 加载数据
func (s *Statistics) load() {
	conn := pool.Get()
	defer conn.Close()

	buf, err := redis.Bytes(conn.Do("GET", stateless.StatisticsKey))
	if err != nil {
		s.log.Errorf("GET %s err:%v", stateless.StatisticsKey, err)
		return
	}

	err = json.Unmarshal(buf, s.stStatistics)
	if err != nil {
		s.log.Errorf("Unmarshal err:%v", err)
		return
	}
}

// Save 保存数据
func (s *Statistics) Save() error {
	conn := pool.Get()
	defer conn.Close()

	s.stStatistics.LastTime = time.Now()
	buf, err := json.Marshal(s.stStatistics)
	if err != nil {
		s.log.Errorf("Marshal err:%v", err)
		return err
	}

	_, err = conn.Do("SET", stateless.StatisticsKey, buf)
	if err != nil {
		s.log.Errorf("SET %s err:%v", stateless.StatisticsKey, err)
		return err
	}

	return nil
}

// ExpireResetData 过期重置数据
func (s *Statistics) ExpireResetData() {
	now := time.Now()
	if s.LastTime.Year() != now.Year() || s.LastTime.YearDay() != now.YearDay() {
		s.NewPlayer = map[int64]bool{}
		s.RechargePlayer = map[int64]bool{}
		s.RechargeCount = 0
		s.OldPlayer = GetAllOnlinePlayer()
		s.LastTime = now
	}
}
