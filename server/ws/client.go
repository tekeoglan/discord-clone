package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github/tekeoglan/discord-clone/model"
)

const (
	writeWait      = 5 * time.Second
	pongWait       = 45 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 10000
)

var newLine = []byte{'\n'}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ID    string
	conn  *websocket.Conn
	hub   *Hub
	send  chan []byte
	rooms map[*Room]bool
}

func newClient(conn *websocket.Conn, hub *Hub, userId string) *Client {
	return &Client{
		ID:    userId,
		conn:  conn,
		hub:   hub,
		send:  make(chan []byte, 256),
		rooms: make(map[*Room]bool),
	}
}

func (client *Client) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)

	_ = client.conn.SetReadDeadline(time.Now().Add(pongWait))

	client.conn.SetPongHandler(func(string) error {
		_ = client.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			break
		}
		client.handleNewMessage(jsonMessage)
	}
}

func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.send:
			_ = client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, _ = w.Write(message)

			n := len(client.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newLine)
				_, _ = w.Write(<-client.send)
			}

			if err = w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.hub.unregister <- client
	for room := range client.rooms {
		room.unregister <- client
	}

	close(client.send)
	_ = client.conn.Close()
}

func ServeWs(hub *Hub, ctx *gin.Context) {
	userId := ctx.MustGet(model.CONTEXT_USER_KEY).(string)

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, hub, userId)

	go client.writePump()
	go client.readPump()

	hub.register <- client
}

func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message model.ReceivedMessage
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmurshal JSON message (%s)", err.Error())
	}

	log.Printf("new message received (%v)", message)

	switch message.Action {
	case JoinUserAction:
		client.handleJoinRoomMessage(message)
	case LeaveRoomAction:
		client.handleLeaveRoomMessage(message)
	}
}

func (client *Client) handleJoinRoomMessage(message model.ReceivedMessage) {
	roomId := message.Room

	room := client.hub.findRoomById(roomId)
	if room == nil {
		room = client.hub.createRoom(roomId)
	}

	client.rooms[room] = true

	room.register <- client
}

func (client *Client) handleLeaveRoomMessage(message model.ReceivedMessage) {
	roomId := message.Room

	room := client.hub.findRoomById(roomId)
	if room != nil {
		delete(client.rooms, room)

		room.unregister <- client
	}
}
