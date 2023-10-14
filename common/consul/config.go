package consul

import (
	"context"
	"fmt"
	"strings"
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/linrongjian/cavy/common/api"
	"github.com/linrongjian/cavy/common/hook"
	"github.com/sirupsen/logrus"
)

const innerPrefix = "config/"

var confMgr struct {
	client *consul.Client
}

func initConf() {
	//log := mlog.NewLogger(mlog.Fields{})

	hook.AddHook(func(s *api.Conf) {
		if s.ConsulAddr == "" {
			return
		}
		config := consul.DefaultConfig()
		config.Address = s.ConsulAddr
		c, err := NewClient(config)
		if err != nil {
			panic(fmt.Errorf("get consul client fail:%s", err.Error()))
		}
		confMgr.client = c
	})
}

// innerKey 增加 "config/" 前缀
func innerKey(key string) string {
	return innerPrefix + key
}

// outerKey 删除 "config/" 前缀
func outerKey(innerKey string) string {
	return strings.TrimPrefix(innerKey, innerPrefix)
}

// GetAndWatchPrefix 先getPrefix并回调一次f，然后WatchPrefix
func GetAndWatchPrefix(prefix string, f func(map[string][]byte)) (cancel func()) {
	confs, index, err := getPrefix(prefix)
	if err == nil {
		f(confs)
	}
	cancel = watchPrefix(prefix, index, f)
	return
}

// getPrefix 获取以prefix为前缀的所有配置
// 返回配置列表
func getPrefix(prefix string) (map[string][]byte, uint64, error) {
	kv := GetClient().KV()

	pairs, meta, err := kv.List(innerKey(prefix), &consul.QueryOptions{})
	if err != nil {
		return nil, 0, fmt.Errorf("consul return err: %s", err.Error())
	}
	configs := make(map[string][]byte, len(pairs))
	for _, pair := range pairs {
		key := strings.TrimPrefix(outerKey(pair.Key), prefix)
		configs[key] = pair.Value
	}
	var lastIndex uint64
	if meta != nil {
		lastIndex = meta.LastIndex
	}
	return configs, lastIndex, nil
}

// watchPrefix 监听以 prefix 为前缀的配置项的变化，并且回调 f
// index 为 GetPrefix 的返回值。
// f 为配置更新处理函数。 f 被 Watch 的内部 goroutine 调用。
// 返回值 cancel 可用于取消监听
func watchPrefix(prefix string, index uint64, f func(map[string][]byte)) (cancel func()) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("conf watchPrefix panic")
			}
		}()

		kv := GetClient().KV()
		for {
			pairs, meta, err := kv.List(innerKey(prefix), &consul.QueryOptions{
				WaitIndex: index,
				WaitTime:  time.Minute,
			})

			select {
			case <-ctx.Done():
				return
			default:
			}

			if err != nil {
				logrus.WithError(err).Errorf("consul return error,prefix=%s ", prefix)
				time.Sleep(time.Minute)
				continue
			}
			configs := make(map[string][]byte, len(pairs))
			for _, pair := range pairs {
				if pair.ModifyIndex > index {
					key := strings.TrimPrefix(outerKey(pair.Key), prefix)
					configs[key] = pair.Value
				}
			}
			if len(configs) > 0 {
				f(configs)
			}
			if meta != nil {
				index = meta.LastIndex
			}
		}
	}()
	return cancel
}
