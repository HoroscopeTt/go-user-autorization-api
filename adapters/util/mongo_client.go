package util

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitMongoDBClient สร้างและคืนค่า MongoDB client
func InitMongoDBClient(mongoURI string, timeout int) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Ping the primary to verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")
	return client, nil
}