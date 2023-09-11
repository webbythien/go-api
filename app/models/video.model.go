package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lido-core/v1/platform/database"
)

const (
	Videos string = "videos"
)

func VideoCollection() *mongo.Collection {
	return database.Collection(Videos)
}

type Video struct {
	Host          string `json:"host"`
	Title         string `json:"title"`
	Background    string `json:"background"`
	View          int64  `json:"view"`
	Id            string `json:"id"`
	Status        string `json:"status"`
	Created       int64  `json:"created"`
	Desc          string `json:"desc"`
	NameGalxe     string `json:"name_galxe" bson:"name_galxe"`
	LinkGalxe     string `json:"link_galxe" bson:"link_galxe"`
	ImageGalxeUrl string `json:"image_galxe_url" bson:"image_galxe_url"`
}
