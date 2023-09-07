package repository

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/redis"
	"time"
)

type cacheRepository struct {
	client redis.Client
}

func NewCacheRepository(cl redis.Client) model.CacheRepository {
	return &cacheRepository{
		client: cl,
	}
}

func (cr *cacheRepository) Set(c context.Context, key string, value interface{}, expiration time.Duration) error {
	cmd := cr.client.Set(c, key, value, expiration)
	return cmd.Err()
}

func (cr *cacheRepository) Get(c context.Context, key string) (string, error) {
	cmd := cr.client.Get(c, key)
	return cmd.Result()
}
