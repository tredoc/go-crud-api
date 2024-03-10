package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/tredoc/go-crud-api/pkg/log"
	"time"
)

var EXPIRATION = time.Hour * 1

type Redis struct {
	rdb *redis.Client
}

func NewRCache(rdb *redis.Client) *Redis {
	return &Redis{
		rdb: rdb,
	}
}

func (c *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	err := c.rdb.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Redis) Get(key string) (string, error) {
	val, err := c.rdb.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrNotFound
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Redis) Invalidate(key string) {
	err := c.rdb.Del(context.Background(), key).Err()
	if err != nil {
		log.Error(err.Error())
	}
}
