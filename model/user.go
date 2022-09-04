package model

import (
	"context"
	"github.com/zakariawahyu/go-echo-mongo-basic/config"
	"github.com/zakariawahyu/go-echo-mongo-basic/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")

func GetALlUser(ctx context.Context) (*mongo.Cursor, error) {
	result, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(ctx context.Context, user entity.User) (*mongo.InsertOneResult, error) {
	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserById(ctx context.Context, id primitive.ObjectID, user *entity.User) error {
	if err := userCollection.FindOne(ctx, bson.M{"id": id}).Decode(user); err != nil {
		return err
	}
	return nil
}

func UpdateUser(ctx context.Context, id primitive.ObjectID, user bson.M) (*mongo.UpdateResult, error) {
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := userCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}
