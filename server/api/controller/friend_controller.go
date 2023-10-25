package controller

import (
	"github/tekeoglan/discord-clone/model"
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
	SocketService  model.SocketService
}

func (fc *FriendController) AddFriend(c *gin.Context) {
	var request model.FriendRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	var sessionId string
	sessionId, err = c.Cookie(model.COOKIE_PREFIX_SESSION)
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

	var sender model.User
	sender, err = fc.AccountService.FetchUser(c, senderHex)
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

	fc.SocketService.EmitAddFriendRequest(receiver.ID.Hex(), &model.FriendRequestWs{
		Id:       sender.ID.Hex(),
		UserName: sender.UserName,
		Image:    sender.Image,
		Type:     model.Incoming,
	})

	c.JSON(http.StatusOK, "Friend is created")
}

func (fc *FriendController) AcceptFriendRequest(c *gin.Context) {
	sessionId, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
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

	friendId := c.Param("friendId")
	if friendId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid param value"})
		return
	}

	var friendResult model.FriendGetResult
	friendResult, err = fc.FriendService.AcceptFriend(c, userIdHex, friendId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	fc.SocketService.EmitAddFriend(&friendResult.FriendInfos[0], &friendResult.FriendInfos[1])

	c.JSON(http.StatusOK, "Friend is accepted")
}

func (fc *FriendController) GetConfirmedFriends(c *gin.Context) {
	sessionId, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
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
	sessionId, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
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

	friend, err := fc.FriendService.GetFriend(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FriendService.Remove(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	fc.SocketService.EmitRemoveFriend(friend.Users[0].Hex(), friend.Users[1].Hex())

	c.JSON(http.StatusOK, "Friend is removed.")
}
