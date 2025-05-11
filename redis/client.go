package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// client singleton holds the single redis.Client for the process
var (
	client *redis.Client
	once   sync.Once
)

type ClientOptions = redis.Options

// InitClient must be called once at startup to configure Redis.
// e.g. redis.InitClient(&redis.Options{Addr:"localhost:6379"})
func InitClient(options *redis.Options) {
	once.Do(func() {
		fmt.Println("Initializing Redis client...")
		c := redis.NewClient(options)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := c.Ping(ctx).Err(); err != nil {
			panic("Redis connection failed: " + err.Error())
		}
		client = c
	})
}

// Client returns the initialized redis client
func Client() *redis.Client {
	if client == nil {
		panic("Redis client not initialized. Please call InitClient first.")
	}
	return client
}

// Create a function to print the current redis client info
func GetClientInfo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	info, err := client.Info(ctx).Result()
	if err != nil {
		panic("Failed to get Redis info: " + err.Error())
	}
	fmt.Println("Redis Info: ", info)
}
