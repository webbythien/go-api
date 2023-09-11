package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lido-core/v1/platform/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	QuizColl   string = "quizzes"
	AnswerColl string = "answer"
)

func QuizCollection() *mongo.Collection {
	return database.Collection(QuizColl)
}

func AnswerCollection() *mongo.Collection {
	return database.Collection(AnswerColl)
}

type Quiz struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status   string             `json:"status"`
	LiveID   string             `json:"live_id" bson:"live_id"`
	Question string             `json:"question"`
	Active   bool               `json:"active"`
}

type Answer struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Qid          string             `json:"qid"`
	LiveID       string             `json:"live_id" bson:"live_id"`
	Text         string             `json:"text" `
	Correct      bool               `json:"correct"`
	UserResponse int64              `json:"user_response" bson:"user_response"`
	Options      string             `json:"options"`
}
