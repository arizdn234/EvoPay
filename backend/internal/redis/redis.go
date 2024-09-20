package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// redis
var RedisClient *redis.Client
var Ctx = context.Background()

func Initialize() {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0, // Default DB
	})

	// Ping Redis server to ensure connection is successful
	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

// Set sets a key-value pair in Redis with an expiration time
func Set(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(Ctx, key, value, expiration).Err()
}

// Get retrieves a value by key from Redis
func Get(key string) (string, error) {
	val, err := RedisClient.Get(Ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// Delete removes a key from Redis
func Delete(key string) error {
	return RedisClient.Del(Ctx, key).Err()
}
