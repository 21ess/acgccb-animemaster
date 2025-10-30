package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	MongoClient *mongo.Client
)

func InitMongo() {
	MongoClient, _ = mongo.Connect(options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = MongoClient.Ping(ctx, readpref.Primary())

}
