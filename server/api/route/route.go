package route

import (
	"github/tekeoglan/discord-clone/api/middleware"
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/redis"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database,
	rc redis.Client, gin *gin.Engine) {

	cr := repository.NewCacheRepository(rc)
	ss := service.NewSessionService(cr)

	publicRouter := gin.Group("")
	publicRouter.Use(middleware.Timeout(timeout))

	wsRouter := gin.Group("/ws")
	wsRouter.Use(middleware.Auth(ss))

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
