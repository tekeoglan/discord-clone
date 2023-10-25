package repository

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const messageCursorLength = 20

type messageRepository struct {
	database   mongo.Database
	collection string
}

func NewMessageRepository(db mongo.Database, collection string) model.MessageRepository {
	return &messageRepository{
		database:   db,
		collection: collection,
	}
}

func (mr *messageRepository) CreateMessage(c context.Context, message *model.Message) (*model.Message, error) {
	collection := mr.database.Collection(mr.collection)
	_, err := collection.InsertOne(c, message)

	return message, err
}

func (mr *messageRepository) UpdateMessage(c context.Context, id string, update interface{}) error {
	collection := mr.database.Collection(mr.collection)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(c, bson.M{"_id": objId}, update)

	return err
}

func (mr *messageRepository) DeleteMessage(c context.Context, id string) error {
	collection := mr.database.Collection(mr.collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectId})
	return err
}

func (mr *messageRepository) GetByID(c context.Context, id string) (*model.Message, error) {
	collection := mr.database.Collection(mr.collection)

	var message model.Message

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": objectId}).Decode(&message)
	return &message, err
}

func (mr *messageRepository) GetChannelMessages(c context.Context, channelId string, cursorPos int) (*[]model.Message, error) {
	collection := mr.database.Collection(mr.collection)

	pipe := bson.A{
		bson.M{
			"$match": bson.M{
				"channelId": channelId,
			},
		},
		bson.M{
			"$skip": cursorPos,
		},
		bson.M{
			"$limit": cursorPos + messageCursorLength,
		},
		bson.M{
			"$set": bson.M{
				"cursorPos": cursorPos + messageCursorLength,
			},
		},
	}

	cur, err := collection.Aggregate(c, pipe)
	if err != nil {
		return nil, err
	}

	var messages []model.Message
	err = cur.All(c, &messages)

	return &messages, err
}
