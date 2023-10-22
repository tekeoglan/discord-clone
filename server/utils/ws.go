package utils

import (
	"encoding/json"
	"github/tekeoglan/discord-clone/model"
	"log"
)

func EncodeWsMessage(message *model.WebsocketMessage) []byte {
	encoding, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return encoding
}
