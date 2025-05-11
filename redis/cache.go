package redis

import (
	"context"
	"time"

	gredis "github.com/redis/go-redis/v9"
)

// Cache defines basic cache operations
type Cache interface {
	//Set key-value pair in cache
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error

	//Get value from cache by key
	Get(ctx context.Context, key string) (interface{}, error)

	//Delete key from cache
	Delete(ctx context.Context, key string) error

	//Exists checks if key exists in cache
	Exists(ctx context.Context, key string) (bool, error)
}

type cacheManager struct {
	client *gredis.Client
}

// New Cache returns a cache key backed by the singleton redis client
func NewCache() Cache {
	return &cacheManager{
		client: Client(),
	}
}
func (c *cacheManager) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}
func (c *cacheManager) Get(ctx context.Context, key string) (interface{}, error) {
	return c.client.Get(ctx, key).Result()
}
func (c *cacheManager) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
func (c *cacheManager) Exists(ctx context.Context, key string) (bool, error) {
	val, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}
