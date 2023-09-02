package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github/tekeoglan/discord-clone/redis"
)

func NewCache(env *Env) redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cacheHost := env.CacheHost
	cachePort := env.CachePort
	cacheUser := env.CacheUser
	cachePass := env.CachePort

	addr := fmt.Sprintf("%s:%s", cacheHost, cachePort)

	opt := redis.Options{
		Addr:     addr,
		User:     cacheUser,
		Password: cachePass,
		DB:       0,
	}

	client := redis.NewClient(opt)

	val, err := client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(val)
	return client
}

func KillCacheClient(env *Env, client redis.Client) {
	if client == nil {
		return
	}

	addr := fmt.Sprintf("%s:%s", env.CacheHost, env.CachePort)
	val := client.ClientKill(context.TODO(), addr)
	if err := val.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connectin to Redis closed.")
}
