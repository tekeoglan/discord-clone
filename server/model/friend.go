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
	BaseModel `bson:",inline"`
	Users     []primitive.ObjectID `bson:"users"`
	Status    string               `bson:"status"`
}

type FriendGetResult struct {
	BaseModel   `bson:",inline"`
	Users       []primitive.ObjectID `bson:"users"`
	FriendInfos []User               `bson:"friendInfos"`
	Status      string               `bson:"status"`
}

type FriendAggragateResult struct {
	BaseModel  `bson:",inline"`
	FriendInfo User   `bson:"friendInfo"`
	Status     string `bson:"status"`
}

type FriendGetAllResult struct {
	Friends   []FriendAggragateResult
	CursorPos int
}

type FriendRequest struct {
	Email string `json:"email"`
}

type RequestType int

const (
	Outgoing RequestType = iota
	Incoming
)

type FriendRequestWs struct {
	Id       string      `json:"id"`
	UserName string      `json:"userName"`
	Image    string      `json:"image"`
	Type     RequestType `json:"type" enums:"0,1"`
}

type FriendRepository interface {
	Add(context.Context, *Friend) error
	Update(context.Context, string, interface{}) error
	Get(context.Context, string) (FriendGetResult, error)
	GetConfirmed(context.Context, string, int) (FriendGetAllResult, error)
	GetPending(context.Context, string, int) (FriendGetAllResult, error)
	Remove(context.Context, string) error
	IsFriends(context.Context, primitive.ObjectID, primitive.ObjectID) (bool, error)
}

type FriendService interface {
	Add(context.Context, *Friend) error
	AcceptFriend(context.Context, string, string) (FriendGetResult, error)
	GetFriend(context.Context, string) (FriendGetResult, error)
	GetConfirmed(context.Context, string, int) (FriendGetAllResult, error)
	GetPending(context.Context, string, int) (FriendGetAllResult, error)
	Remove(context.Context, string) error
}
