package main

import (
	"context"
	"log"

	"github.com/KelpGF/Go-Auction/config/database/mongodb"
	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/infra/api/web/controller/auction_controller"
	"github.com/KelpGF/Go-Auction/internal/infra/api/web/controller/bid_controller"
	"github.com/KelpGF/Go-Auction/internal/infra/api/web/controller/user_controller"
	"github.com/KelpGF/Go-Auction/internal/infra/database/auction_repository"
	"github.com/KelpGF/Go-Auction/internal/infra/database/bid_repository"
	"github.com/KelpGF/Go-Auction/internal/infra/database/user_repository"
	"github.com/KelpGF/Go-Auction/internal/usecase/auction_usecase"
	"github.com/KelpGF/Go-Auction/internal/usecase/bid_usecase"
	"github.com/KelpGF/Go-Auction/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("Starting the application...")

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	mongoDatabase, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
		return
	}

	logger.Info("MongoDB connection successful")

	userController, auctionController, bidController := initDependencies(mongoDatabase)

	router := gin.Default()

	router.GET("/auction", auctionController.FindAuctions)
	router.POST("/auction", auctionController.CreateAuction)
	router.GET("/auction/:id", auctionController.FindById)
	router.GET("/auction/:id/winner", auctionController.FindWinningBidByAuctionID)

	router.POST("/bid", bidController.Create)
	router.GET("/bid/:auction_id", bidController.FindByAuctionId)

	router.GET("/user/:id", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(
	mongoDatabase *mongo.Database,
) (
	*user_controller.UserController,
	*auction_controller.AuctionController,
	*bid_controller.BidController,
) {
	auctionRepos := auction_repository.NewAuctionRepository(mongoDatabase)
	bidRepos := bid_repository.NewBidRepository(mongoDatabase, auctionRepos)
	userRepos := user_repository.NewUserRepository(mongoDatabase)

	auctionUseCases := auction_usecase.NewAuctionUseCase(auctionRepos, bidRepos)
	bidUseCases := bid_usecase.NewBidUseCase(bidRepos)
	userUseCases := user_usecase.NewUserUseCase(userRepos)

	userController := user_controller.NewUserController(userUseCases)
	auctionController := auction_controller.NewAuctionController(auctionUseCases)
	bidController := bid_controller.NewBidController(bidUseCases)

	return userController, auctionController, bidController
}
