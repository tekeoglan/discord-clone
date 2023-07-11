package route

import (
	"github/tekeoglan/discord-clone/api/middleware"
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	publicRouter.Use(middleware.Timeout(timeout))

	NewRegisterRouter(env, timeout, db, publicRouter)
}
