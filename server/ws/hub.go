package ws

import (
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/redis"
)

var hubInstance *Hub = nil

type Hub struct {
	clients        map[*Client]bool
	register       chan *Client
	unregister     chan *Client
	broadcast      chan []byte
	rooms          map[*Room]bool
	channelService model.ChannelService
	accountService model.AccountService
	sessionService model.SessionService
	redisClient    redis.Client
}

func NewWebsocketHub(cs model.ChannelService, as model.AccountService, ss model.SessionService, rc redis.Client) *Hub {
	if hubInstance != nil {
		return hubInstance
	}

	hubInstance = &Hub{
		clients:        make(map[*Client]bool),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		broadcast:      make(chan []byte),
		rooms:          make(map[*Room]bool),
		sessionService: ss,
		channelService: cs,
		accountService: as,
		redisClient:    rc,
	}

	return hubInstance
}

func GetHub() *Hub {
	return hubInstance
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.registerClient(client)
		case client := <-hub.unregister:
			hub.unregisterClient(client)
		case message := <-hub.broadcast:
			hub.broadcastToClients(message)
		}
	}
}

func (hub *Hub) registerClient(client *Client) {
	hub.clients[client] = true
}

func (hub *Hub) unregisterClient(client *Client) {
	delete(hub.clients, client)
}

func (hub *Hub) broadcastToClients(message []byte) {
	for client := range hub.clients {
		client.send <- message
	}
}

func (hub *Hub) BroadCastToRoom(message []byte, roomId string) {
	if room := hub.findRoomById(roomId); room != nil {
		room.publishRoomMessage(message)
	}
}

func (hub *Hub) findRoomById(id string) *Room {
	var foundRoom *Room
	for room := range hub.rooms {
		if room.GetId() == id {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

func (hub *Hub) createRoom(id string) *Room {
	room := NewRoom(id, hub.redisClient)
	go room.RunRoom()
	hub.rooms[room] = true

	return room
}
