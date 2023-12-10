package controller

import (
	"github/tekeoglan/discord-clone/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	MessageService model.MessageService
	SocketService  model.SocketService
}

func (mc *MessageController) PostMessage(c *gin.Context) {
	var request model.MessageMeta
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	message := &model.Message{
		BaseModel: model.BaseModel{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		MessageMeta: model.MessageMeta{
			UserID:    request.UserID,
			ChannelID: request.ChannelID,
			UserName:  request.UserName,
			Text:      request.Text,
		},
	}

	message, err = mc.MessageService.CreateMessage(c, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	mc.SocketService.EmitNewMessage(message.ChannelID, message)

	c.JSON(http.StatusOK, "message has been posted")
}

func (mc *MessageController) UpdateMessage(c *gin.Context) {
	var request model.UpdateMessageRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	err = mc.MessageService.UpdateMessage(c, request.ID, request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	message, _ := mc.MessageService.GetMessage(c, request.ID)
	if message != nil {
		mc.SocketService.EmitEditMessage(message.ChannelID, message)
	}

	c.JSON(http.StatusOK, "message updated")
}

func (mc *MessageController) DeleteMessage(c *gin.Context) {
	channelId := c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid query"})
		return
	}

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid query"})
		return
	}

	err := mc.MessageService.DeleteMessage(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	mc.SocketService.EmitDeleteMessage(channelId, id)

	c.JSON(http.StatusOK, "message deleted")
}

func (mc *MessageController) GetMessage(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid query"})
		return
	}

	message, err := mc.MessageService.GetMessage(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (mc *MessageController) GetChannelMessages(c *gin.Context) {
	channelId := c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid query"})
		return
	}

	cursorPosString := c.Query("cursorPos")
	if cursorPosString == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid query"})
		return
	}

	cursorPos, err := strconv.Atoi(cursorPosString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var result model.MessageGetAllResult
	result, err = mc.MessageService.GetChannelMessages(c, channelId, cursorPos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
