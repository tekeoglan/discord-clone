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

type Friend struct {
	BaseModel
	Users  []primitive.ObjectID `bson:"users"`
	Status string               `bson:"status"`
}

type FriendRequest struct {
	Email string `bson:"email"`
}

type FriendRepository interface {
	Add(context.Context, *Friend) error
	GetConfirmed(context.Context, string, int) ([]interface{}, error)
	GetPending(context.Context, string, int) ([]interface{}, error)
	Remove(context.Context, string) error
}

type FriendService interface {
	Add(context.Context, *Friend) error
	GetConfirmed(context.Context, string, int) ([]interface{}, error)
	GetPending(context.Context, string, int) ([]interface{}, error)
	Remove(context.Context, string) error
}
