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

func NewRegisterRouter(db mongo.Database,
	cl redis.Client, group *gin.RouterGroup) {

	ur := repository.NewUserRepository(db, model.CollectionUser)

	rc := controller.RegisterController{
		AccountService: service.NewAccountService(ur),
	}

	group.POST("/register", rc.Register)
}
