package main

import (
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

	cache := app.Redis
	defer cache.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, cache, gin)

	gin.Run(env.ServerAddress)
}
