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

func NewFriendRoute(db mongo.Database, cl redis.Client, group *gin.RouterGroup) {

	fr := repository.NewFriendRepository(db, model.CollectionFriend)

	cr := repository.NewCacheRepository(cl)

	ar := repository.NewUserRepository(db, model.CollectionUser)

	chr := repository.NewChannelRepository(db, model.CollectionChannel)

	hub := ws.GetHub()

	fc := controller.FriendController{
		FriendService:  service.NewFriendService(fr),
		SessionService: service.NewSessionService(cr),
		AccountService: service.NewAccountService(ar),
		SocketService:  service.NewSocketService(hub, chr),
	}

	group.POST("/add", fc.AddFriend)
	group.POST("/confirm/:friendId", fc.AcceptFriendRequest)
	group.GET("/getConfirmed", fc.GetConfirmedFriends)
	group.GET("/getPending", fc.GetPendingFriends)
	group.POST("/removeByUserIds", fc.RemoveByUserIds)
	group.POST("/remove", fc.RemoveFriend)
}
