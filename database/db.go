package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDB(ctx context.Context) (*mongo.Client, error) {
	mongoCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(mongoCtx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(mongoCtx, nil)
	if err != nil {
		return nil, err
	}

	//fmt.Println("Connected to MongoDB!")
	return client, nil
}
