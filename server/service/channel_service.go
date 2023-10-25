package service

import (
	"context"
	"errors"
	"github/tekeoglan/discord-clone/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type channelService struct {
	channelRepository model.ChannelRepository
}

func NewChannelService(cr model.ChannelRepository) model.ChannelService {
	return &channelService{
		channelRepository: cr,
	}
}

func (cs *channelService) CreateFriendChannel(c context.Context, userId string, friendId string) (*model.FriendChannel, error) {
	_, err := cs.channelRepository.GetFcByUserIds(c, userId, friendId)
	var friendChannel model.FriendChannel

	if err == nil {
		return &friendChannel, errors.New("Friend channel is already exist")
	}

	var userObjId primitive.ObjectID
	userObjId, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		return &friendChannel, err
	}

	var friendObjId primitive.ObjectID
	friendObjId, err = primitive.ObjectIDFromHex(friendId)
	if err != nil {
		return &friendChannel, err
	}

	friendChannel.BaseModel = model.BaseModel{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	friendChannel.ChannelID = GenerateId()
	friendChannel.Type = model.ChannelTypeFriend
	friendChannel.Users = []primitive.ObjectID{userObjId, friendObjId}

	err = cs.channelRepository.CreateChannel(c, &friendChannel)
	return &friendChannel, err
}

func (cs *channelService) GetFcById(c context.Context, id string) (*model.FriendChannel, error) {
	return cs.channelRepository.GetFcById(c, id)
}

func (cs *channelService) GetFcByUserIds(c context.Context, user1Id string, user2Id string) (*model.FriendChannel, error) {
	return cs.channelRepository.GetFcByUserIds(c, user1Id, user2Id)
}

func (cs *channelService) GetFriendChannels(c context.Context, userId string) (*[]model.FriendChannelResult, error) {
	return cs.channelRepository.GetFriendChannels(c, userId)
}

func (cs *channelService) DeleteChannel(c context.Context, id string) error {
	return cs.channelRepository.DeleteChannel(c, id)
}
