package mongo

import (
	"conn/internal/models"
	"context"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

var client *mongo.Client
var collection *mongo.Collection

func InitMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	collection = client.Database("webhooks").Collection("repository_events")
	return nil
}

func DisconnectMongoDB() {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect MongoDB: %v", err)
	}
}

func SaveToMongoDB(event models.RepositoryEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	document := bson.M{
		"action": event.Action,
		"repository": bson.M{
			"name":      event.Repository.Name,
			"full_name": event.Repository.FullName,
		},
		"received_at": time.Now(),
	}

	_, err := collection.InsertOne(ctx, document)
	return err
}
