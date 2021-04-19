package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection interface {
	GetCollection() *mongo.Collection
	InsertItem(item interface{}) (string, error)
	FindOneItemById(objectId string) (*mongo.SingleResult, error)
	UpdateItemById(objectId string, operation *Operation) error
	UpdateItemByKey(matchArr []*MatchItem, operation *Operation) error
	DeleteItemById(objectId string) error
	DeleteItemByKey(matchArr []*MatchItem) error
	GetModel() interface{}
}
