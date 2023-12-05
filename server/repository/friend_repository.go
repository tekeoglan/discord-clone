package repository

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const friendCursorLength = 20

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

func (fr *friendRepository) Update(c context.Context, id string, update interface{}) error {
	collection := fr.database.Collection(fr.collection)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": objId}, update)

	return err
}

func (fr *friendRepository) Get(c context.Context, id string) (model.FriendGetResult, error) {
	collection := fr.database.Collection(fr.collection)

	var friend model.FriendGetResult

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return friend, err
	}

	pipe := bson.A{
		bson.M{
			"$match": bson.M{
				"_id": objId,
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "users",
				"foreignField": "_id",
				"as":           "friendInfos",
			},
		},
	}

	var cursor mongo.Cursor
	cursor, err = collection.Aggregate(c, pipe)
	defer cursor.Close(c)

	for cursor.Next(c) {
		err = cursor.Decode(&friend)
	}

	return friend, err
}

func (fr *friendRepository) GetByUserIds(c context.Context, user1 string, user2 string) (model.Friend, error) {
	collection := fr.database.Collection(fr.collection)

	var result model.Friend

	user1ObjId, err := primitive.ObjectIDFromHex(user1)
	if err != nil {
		return result, err
	}

	var user2ObjId primitive.ObjectID
	user2ObjId, err = primitive.ObjectIDFromHex(user2)
	if err != nil {
		return result, err
	}

	err = collection.FindOne(c, bson.M{
		"users": bson.M{
			"$all": bson.A{user1ObjId, user2ObjId},
		},
	}).Decode(&result)

	return result, err
}

func (fr *friendRepository) GetConfirmed(c context.Context, id string, cursorPos int) (model.FriendGetAllResult, error) {
	collection := fr.database.Collection(fr.collection)

	var friends []model.FriendAggragateResult
	var result model.FriendGetAllResult

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
			"$limit": friendCursorLength,
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
	}

	cursor, err := collection.Aggregate(c, pipe)
	defer cursor.Close(c)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &friends)

	result.Friends = friends
	result.CursorPos = cursorPos + friendCursorLength

	return result, err
}

func (fr *friendRepository) GetPending(c context.Context, id string, cursorPos int) (model.FriendGetAllResult, error) {
	collection := fr.database.Collection(fr.collection)

	var friends []model.FriendAggragateResult
	var result model.FriendGetAllResult

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	pipe := bson.A{
		bson.M{
			"$match": bson.M{
				"$and": bson.A{
					bson.M{
						"users.1": _id,
					},
					bson.M{
						"status": "pending",
					},
				},
			},
		},
		bson.M{
			"$skip": cursorPos,
		},
		bson.M{
			"$limit": friendCursorLength,
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
	}

	cursor, err := collection.Aggregate(c, pipe)
	defer cursor.Close(c)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &friends)

	result.Friends = friends
	result.CursorPos = cursorPos + friendCursorLength

	return result, err
}

func (fr *friendRepository) Remove(c context.Context, id string) error {
	collection := fr.database.Collection(fr.collection)

	objId, err := primitive.ObjectIDFromHex(id)

	_, err = collection.DeleteOne(c, bson.M{"_id": objId})

	return err
}

func (fr *friendRepository) IsFriends(c context.Context, sender primitive.ObjectID, receiver primitive.ObjectID) (bool, error) {
	collection := fr.database.Collection(fr.collection)

	result := collection.FindOne(c, bson.M{"users": bson.M{
		"$all": bson.A{sender, receiver},
	}})

	var decodedResult bson.M
	err := result.Decode(&decodedResult)
	if err != nil && err != mongo.ErrNoDoc {
		return false, err
	}

	if err != nil && err == mongo.ErrNoDoc {
		return false, nil
	}

	return true, nil
}
