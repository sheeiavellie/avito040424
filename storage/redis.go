package storage

import (
	"context"
	"hash"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sheeiavellie/avito040424/data"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(
	redisAddr string,
	redisPass string,
) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &RedisStorage{
		client: client,
	}, nil
}

func (rs *RedisStorage) GetBanner(
	ctx context.Context,
	bannerKey hash.Hash32,
) (*data.Banner, error) {
	return nil, nil
}

func (rs *RedisStorage) SetBanner(
	ctx context.Context,
	banner *data.Banner,
) error {
	return nil
}
