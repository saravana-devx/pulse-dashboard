package auth

import (
	"pulseDashboard/internal/redis"
)

type JTIStore struct {
	rdb *redis.Redis
}

func NewJTIStore(rdb *redis.Redis) *JTIStore {
	return &JTIStore{rdb: rdb}
}
