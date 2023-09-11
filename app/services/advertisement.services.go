package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"lido-core/v1/app/models"
	"lido-core/v1/app/schemas"
	"log"
)

func GetAllAdvertisements() ([]schemas.Advertisement, error) {
	collection := models.AdvertisementCollection()
	filter := bson.M{}
	var advertisements []schemas.Advertisement
	// Find documents in the collection that match the filter
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and decode each document into an Advertisement
	for cursor.Next(context.TODO()) {
		var advertisement models.Advertisement
		err := cursor.Decode(&advertisement)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Append the decoded advertisement to the slice
		advertisements = append(advertisements, schemas.Advertisement{
			Id:       advertisement.Id,
			ImageUrl: advertisement.ImageUrl,
			Link:     advertisement.Link,
			Created:  advertisement.Created,
		})
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return advertisements, nil
}
