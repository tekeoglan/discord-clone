package model

import "context"

type ReceivedMessage struct {
	Action  string       `json:"action"`
	Room    string       `json:"room"`
	Message *interface{} `json:"message"`
}

type WebsocketMessage struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type SocketService interface {
	EmitNewMessage(string, *Message)
	EmitEditMessage(string, *Message)
	EmitDeleteMessage(string, string)

	EmitNewChannel(string, *ChannelResponseWs)
	EmitEditChannel(string, *ChannelResponseWs)

	EmitNewDMNotification(context.Context, string, *User)

	EmitSendRequest(string)
	EmitAddFriendRequest(string, *FriendRequestWs)
	EmitAddFriend(*User, *User)
	EmitRemoveFriend(string, string)
}
