package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	Ping(context.Context) (string, error)
	Get(context.Context, string) *redis.StringCmd
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
	Publish(context.Context, string, interface{}) *redis.IntCmd
	Subscribe(context.Context, ...string) *redis.PubSub
	Close() error
}

func NewClient(opt Options) Client {

	cl := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Username: opt.User,
		Password: opt.Password,
		DB:       opt.DB,
	})

	return &redisClient{cl: cl}
}

type Options struct {
	Addr     string
	User     string
	Password string
	DB       int
}

type redisClient struct {
	cl *redis.Client
}

func (rc *redisClient) Ping(ctx context.Context) (string, error) {
	return rc.cl.Ping(ctx).Result()
}

func (rc *redisClient) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) *redis.StatusCmd {
	return rc.cl.Set(ctx, key, val, expiration)
}

func (rc *redisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return rc.cl.Get(ctx, key)
}

func (rc *redisClient) Close() error {
	return rc.cl.Close()
}

func (rc *redisClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	return rc.cl.Publish(ctx, channel, message)
}

func (rc *redisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return rc.cl.Subscribe(ctx, channels...)
}
