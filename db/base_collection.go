package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseCollection struct {
	collectionName string
}

func NewBaseCollection(collectionName string) *BaseCollection{
	return &BaseCollection{
		collectionName: collectionName,
	}
}

func (baseColl *BaseCollection) GetCollection() *mongo.Collection {
	return Mc.GetCollection(baseColl.collectionName, nil)
}

func (baseColl *BaseCollection) InsertItem(item interface{}) (string, error) {
	collection := baseColl.GetCollection()
	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).String(), nil
}

func (baseColl *BaseCollection) FindOneItemById(objectId string) (*mongo.SingleResult, error) {
	collection := baseColl.GetCollection()
	itemIdObject, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return nil, err
	}
	result := collection.FindOne(context.TODO(), bson.M{"_id": itemIdObject})
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}

func (baseColl *BaseCollection) FindOneItemByKey(key string, value interface{}) (*mongo.SingleResult, error) {
	//collection := baseColl.GetCollection()
	return nil, nil
}

func (baseColl *BaseCollection) UpdateItemById(objectId string, newItem interface{}) error {
	panic("implement me")
}

func (baseColl *BaseCollection) UpdateItemByKey(key string, value interface{}, newItem interface{}) error {
	panic("implement me")
}

func (baseColl *BaseCollection) DeleteItemById(objectId string) error {
	panic("implement me")
}

func (baseColl *BaseCollection) DeleteItemByKey(key string, value interface{}) error {
	panic("implement me")
}
