package route

import (
	"github/tekeoglan/discord-clone/api/controller"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"
	"github/tekeoglan/discord-clone/repository"
	"github/tekeoglan/discord-clone/service"
	"github/tekeoglan/discord-clone/ws"

	"github.com/gin-gonic/gin"
)

func NewMessageRouter(db mongo.Database, group *gin.RouterGroup) {
	mr := repository.NewMessageRepository(db, model.CollectionMessage)

	cr := repository.NewChannelRepository(db, model.CollectionChannel)

	hub := ws.GetHub()

	mc := controller.MessageController{
		MessageService: service.NewMessageService(mr),
		SocketService:  service.NewSocketService(hub, cr),
	}

	group.GET("", mc.GetMessage)
	group.POST("/add", mc.PostMessage)
	group.PATCH("/update", mc.UpdateMessage)
	group.DELETE("/delete", mc.DeleteMessage)
	group.GET("/channel", mc.GetChannelMessages)
}
