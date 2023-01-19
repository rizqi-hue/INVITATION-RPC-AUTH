package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val string, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Incr(key string) ([]byte, error)
}
