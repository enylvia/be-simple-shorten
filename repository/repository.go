package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Set(key, val string) (string, error)
	Get(key string) (string, error)
}

type RedisRepositoryImplement struct {
	rd *redis.Client
}

func NewRedisRepository(rd *redis.Client) RedisRepository {
	return &RedisRepositoryImplement{rd}
}

func (r *RedisRepositoryImplement) Set(key, val string) (string, error) {
	err := r.rd.Set(context.Background(), key, val, 10*time.Minute).Err()
	if err != nil {
		return "failed", err
	}
	return "success", nil

}

func (r *RedisRepositoryImplement) Get(key string) (string, error) {
	val, err := r.rd.Get(context.Background(), key).Result()
	if err != nil {
		return "data with this " + key + "is not found", err
	}
	return val, nil
}
