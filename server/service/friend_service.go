package service

import (
	"context"
	"errors"
	"github/tekeoglan/discord-clone/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type friendService struct {
	friendRepository model.FriendRepository
}

func NewFriendService(friendRepository model.FriendRepository) model.FriendService {
	return &friendService{
		friendRepository: friendRepository,
	}
}

func (fs *friendService) Add(c context.Context, friend *model.Friend) error {
	isFriends, err := fs.friendRepository.IsFriends(c, friend.Users[0], friend.Users[1])
	if err != nil {
		return err
	}

	if isFriends {
		return errors.New("Users are already befriended.")
	}

	return fs.friendRepository.Add(c, friend)
}

func (fs *friendService) AcceptFriend(c context.Context, userId string, friendId string) (model.FriendGetResult, error) {
	friend, err := fs.friendRepository.Get(c, friendId)
	if err != nil {
		return friend, err
	}

	if friend.Users[0].Hex() == userId {
		return friend, errors.New("cannot accept the friend request created by yourself")
	}

	update := bson.M{"$set": bson.M{"updatedAt": time.Now(), "status": model.Confirmed}}

	err = fs.friendRepository.Update(c, friendId, update)

	return friend, err
}

func (fs *friendService) GetFriend(c context.Context, id string) (model.FriendGetResult, error) {
	return fs.friendRepository.Get(c, id)
}

func (fs *friendService) GetConfirmed(c context.Context, id string, cursorPos int) (model.FriendGetAllResult, error) {
	return fs.friendRepository.GetConfirmed(c, id, cursorPos)
}

func (fs *friendService) GetPending(c context.Context, id string, cursorPos int) (model.FriendGetAllResult, error) {
	return fs.friendRepository.GetPending(c, id, cursorPos)
}

func (fs *friendService) RemoveByUserIds(c context.Context, requesterId, userId string) error {
	friend, err := fs.friendRepository.GetByUserIds(c, requesterId, userId)
	if err != nil {
		return err
	}

	return fs.friendRepository.Remove(c, friend.ID.Hex())
}

func (fs *friendService) Remove(c context.Context, id string) error {
	return fs.friendRepository.Remove(c, id)
}
