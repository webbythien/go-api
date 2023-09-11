package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Advertisement struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ImageUrl string             `json:"image_url" bson:"image_url"`
	Link     string             `json:"link"`
	Created  int64              `json:"created"`
}
