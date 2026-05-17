package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type Redis struct {
	Client *redis.Client
}

var (
	redisInstance *Redis
	once          sync.Once
)

// singleton pattern gurantees that this redis connection will create only once
// every goroutines / functions will gaurantee that they well get the same pointer reference
// without singleton pattern we might can get multiple redis connections
func NewRedis() *Redis {
	once.Do(func() {
		addr := viper.GetString("REDIS_ADDR")
		password := viper.GetString("REDIS_PASSWORD")
		if addr == "" {
			addr = "localhost:6379"
		}

		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		redisInstance = &Redis{Client: rdb}
	})
	return redisInstance
}

func (r *Redis) SetJTI(ctx context.Context, key, userID string, exp time.Time) error {
	return r.Client.Set(ctx, key, userID, time.Until(exp)).Err()
}

func (r *Redis) DelJTI(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *Redis) GetUserByJTI(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
