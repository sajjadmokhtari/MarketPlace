package db

import (
	"MarketPlace/logging"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() error {
	uri := "mongodb://admin:password123@localhost:27018"
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ساخت و اتصال مستقیم با mongo.Connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logging.GetLogger().Errorw("❌ failed to connect to MongoDB", "error", err)
		return err
	}

	// تست اتصال با Ping
	if err := client.Ping(ctx, nil); err != nil {
		logging.GetLogger().Errorw("❌ MongoDB ping failed", "error", err)
		return err
	}

	logging.GetLogger().Infow("✅ connected to MongoDB successfully", "uri", uri)
	MongoClient = client
	return nil
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}
