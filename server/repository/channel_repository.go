package repository

import (
	"context"
	"github/tekeoglan/discord-clone/model"
	"github/tekeoglan/discord-clone/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type channelRepository struct {
	database   mongo.Database
	collection string
}

func NewChannelRepository(db mongo.Database, collection string) model.ChannelRepository {
	return &channelRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *channelRepository) CreateChannel(c context.Context, channel interface{}) error {
	col := cr.database.Collection(cr.collection)

	_, err := col.InsertOne(c, channel)

	return err
}

func (cr *channelRepository) GetFriendChannels(c context.Context, userId string) (*[]model.FriendChannelResult, error) {
	col := cr.database.Collection(cr.collection)

	var result []model.FriendChannelResult

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	pipe := bson.A{
		bson.M{
			"$match": bson.M{
				"users": objectId,
			},
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
					"$ne": objectId,
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"messages":            0,
				"users":               0,
				"friendInfo.password": 0,
			},
		},
	}

	cur, err := col.Aggregate(c, pipe)

	err = cur.All(c, &result)

	return &result, err
}

func (cr *channelRepository) GetFcById(c context.Context, id string) (*model.FriendChannel, error) {
	col := cr.database.Collection(cr.collection)

	var channel model.FriendChannel
	err := col.FindOne(c, bson.M{"channelId": id}).Decode(&channel)

	return &channel, err
}

func (cr *channelRepository) GetFcByUserIds(c context.Context, user1Id string, user2Id string) (*model.FriendChannel, error) {
	col := cr.database.Collection(cr.collection)

	user1ObjId, err := primitive.ObjectIDFromHex(user1Id)
	if err != nil {
		return nil, err
	}

	var user2ObjId primitive.ObjectID
	user2ObjId, err = primitive.ObjectIDFromHex(user2Id)
	if err != nil {
		return nil, err
	}

	var channel model.FriendChannel
	err = col.FindOne(c, bson.M{
		"users": bson.M{
			"$all": bson.A{user1ObjId, user2ObjId},
		},
	}).Decode(&channel)

	return &channel, err
}

func (cr *channelRepository) DeleteChannel(c context.Context, id string) error {
	col := cr.database.Collection(cr.collection)

	_, err := col.DeleteOne(c, bson.M{"channelId": id})

	return err
}
