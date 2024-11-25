package auction_repository

import (
	"context"

	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

func (r *AuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	auctionMongo := AuctionEntityMongo{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	_, err := r.Collection.InsertOne(ctx, auctionMongo)
	if err != nil {
		logger.Error("Error inserting auction", err)
		return internal_error.NewInternalServerError("Error creating auction")
	}

	return nil
}
