package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository() (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	return &RedisRepository{
		client: client,
	}, nil
}

func (r *RedisRepository) GetFileSize(filename string) (int64, error) {
	ctx := context.Background()
	size, err := r.client.Get(ctx, filename).Int64()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	return size, nil
}

func (r *RedisRepository) SaveFileMapping(filename string, size int64) error {
	ctx := context.Background()
	return r.client.Set(ctx, filename, size, 0).Err()
}

func (r *RedisRepository) Close() error {
	return r.client.Close()
}

func (r *RedisRepository) Set(fileName string, fileSize int64) error {
	ctx := context.TODO()
	return r.client.Set(ctx, fileName, fileSize, 0).Err()
}
