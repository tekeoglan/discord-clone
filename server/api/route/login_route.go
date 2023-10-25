package route

import (
	"github/tekeoglan/discord-clone/api/controller"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/redis"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(db mongo.Database,
	cl redis.Client, group *gin.RouterGroup) {

	ur := repository.NewUserRepository(db, model.CollectionUser)

	cr := repository.NewCacheRepository(cl)

	lr := controller.LoginController{
		AccountService: service.NewAccountService(ur),
		SessionService: service.NewSessionService(cr),
	}

	group.POST("/login", lr.Login)
	group.GET("/fetchUser", lr.FetchUser)
}
