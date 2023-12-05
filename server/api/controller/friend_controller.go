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
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	senderHex := c.MustGet(model.CONTEXT_USER_KEY).(string)

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
		Id:       friend.ID.Hex(),
		UserId:   sender.ID.Hex(),
		UserName: sender.UserName,
		Image:    sender.Image,
	})

	c.JSON(http.StatusOK, "Friend is created")
}

func (fc *FriendController) AcceptFriendRequest(c *gin.Context) {
	userIdHex := c.MustGet(model.CONTEXT_USER_KEY).(string)

	friendId := c.Param("friendId")
	if friendId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid param value"})
		return
	}

	friendResult, err := fc.FriendService.AcceptFriend(c, userIdHex, friendId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	fc.SocketService.EmitAddFriend(&friendResult.FriendInfos[0], &friendResult.FriendInfos[1])

	c.JSON(http.StatusOK, "Friend is accepted")
}

func (fc *FriendController) GetConfirmedFriends(c *gin.Context) {
	userIdHex := c.MustGet(model.CONTEXT_USER_KEY).(string)

	cursorPosString := c.Query("cursorPos")
	if cursorPosString == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid query value"})
		return
	}

	cursorPos, err := strconv.Atoi(cursorPosString)
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
	userIdHex := c.MustGet(model.CONTEXT_USER_KEY).(string)

	cursorPosString := c.Query("cursorPos")
	if cursorPosString == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid query value"})
		return
	}

	cursorPos, err := strconv.Atoi(cursorPosString)
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

func (fc *FriendController) RemoveByUserIds(c *gin.Context) {
	requesterId := c.MustGet(model.CONTEXT_USER_KEY).(string)

	userId := c.Query("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid query value."})
		return
	}

	err := fc.FriendService.RemoveByUserIds(c, requesterId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	fc.SocketService.EmitRemoveFriend(requesterId, userId)

	c.JSON(http.StatusOK, "Friend is removed.")
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
