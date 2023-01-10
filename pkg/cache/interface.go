package cache

import (
	"time"
)

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, expiration time.Duration) error
	Delete(key string) error
	Incr(key string) ([]byte, error)
}
