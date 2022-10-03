package redis

import (
	"github.com/go-redis/redis/v9"
	"sync"
)

var (
	once     sync.Once
	instance *redis.Client
)

func RedisClient() *redis.Client {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		instance = rdb
	})
	return instance
}
