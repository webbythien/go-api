package services

import (
	"context"
	"errors"
	"fmt"
	"lido-core/v1/app/models"
	"lido-core/v1/app/schemas"
	"lido-core/v1/pkg/utils"
	"log"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func SignIn(request *schemas.SignIn) (string, error) {
	if !utils.VerifySig(request.Address, request.Signature, []byte(request.Message)) {
		return "", errors.New("invalid signature")
	}
	token, err := utils.GenerateToken(request.Address)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateMessage() string {
	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")
	return fmt.Sprintf("Signature Message ID: #%s", id)
}

func GetQuizResult(address, liveId string) (interface{}, error) {
	collection := models.PlayerCollection()
	filter := bson.M{
		"wallet_address": address,
		"live_id":        liveId,
	}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cur.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	var players []models.Player
	for cur.Next(context.TODO()) {
		var player models.Player
		if err := cur.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}
