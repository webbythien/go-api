package services

import (
	"context"
	"encoding/json"
	"errors"
	"lido-core/v1/app/models"
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/tasks"
	"lido-core/v1/pkg/workers"
	"lido-core/v1/platform/database"
	"lido-core/v1/platform/socket"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type m map[string]interface{}

func CreateQuiz(quiz *schemas.Quiz) error {
	//quizCollection := database.Collection(models.QuizColl)
	//ansCollection := database.Collection(models.AnswerColl)
	quizCollection := models.QuizCollection()
	ansCollection := models.AnswerCollection()
	quizDocs, err := quizCollection.InsertOne(context.TODO(), models.Quiz{
		LiveID:   quiz.LiveID,
		Question: quiz.Question,
		Status:   "close",
		Active:   false,
	})
	if err != nil {
		return err
	}
	quizID := quizDocs.InsertedID.(primitive.ObjectID)
	for _, ans := range quiz.Answers {
		_, err = ansCollection.InsertOne(context.TODO(), models.Answer{
			Qid:          quizID.Hex(),
			LiveID:       quiz.LiveID,
			Text:         ans.Text,
			Correct:      ans.Correct,
			UserResponse: 0,
			Options:      ans.Options,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetQuiz(id string) (interface{}, error) {
	//quizColl := database.Collection(models.QuizColl)
	//ansColl := database.Collection(models.AnswerColl)
	quizColl := models.QuizCollection()
	ansColl := models.AnswerCollection()
	quizFilter := bson.M{
		"live_id": id,
		"status":  "open",
	}
	var quiz models.Quiz
	if err := quizColl.FindOne(context.TODO(), quizFilter).Decode(&quiz); err != nil {
		return nil, err
	}
	ansFilter := bson.M{"qid": quiz.Id.Hex()}
	cur, err := ansColl.Find(context.TODO(), ansFilter)
	if err != nil {
		return nil, err
	}
	var answers []schemas.AnswerResponse
	for cur.Next(context.TODO()) {
		var ans models.Answer
		err = cur.Decode(&ans)
		if err != nil {
			return nil, err
		}
		answers = append(answers, schemas.AnswerResponse{
			Text:    ans.Text,
			Id:      ans.Id.Hex(),
			Options: ans.Options,
		})
	}
	return m{
		"question": quiz,
		"answers":  answers,
	}, nil
}

func QuizAnswer(qid, aid, address string) error {
	quizCollection := models.QuizCollection()
	_id, err := primitive.ObjectIDFromHex(qid)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id":    _id,
		"status": "open",
	}
	var quiz models.Quiz
	if err := quizCollection.FindOne(context.TODO(), filter).Decode(&quiz); err != nil {
		return errors.New("quiz is not available")
	}
	go workers.Delay(
		"core_api_queue",
		"Worker.QuizAnswer",
		tasks.QuizAnswer,
		qid,
		aid,
		address,
	)
	return nil
}

func QuizStart(id string) error {
	collection := models.QuizCollection()
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// open := bson.M{"status": "open"}
	// _update := bson.M{"$set": bson.M{"status": "closed"}}
	// _, err = collection.UpdateOne(context.TODO(), open, _update)
	// if err != nil {
	// 	return err
	// }

	filter := bson.M{
		"status": "close",
		"_id":    _id,
	}
	update := bson.M{"$set": bson.M{"status": "open", "active": true}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	// find live id
	var quiz models.Quiz
	err = collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&quiz)
	if err != nil {
		return err
	}

	_quiz, err := GetQuiz(quiz.LiveID)
	if err != nil {
		return err
	}
	data, err := json.Marshal(_quiz)
	if err != nil {
		return err
	}
	if ok := socket.New().In(quiz.LiveID).Emit("games", data); !ok {
		log.Fatal("Error creating socket")
	}
	return nil
}

func QuizStat(id, address string) (interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var playAns models.Player
	if address != "" {
		if err := database.Collection(models.PlayerColl).FindOne(
			context.TODO(),
			bson.M{"wallet_address": address},
		).Decode(&playAns); err != nil {
			return nil, err
		}
	}

	quizColl := models.QuizCollection()
	quizFilter := bson.M{"_id": _id}
	quiz := new(models.Quiz)
	if err := quizColl.FindOne(context.TODO(), quizFilter).Decode(quiz); err != nil {
		return nil, err
	}

	ansColl := models.AnswerCollection()
	ansFilter := bson.M{"qid": quiz.Id.Hex()}
	var answers []schemas.AnswerStatResponse
	cur, err := ansColl.Find(context.TODO(), ansFilter)
	defer func() {
		if err := cur.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var ans models.Answer
		err = cur.Decode(&ans)
		if err != nil {
			return nil, err
		}
		answers = append(answers, schemas.AnswerStatResponse{
			Text:         ans.Text,
			Id:           ans.Id.Hex(),
			UserResponse: ans.UserResponse,
			Options:      ans.Options,
			Correct:      ans.Correct,
			YourAnswer:   playAns.AnswerID == ans.Id.Hex(),
		})
	}

	return m{
		"question": quiz,
		"answers":  answers,
	}, nil
}

func QuizClose(id string) error {
	collection := models.QuizCollection()
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"status": "open",
		"_id":    _id,
	}
	update := bson.M{"$set": bson.M{"status": "close"}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func QuizActivate(liveId string) error {
	collection := models.QuizCollection()
	//_id, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return err
	//}
	open := bson.M{"status": "open"}
	_update := bson.M{"$set": bson.M{"status": "close"}}
	_, err := collection.UpdateMany(context.TODO(), open, _update)
	if err != nil {
		return err
	}

	filter := bson.M{
		"live_id": liveId,
		"status":  "close",
		"active":  false,
	}
	update := bson.M{"$set": bson.M{"status": "open", "active": true}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	quizData, err := GetQuiz(liveId)
	if err != nil {
		return errors.New("no quiz found for live")
	}
	quizByte, err := json.Marshal(quizData)
	if err != nil {
		return err
	}
	if ok := socket.New().In(liveId).Emit("games", quizByte); !ok {
		log.Fatal("Error creating socket")
	}
	return nil
}
