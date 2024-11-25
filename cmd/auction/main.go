package main

import (
	"context"
	"log"

	"github.com/KelpGF/Go-Auction/config/database/mongodb"
	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
		return
	}

	logger.Info("MongoDB connection successful")
}
