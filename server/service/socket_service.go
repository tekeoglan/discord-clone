package service

import (
	"context"
	"encoding/json"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/ws"
	"log"
)

type socketService struct {
	Hub               *ws.Hub
	ChannelRepository model.ChannelRepository
}

func NewSocketService(hub *ws.Hub, cr model.ChannelRepository) model.SocketService {
	return &socketService{
		Hub:               hub,
		ChannelRepository: cr,
	}
}

func (ss *socketService) EmitNewMessage(room string, message *model.Message) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.NewMessageAction,
		Data:   message,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitEditMessage(room string, message *model.Message) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.EditMessageAction,
		Data:   message,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitDeleteMessage(room, messageId string) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.DeleteMessageAction,
		Data:   messageId,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitNewChannel(room string, channel *model.ChannelResponseWs) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.AddChannelAction,
		Data:   channel,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitNewFcChannel(room string, channel *model.ChannelResponseWs) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.AddChannelAction,
		Data:   channel,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitEditChannel(room string, channel *model.ChannelResponseWs) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.EditChannelAction,
		Data:   channel,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitNewDMNotification(c context.Context, channelId string, user *model.User) {

	response := model.DirectMessage{
		ChannelID: channelId,
		User:      user,
	}

	notification, err := json.Marshal(model.WebsocketMessage{
		Action: ws.NewDMNotificationAction,
		Data:   response,
	})

	if err != nil {
		log.Printf("error marshalling notification: %v\n", err)
	}

	var fc *model.FriendChannelInfosResult
	fc, err = ss.ChannelRepository.GetFcById(c, channelId)
	if err != nil {
		log.Printf("error getting the channel: %v\n", err)
	}

	for _, user := range fc.FriendInfo {
		if user.ID != user.ID {
			ss.Hub.BroadCastToRoom(notification, user.ID.Hex())
		}
	}
}

func (ss *socketService) EmitSendRequest(room string) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.SendRequestAction,
		Data:   "",
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitAddFriendRequest(room string, request *model.FriendRequestWs) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.AddRequestAction,
		Data:   request,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, room)
}

func (ss *socketService) EmitAddFriend(user, member *model.User) {

	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.AddFriendAction,
		Data:   user,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, member.ID.Hex())

	data, err = json.Marshal(model.WebsocketMessage{
		Action: ws.AddFriendAction,
		Data:   member,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, user.ID.Hex())
}

func (ss *socketService) EmitRemoveFriend(userId, memberId string) {
	data, err := json.Marshal(model.WebsocketMessage{
		Action: ws.RemoveFriendAction,
		Data:   memberId,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, userId)

	data, err = json.Marshal(model.WebsocketMessage{
		Action: ws.RemoveFriendAction,
		Data:   userId,
	})

	if err != nil {
		log.Printf("error marshalling response: %v\n", err)
	}

	ss.Hub.BroadCastToRoom(data, memberId)
}
