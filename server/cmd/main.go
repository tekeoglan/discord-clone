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

	if env.AppEnv == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := gin.Default()

	route.Setup(env, timeout, db, cache, ginEngine)

	ginEngine.Run(env.ServerAddress)
}
