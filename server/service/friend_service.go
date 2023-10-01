package service

import (
	"context"
	"errors"
	"github/tekeoglan/discord-clone/model"
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

func (fs *friendService) GetConfirmed(c context.Context, id string, cursorPos int) ([]model.FriendResult, error) {
	return fs.friendRepository.GetConfirmed(c, id, cursorPos)
}

func (fs *friendService) GetPending(c context.Context, id string, cursorPos int) ([]model.FriendResult, error) {
	return fs.friendRepository.GetPending(c, id, cursorPos)
}

func (fs *friendService) Remove(c context.Context, id string) error {
	return fs.friendRepository.Remove(c, id)
}
