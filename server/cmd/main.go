package main

import (
	"context"
	"fmt"
	"github/tekeoglan/discord-clone/api/route"
	"github/tekeoglan/discord-clone/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	ipPort := fmt.Sprintf("%s:%s", env.CacheHost, env.CachePort)
	cache := app.Redis
	defer cache.ClientKill(context.TODO(), ipPort)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
