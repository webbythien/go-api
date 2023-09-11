package tasks

import (
	"context"
	"errors"
	"lido-core/v1/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func QuizAnswer(qid, aid, address string) error {
	ansColl := models.AnswerCollection()
	playerColl := models.PlayerCollection()
	// quizColl := models.QuizCollection()
	_id, err := primitive.ObjectIDFromHex(aid)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": _id,
		"qid": qid,
	}
	update := bson.M{
		"$inc": bson.M{
			"user_response": 1,
		},
	}
	var ans models.Answer
	if err := ansColl.FindOne(context.TODO(), filter).Decode(&ans); err != nil {
		return errors.New("answer is not available")
	}
	// var quiz models.Quiz
	// quizFilter := bson.M{
	// 	"live_id": ans.LiveID,
	// 	"status":  "open",
	// }
	// if err := quizColl.FindOne(context.TODO(), quizFilter).Decode(&quiz); err != nil {
	// 	return errors.New("quiz is not available")
	// }
	_, err = playerColl.InsertOne(context.TODO(), models.Player{
		WalletAddress: address,
		LiveID:        ans.LiveID,
		QuizID:        ans.Qid,
		AnswerID:      aid,
		Correct:       ans.Correct,
	})
	if err != nil {
		return err
	}
	_, err = ansColl.UpdateOne(context.TODO(), filter, update)
	return err
}

func QuizCreate(quiz []byte) error {
	return nil
}

func QuizClose() error {
	return nil
}
