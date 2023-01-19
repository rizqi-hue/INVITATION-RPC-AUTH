package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/INVITATION-RPC-AUTH/pkg/cache"
)

type AuthRedis interface {
	GetAuthRedis(ctx context.Context, key string) (string, error)
	SetAuthRedis(ctx context.Context, key string, value string) error
	DeleteAuthRedis(ctx context.Context, key string) error
}

// List of keys for user's caching
const (
	userKey string = "user:%v"
)

type authRedis struct {
	redis cache.Cache
}

// NewUserRedis returns new instance of userRedis.
func NewAuthRedis(ctx context.Context) *authRedis {
	cache, err := cache.NewRedis(ctx)
	if err != nil {
		log.Fatal("failed getting auth redis cache functions")
		return nil
	}

	return &authRedis{redis: cache}
}

func (ur *authRedis) GetAuthRedis(ctx context.Context, key string) (string, error) {
	val, err := ur.redis.Get(ctx, fmt.Sprintf(userKey, key))
	return val, err
}

func (ur *authRedis) SetAuthRedis(ctx context.Context, key string, value string) error {
	err := ur.redis.Set(ctx, fmt.Sprintf(userKey, key), value, 168*time.Hour)
	return err
}

func (ur *authRedis) DeleteAuthRedis(ctx context.Context, key string) error {
	err := ur.redis.Delete(ctx, fmt.Sprintf(userKey, key))
	return err
}
