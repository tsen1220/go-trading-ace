package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRedisHelper interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

type RedisHelper struct {
	RedisClient *redis.Client
}

func NewRedisHelper(redisClient *redis.Client) IRedisHelper {
	return &RedisHelper{
		RedisClient: redisClient,
	}
}

func (r *RedisHelper) Set(key string, value string, expiration time.Duration) error {
	err := r.RedisClient.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s: %w", key, err)
	}

	return nil
}

func (r *RedisHelper) Get(key string) (string, error) {
	val, err := r.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist", key)
		}

		return "", fmt.Errorf("failed to get key %s: %w", key, err)
	}

	return val, nil
}

func (r *RedisHelper) Delete(key string) error {
	err := r.RedisClient.Del(context.Background(), key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key %s: %w", key, err)
	}

	return nil
}
