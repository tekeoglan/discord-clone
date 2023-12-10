package service

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type messageService struct {
	messageRepository model.MessageRepository
}

func NewMessageService(mr model.MessageRepository) model.MessageService {
	return &messageService{
		messageRepository: mr,
	}
}

func (ms *messageService) CreateMessage(c context.Context, m *model.Message) (*model.Message, error) {
	return ms.messageRepository.CreateMessage(c, m)
}

func (ms *messageService) UpdateMessage(c context.Context, id string, text string) error {
	update := bson.M{"$set": bson.M{"updatedAt": time.Now(), "text": text}}
	return ms.messageRepository.UpdateMessage(c, id, update)
}

func (ms *messageService) DeleteMessage(c context.Context, id string) error {
	return ms.messageRepository.DeleteMessage(c, id)
}
func (ms *messageService) GetMessage(c context.Context, id string) (*model.Message, error) {
	return ms.messageRepository.GetByID(c, id)
}

func (ms *messageService) GetChannelMessages(c context.Context, channelId string, cursorPos int) (model.MessageGetAllResult, error) {
	return ms.messageRepository.GetChannelMessages(c, channelId, cursorPos)
}
