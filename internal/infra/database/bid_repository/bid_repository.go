package bid_repository

import (
	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	ID        string  `bson:"_id"`
	UserID    string  `bson:"user_id"`
	AuctionID string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository auction_entity.AuctionRepositoryInterface
}

func NewBidRepository(
	database *mongo.Database,
	auctionRepository auction_entity.AuctionRepositoryInterface,
) bid_entity.BidRepositoryInterface {
	return &BidRepository{
		Collection:        database.Collection("bids"),
		AuctionRepository: auctionRepository,
	}
}
