package services

import (
	"context"
	"lido-core/v1/app/models"
	"lido-core/v1/app/schemas"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LiveMain() schemas.VideoMain {
	collection := models.VideoCollection()
	filter := bson.M{"status": "live"}
	var video models.Video
	err := collection.FindOne(context.TODO(), filter).Decode(&video)
	if err != nil {
		log.Fatal(err)
	}
	return schemas.VideoMain{
		Host:          video.Host,
		Title:         video.Title,
		Background:    video.Background,
		View:          video.View,
		Id:            video.Id,
		Status:        video.Status,
		Started:       video.Created,
		Desc:          video.Desc,
		NameGalxe:     video.NameGalxe,
		LinkGalxe:     video.LinkGalxe,
		ImageGalxeUrl: video.ImageGalxeUrl,
	}
}

func RecommendVideo(limit int64) ([]schemas.VideoPrevious, error) {
	collection := models.VideoCollection()
	otp := options.Find().SetLimit(limit)
	cur, err := collection.Find(context.TODO(), bson.D{}, otp)
	if err != nil {
		log.Fatal(err)
	}
	var videos []schemas.VideoPrevious
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, context.TODO())
	for cur.Next(context.TODO()) {
		var video models.Video
		err = cur.Decode(&video)
		if err != nil {
			return nil, err
		}
		videos = append(videos, schemas.VideoPrevious{
			Host:       video.Host,
			Title:      video.Title,
			Background: video.Background,
			View:       video.View,
			Id:         video.Id,
			End:        video.Created,
		})
	}
	return videos, nil
}

func PreviousStreamDetails(id string) *schemas.PreviousStream {
	collection := models.VideoCollection()
	filter := bson.M{"id": id}
	var video = new(models.Video)
	if err := collection.FindOne(context.TODO(), filter).Decode(video); err != nil {
		return &schemas.PreviousStream{}
	}
	return &schemas.PreviousStream{
		Host:       video.Host,
		Title:      video.Title,
		Background: video.Title,
		View:       video.View,
		Id:         video.Id,
		Started:    video.Created,
		Desc:       video.Desc,
	}
}

func PreviousStreamRecommend(id string) ([]schemas.PreviousStreamRecommend, error) {
	collection := models.VideoCollection()
	otp := options.Find().SetLimit(5)
	cur, err := collection.Find(context.TODO(), bson.M{}, otp)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cur.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	var videos []schemas.PreviousStreamRecommend
	for cur.Next(context.TODO()) {
		var video models.Video
		err = cur.Decode(&video)
		if err != nil {
			return nil, err
		}
		if video.Id == id {
			continue
		}
		videos = append(videos, schemas.PreviousStreamRecommend{
			Host:       video.Host,
			Title:      video.Title,
			Background: video.Background,
			View:       video.View,
			Id:         video.Id,
			Started:    video.Created,
		})
	}
	return videos, nil
}
