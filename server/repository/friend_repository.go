package repository

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CURSOR_LENGTH = 20

type friendRepository struct {
	database   mongo.Database
	collection string
}

func NewFriendRepository(db mongo.Database, collection string) model.FriendRepository {
	return &friendRepository{
		database:   db,
		collection: collection,
	}
}

func (fr *friendRepository) Add(c context.Context, friend *model.Friend) error {
	collection := fr.database.Collection(fr.collection)
	_, err := collection.InsertOne(c, friend)

	return err
}

func (fr *friendRepository) GetConfirmed(c context.Context, id string, cursorPos int) ([]interface{}, error) {
	collection := fr.database.Collection(fr.collection)

	var result []interface{} = make([]interface{}, CURSOR_LENGTH)

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	pipe := bson.A{
		bson.M{
			"$match": bson.M{
				"$and": bson.A{
					bson.M{
						"users": _id,
					},
					bson.M{
						"status": "confirmed",
					},
				},
			},
		},
		bson.M{
			"$skip": cursorPos,
		},
		bson.M{
			"$limit": cursorPos + CURSOR_LENGTH,
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "users",
				"foreignField": "_id",
				"as":           "friendInfo",
			},
		},
		bson.M{
			"$unwind": "$friendInfo",
		},
		bson.M{
			"$match": bson.M{
				"friendInfo._id": bson.M{
					"$ne": _id,
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"users":               0,
				"friendInfo.password": 0,
			},
		},
		bson.M{
			"$set": bson.M{
				"cursorPos": cursorPos + CURSOR_LENGTH,
			},
		},
	}

	cursor, err := collection.Aggregate(c, pipe)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func (fr *friendRepository) GetPending(c context.Context, id string, cursorPos int) ([]interface{}, error) {
	var result = make([]interface{}, 20)

	return result, nil
}

func (fr *friendRepository) Remove(c context.Context, id string) error {

	return nil
}
