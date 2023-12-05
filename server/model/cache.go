package model

import (
	"context"
	"time"
)

type CacheRepository interface {
	Set(c context.Context, key string, value interface{}, expiration time.Duration) error
	Get(c context.Context, key string) (string, error)
	Delete(c context.Context, key ...string) (int64, error)
}
