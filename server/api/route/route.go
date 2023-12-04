package route

import (
	"github/tekeoglan/discord-clone/api/middleware"
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/redis"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database,
	rc redis.Client, gin *gin.Engine) {

	cr := repository.NewCacheRepository(rc)
	ss := service.NewSessionService(cr)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{env.ClientAddress}
	corsConfig.AllowCredentials = true

	gin.Use(cors.New(corsConfig))

	wsRouter := gin.Group("/ws")
	wsRouter.Use(middleware.Auth(ss))

	gin.Use(middleware.Timeout(timeout))

	publicRouter := gin.Group("")

	friendRouter := gin.Group("/friend")
	friendRouter.Use(middleware.Auth(ss))

	messageRouter := gin.Group("/message")
	messageRouter.Use(middleware.Auth(ss))

	channelRouter := gin.Group("/channel")
	channelRouter.Use(middleware.Auth(ss))

	NewWsRoute(db, rc, wsRouter)
	NewRegisterRouter(db, rc, publicRouter)
	NewLoginRoute(db, rc, publicRouter)
	NewFriendRoute(db, rc, friendRouter)
	NewMessageRouter(db, messageRouter)
	NewChannelRoute(db, rc, channelRouter)
}
