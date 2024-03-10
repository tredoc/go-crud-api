package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type RCache interface {
	Set(string, interface{}, time.Duration) error
	Get(string) (string, error)
	Invalidate(string)
}

type Cache struct {
	Redis RCache
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{
		Redis: NewRCache(rdb),
	}
}
