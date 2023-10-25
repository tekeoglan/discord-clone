package route

import (
	"github/tekeoglan/discord-clone/api/controller"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/redis"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"
	"github/tekeoglan/discord-clone/ws"

	"github.com/gin-gonic/gin"
)

func NewChannelRoute(db mongo.Database, rc redis.Client, group *gin.RouterGroup) {
	cr := repository.NewChannelRepository(db, model.CollectionChannel)

	cacr := repository.NewCacheRepository(rc)

	hub := ws.GetHub()

	cc := controller.ChannelController{
		ChannelService: service.NewChannelService(cr),
		SessionService: service.NewSessionService(cacr),
		SocketService:  service.NewSocketService(hub, cr),
	}

	group.POST("/fc", cc.CreateFc)
	group.GET("/fc/:channelId", cc.GetFcById)
	group.GET("/fc/byUserIds", cc.GetFcByUserIds)
	group.GET("/fc/user", cc.GetFcs)
	group.DELETE("/fc/:channelId", cc.Delete)
}
