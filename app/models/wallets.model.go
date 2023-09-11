package models

import (
	"lido-core/v1/platform/database"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Wallets string = "wallets"
)

func WalletCollection() *mongo.Collection {
	return database.Collection(Wallets)
}

type Wallet struct {
	Address   string `json:"address"`
	Balance   string `json:"balance"`
	LastBlock int64  `json:"last_block" bson:"last_block"`
}
