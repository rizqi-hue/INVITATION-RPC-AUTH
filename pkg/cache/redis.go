package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/INVITATION-RPC-AUTH/pkg/config"
	"github.com/INVITATION-RPC-AUTH/pkg/logger"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

type redisCache struct {
	client redis.Client
	logger *logrus.Entry
}

func NewRedis(ctx context.Context) (Cache, error) {
	log := logger.GetLoggerContext(ctx, "cache", "NewRedis")

	jsonByte, err := json.Marshal(config.Get("redis"))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var cfg RedisConfig
	err = json.Unmarshal(jsonByte, &cfg)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//Default Pool Size
	poolSize := 10

	if cfg.PoolSize != 0 {
		poolSize = cfg.PoolSize
	}

	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Server,
		Password: cfg.AuthPass, // no password set
		DB:       0,            // use default DB
		PoolSize: poolSize,
	})

	pong, err := conn.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed creating new redis pool [%v]", err)
		return nil, err
	}

	fmt.Println(pong, err)
	log.Info("Created new connection to redis")

	c := &redisCache{
		client: *conn,
		logger: logger.GetLogger("cache", "redisFunc"),
	}

	return c, err
}

// Get the item with the provided key.
// Return nil byte if the item didn't already exist in the cache.
func (m *redisCache) Get(key string) (rcv []byte, err error) {
	// err = m.client.Do(radix.Cmd(&rcv, "GET", key))
	// if err != nil {
	// 	m.logger.Error(fmt.Sprintf("%s %s %s", key, string(rcv), err.Error()))
	// 	return
	// }
	return
}

// Set writes the given item, unconditionally.
func (m *redisCache) Set(key string, val []byte, expiration time.Duration) (err error) {

	// args := []string{key, string(val)}

	// if expiration != 0 {
	// 	//EX seconds -- Set the specified expire time, in seconds.
	// 	//PX milliseconds -- Set the specified expire time, in milliseconds.
	// 	args = append(args, "EX", fmt.Sprintf("%d", int(expiration.Seconds())))
	// }

	// err = m.client.Do(radix.Cmd(nil, "SET", args...))
	// if err != nil {
	// 	m.logger.Error(fmt.Sprintf("%s %s %s", key, string(val), err.Error()))
	// 	return
	// }

	return
}

// Delete deletes the item with the provided key.
// return nil error if the item didn't already exist in the cache.
func (m *redisCache) Delete(key string) (err error) {
	// err = m.client.Do(radix.Cmd(nil, "DEL", key))
	// if err != nil {
	// 	m.logger.Error(fmt.Sprintf("%s %s", key, err.Error()))
	// 	return
	// }
	return
}

// Incr the item with the provided key.
// Return incremented byte if the item didn't already exist in the cache.
func (m *redisCache) Incr(key string) (rcv []byte, err error) {
	// err = m.client.Do(radix.Cmd(&rcv, "INCR", key))
	// if err != nil {
	// 	m.logger.Error(fmt.Sprintf("%s %s %s", key, string(rcv), err.Error()))
	// 	return
	// }
	return
}
