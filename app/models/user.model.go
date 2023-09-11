package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lido-core/v1/platform/database"
)

const (
	PlayerColl string = "player"
)

func PlayerCollection() *mongo.Collection {
	return database.Collection(PlayerColl)
}

type Player struct {
	WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	LiveID        string `json:"live_id" bson:"live_id"`
	QuizID        string `json:"quiz_id" bson:"quiz_id"`
	AnswerID      string `json:"answer_id" bson:"answer_id"`
	Correct       bool   `json:"correct"`
}
