package redis

import (
	"context"
	"servergo/core/util"
	"fmt"
	"sync"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

var (
	single *Store
	once   sync.Once
)

func S() *Store {
	once.Do(func() {
		single = NewStore()
	})
	return single
}

func Connect() error {
	return S().Connect()
}

func Disconnect() error {
	return S().Disconnect()
}

func Client() *redis.Client {
	return S().client
}

func Locker() *redislock.Client {
	return S().Locker()
}

func Do(args ...interface{}) *redis.Cmd {
	return S().Do(args...)
}

// Lock 分布式锁 (默认生存周期, 默认重试时间)
func Lock(ctx context.Context, key string, index interface{}) (*redislock.Lock, error) {
	return LockBackoff(ctx, key, index, 3*time.Second, 100*time.Millisecond)
}

// LockTTL 分布式锁 (指定生存周期)
func LockTTL(ctx context.Context, key string, index interface{}, ttl time.Duration) (*redislock.Lock, error) {
	return LockBackoff(ctx, key, index, ttl, 100*time.Millisecond)
}

// LockBackoff 分布式锁 (指定生存周期, 设置重试时间)
func LockBackoff(ctx context.Context, key string, index interface{}, ttl time.Duration, backoff time.Duration) (*redislock.Lock, error) {
	opt := &redislock.Options{}
	if backoff > 0 {
		opt.RetryStrategy = redislock.LinearBackoff(backoff)
	}
	keyName := fmt.Sprintf("LOCK:%s:%s", key, util.ParseStr(index))
	return Locker().Obtain(ctx, keyName, ttl, opt)
}
