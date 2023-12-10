package controller

import (
	"github/tekeoglan/discord-clone/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	ChannelService model.ChannelService
	SessionService model.SessionService
	SocketService  model.SocketService
	AccountService model.AccountService
}

func (cc *ChannelController) CreateFc(c *gin.Context) {
	userId := c.MustGet(model.CONTEXT_USER_KEY).(string)

	friendId := c.Query("friendId")
	if friendId == "" {
		c.JSON(http.StatusBadRequest, "invalid query")
		return
	}

	fc, err := cc.ChannelService.CreateFriendChannel(c, userId, friendId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var user model.User
	user, err = cc.AccountService.FetchUser(c, userId)
	if err == nil {
		cc.SocketService.EmitNewChannel(friendId, &model.ChannelResponseWs{
			Id:              fc.ChannelID,
			Name:            user.UserName,
			CreatedAt:       fc.CreatedAt,
			UpdatedAt:       fc.UpdatedAt,
			HasNotification: false,
		})
	}

	var friend model.User
	friend, err = cc.AccountService.FetchUser(c, friendId)
	if err == nil {
		cc.SocketService.EmitNewChannel(userId, &model.ChannelResponseWs{
			Id:              fc.ChannelID,
			Name:            friend.UserName,
			CreatedAt:       fc.CreatedAt,
			UpdatedAt:       fc.UpdatedAt,
			HasNotification: false,
		})
	}

	c.JSON(http.StatusOK, fc)
}

func (cc *ChannelController) GetFcById(c *gin.Context) {
	id, _ := c.Params.Get("channelId")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid param"})
		return
	}

	channel, err := cc.ChannelService.GetFcById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, channel)
}

func (cc *ChannelController) GetFcByUserIds(c *gin.Context) {
	userId := c.MustGet(model.CONTEXT_USER_KEY).(string)

	friendId := c.Query("friendId")
	if friendId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid query"})
		return
	}

	channel, err := cc.ChannelService.GetFcByUserIds(c, userId, friendId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, channel)
}

func (cc *ChannelController) GetFcs(c *gin.Context) {

	userId := c.MustGet(model.CONTEXT_USER_KEY).(string)

	channels, err := cc.ChannelService.GetFriendChannels(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, channels)
}

func (cc *ChannelController) Delete(c *gin.Context) {
	channelId, _ := c.Params.Get("channelId")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid param"})
		return
	}

	err := cc.ChannelService.DeleteChannel(c, channelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "channel has been removed")
}
