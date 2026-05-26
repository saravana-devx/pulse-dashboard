package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"

	"pulseDashboard/internal/config"
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
		c := config.Get()
		rdb := redis.NewClient(&redis.Options{
			Addr:     c.RedisAddr,
			Password: c.RedisPassword,
			DB:       0,
		})
		redisInstance = &Redis{Client: rdb}
	})
	return redisInstance
}

