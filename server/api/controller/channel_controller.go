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
}

func (cc *ChannelController) CreateFc(c *gin.Context) {
	session, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var userId string
	userId, err = cc.SessionService.RetriveSession(c, session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

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

	crws := model.ChannelResponseWs{
		Id:              fc.ChannelID,
		CreatedAt:       fc.CreatedAt,
		UpdatedAt:       fc.UpdatedAt,
		HasNotification: false,
	}

	cc.SocketService.EmitNewChannel(fc.ChannelID, &crws)

	c.JSON(http.StatusOK, "channel has been created")
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
	session, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var userId string
	userId, err = cc.SessionService.RetriveSession(c, session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

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
	session, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var userId string
	userId, err = cc.SessionService.RetriveSession(c, session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

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
