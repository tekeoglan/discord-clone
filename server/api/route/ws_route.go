package route

import (
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/redis"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"
	"github/tekeoglan/discord-clone/ws"

	"github.com/gin-gonic/gin"
)

func NewWsRoute(db mongo.Database, rc redis.Client, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, model.CollectionUser)

	cr := repository.NewChannelRepository(db, model.CollectionChannel)

	cacr := repository.NewCacheRepository(rc)

	as := service.NewAccountService(ur)

	cs := service.NewChannelService(cr)

	ss := service.NewSessionService(cacr)

	hub := ws.NewWebsocketHub(cs, as, ss, rc)
	go hub.Run()

	group.GET("", func(c *gin.Context) {
		ws.ServeWs(hub, c)
	})
}
