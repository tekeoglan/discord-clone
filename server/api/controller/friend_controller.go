package controller

import (
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendController struct {
	FriendService  model.FriendService
	SessionService model.SessionService
	AccountService model.AccountService
}

func (fc *FriendController) AddFriend(c *gin.Context) {
	var request model.FriendRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	var sessionId string
	sessionId, err = c.Cookie(service.COOKIE_PREFIX)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var senderHex string
	senderHex, err = fc.SessionService.RetriveSession(c, sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var senderId primitive.ObjectID
	senderId, err = primitive.ObjectIDFromHex(senderHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var receiver model.User
	receiver, err = fc.AccountService.GetByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	friend := &model.Friend{
		BaseModel: model.BaseModel{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Users:  []primitive.ObjectID{senderId, receiver.ID},
		Status: model.Pending,
	}

	err = fc.FriendService.Add(c, friend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Friend is created")
}

func (fc *FriendController) GetConfirmedFriends(c *gin.Context) {
	sessionId, err := c.Cookie(service.COOKIE_PREFIX)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var userIdHex string
	userIdHex, err = fc.SessionService.RetriveSession(c, sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	cursorPosString := c.Query("cursorPos")
	if cursorPosString == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid query value"})
		return
	}

	var cursorPos int
	cursorPos, err = strconv.Atoi(cursorPosString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	confirmed, err := fc.FriendService.GetConfirmed(c, userIdHex, cursorPos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, confirmed)
}

func (fc *FriendController) GetPendingFriends(c *gin.Context) {
	sessionId, err := c.Cookie(service.COOKIE_PREFIX)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var userIdHex string
	userIdHex, err = fc.SessionService.RetriveSession(c, sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	cursorPosString := c.Query("cursorPos")
	if cursorPosString == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid query value"})
		return
	}

	var cursorPos int
	cursorPos, err = strconv.Atoi(cursorPosString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	pending, err := fc.FriendService.GetPending(c, userIdHex, cursorPos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, pending)
}

func (fc *FriendController) RemoveFriend(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid query value."})
		return
	}

	err := fc.FriendService.Remove(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Friend is removed.")
}
