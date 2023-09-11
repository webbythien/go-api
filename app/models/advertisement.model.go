package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lido-core/v1/platform/database"
)

const (
	Advertisements string = "advertisement"
)

func AdvertisementCollection() *mongo.Collection {
	return database.Collection(Advertisements)
}

type Advertisement struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ImageUrl string             `json:"image_url" bson:"image_url"`
	Link     string             `json:"link"`
	Created  int64              `json:"created"`
}
