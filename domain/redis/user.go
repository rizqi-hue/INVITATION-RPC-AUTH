package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/INVITATION-RPC-AUTH/pkg/cache"
)

type UserRedisRepository interface {
	GetUserRedis() ([]byte, error)
	SetUserRedis() error
}

// List of keys for user's caching
const (
	userKey string = "user:%d"
)

type userRedis struct {
	redis cache.Cache
}

// NewUserRedis returns new instance of userRedis.
func NewUserRedis(ctx context.Context) *userRedis {
	cache, err := cache.NewRedis(ctx)
	if err != nil {
		log.Fatal("failed getting redis cache functions")
		return nil
	}

	return &userRedis{redis: cache}
}

func (ur *userRedis) GetUserRedis() ([]byte, error) {
	val, err := ur.redis.Get(fmt.Sprintf(userKey, 12))
	return val, err
}

func (ur *userRedis) SetUserRedis() error {
	err := ur.redis.Set(fmt.Sprintf(userKey, 12), []byte("test"), 3*time.Second)
	return err
}
