package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func NewCache(redisAddr string, redisPassword string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
	})

	return &Cache{
		Client: rdb,
	}
}

func (c *Cache) Get(key string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *Cache) Set(key string, value float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val := strconv.FormatFloat(value, 'f', -1, 64)

	return c.Client.Set(ctx, key, val, 2*time.Hour).Err()
}
