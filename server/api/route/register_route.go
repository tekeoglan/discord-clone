package route

import (
	"github/tekeoglan/discord-clone/api/controller"
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRegisterRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, model.CollectionUser)
	rc := controller.RegisterController{
		RegisterService: service.NewRegisterService(ur),
		Env:             env,
	}

	group.POST("/register", rc.Register)
}
