package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionChannel = "channels"

const (
	ChannelTypeFriend = "friend"
	ChannelTypeGuild  = "guild"
)

type Channel struct {
	BaseModel `bson:",inline"`
	ChannelID string               `bson:"channelId"`
	Messages  []primitive.ObjectID `bson:"messages"`
	Type      string               `bson:"type"`
}

type FriendChannel struct {
	Channel `bson:",inline"`
	Users   []primitive.ObjectID `bson:"users"`
}

type GuildChannel struct {
	Channel `bson:",inline"`
	Name    string `bson:"name"`
}

type FriendChannelResult struct {
	Channel    `bson:",inline"`
	FriendInfo User `bson:"friendInfo"`
}

type ChannelResponseWs struct {
	Id              string    `json:"id"`
	Name            string    `json:"name,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	HasNotification bool      `json:"hasNotification"`
}

type ChannelService interface {
	CreateFriendChannel(context.Context, string, string) (*FriendChannel, error)
	GetFcById(context.Context, string) (*FriendChannel, error)
	GetFcByUserIds(context.Context, string, string) (*FriendChannel, error)
	GetFriendChannels(context.Context, string) (*[]FriendChannelResult, error)
	DeleteChannel(context.Context, string) error
}

type ChannelRepository interface {
	CreateChannel(context.Context, interface{}) error
	GetFcById(context.Context, string) (*FriendChannel, error)
	GetFcByUserIds(context.Context, string, string) (*FriendChannel, error)
	GetFriendChannels(context.Context, string) (*[]FriendChannelResult, error)
	DeleteChannel(context.Context, string) error
}
