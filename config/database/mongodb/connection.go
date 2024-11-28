package mongodb

import (
	"context"
	"os"

	"github.com/KelpGF/Go-Auction/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL    = "MONGODB_URL"
	MONGODB_DBNAME = "MONGODB_DBNAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDBName := os.Getenv(MONGODB_DBNAME)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoURL),
	)

	if err != nil {
		logger.Error("Error connecting to MongoDB", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error pinging to MongoDB", err)
		return nil, err
	}

	return client.Database(mongoDBName), nil
}
