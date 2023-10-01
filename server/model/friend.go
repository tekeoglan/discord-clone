package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionFriend = "friends"

const (
	Confirmed string = "confirmed"
	Pending          = "pending"
)

type FriendResult struct {
	BaseModel  `bson:",inline"`
	FriendInfo User   `bson:"friendInfo"`
	Status     string `bson:"status"`
}

type Friend struct {
	BaseModel `bson:",inline"`
	Users     []primitive.ObjectID `bson:"users"`
	Status    string               `bson:"status"`
}

type FriendRequest struct {
	Email string `bson:"email"`
}

type FriendRepository interface {
	Add(context.Context, *Friend) error
	GetConfirmed(context.Context, string, int) ([]FriendResult, error)
	GetPending(context.Context, string, int) ([]FriendResult, error)
	Remove(context.Context, string) error
	IsFriends(context.Context, primitive.ObjectID, primitive.ObjectID) (bool, error)
}

type FriendService interface {
	Add(context.Context, *Friend) error
	GetConfirmed(context.Context, string, int) ([]FriendResult, error)
	GetPending(context.Context, string, int) ([]FriendResult, error)
	Remove(context.Context, string) error
}
