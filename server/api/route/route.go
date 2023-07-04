package route

import (
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewRegisterRouter(env, timeout, db, publicRouter)
}
