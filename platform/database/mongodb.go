package database

import (
	"context"
	"lido-core/v1/pkg/utils"
	"log"

	"lido-core/v1/pkg/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func NewContext() (context.Context, context.CancelFunc) {
	// return context.WithTimeout(context.Background(), 10*time.Second)
	return context.TODO(), nil
}

func initDatabase() {
	uri, _ := utils.ConnectionURLBuilder("mongo")
	ctx, _ := NewContext()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("[MONGO_DB] Cannot connect to mongoDB")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Cannot ping MongoDB: %v", err)
	}
	log.Println("[MONGO_DB] connect successfully")
}

func init() {
	initDatabase()
}

func Shutdown() {
	ctx, _ := NewContext()
	if client != nil {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Cannot disconnect MongoDB: %v", err)
		}
	}
	log.Println("[MONGO_DB] disconnect successfully")
}

func Collection(col string) *mongo.Collection {
	if client == nil {
		initDatabase()
	}
	return client.Database(configs.DbName).Collection(col)
}
