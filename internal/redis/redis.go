package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"sync"
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

