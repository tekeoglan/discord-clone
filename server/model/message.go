package model

import "context"

const CollectionMessage = "messages"

type MessageMeta struct {
	UserID    string `bson:"userId" form:"userId" binding:"required"`
	ChannelID string `bson:"channelId" form:"channelId" binding:"required"`
	Text      string `bson:"text" form:"text" binding:"required"`
}

type Message struct {
	BaseModel   `bson:",inline"`
	MessageMeta `bson:",inline"`
}

type DirectMessage struct {
	ChannelID string `json:"channelId"`
	User      *User  `json:",inline"`
}

type UpdateMessageRequest struct {
	ID   string `form:"id" binding:"required"`
	Text string `form:"text" binding:"required"`
}

type MessageService interface {
	CreateMessage(context.Context, *Message) (*Message, error)
	UpdateMessage(context.Context, string, string) error
	DeleteMessage(context.Context, string) error
	GetMessage(context.Context, string) (*Message, error)
	GetChannelMessages(context.Context, string, int) (*[]Message, error)
}

type MessageRepository interface {
	CreateMessage(context.Context, *Message) (*Message, error)
	UpdateMessage(context.Context, string, interface{}) error
	DeleteMessage(context.Context, string) error
	GetByID(context.Context, string) (*Message, error)
	GetChannelMessages(context.Context, string, int) (*[]Message, error)
}
