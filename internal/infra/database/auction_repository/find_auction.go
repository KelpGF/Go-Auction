package auction_repository

import (
	"context"
	"errors"
	"time"

	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	var auctionMongo AuctionEntityMongo

	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(ctx, filter).Decode(&auctionMongo)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("Auction not found with ID: "+id, err)
			return nil, internal_error.NewNotFoundError("Auction not found with ID: " + id)
		}

		logger.Error("Error finding auction with ID: "+id, err)
		return nil, internal_error.NewInternalServerError("Error finding auction with ID: " + id)
	}

	auctionEntity := &auction_entity.Auction{
		ID:          auctionMongo.ID,
		ProductName: auctionMongo.ProductName,
		Category:    auctionMongo.Category,
		Description: auctionMongo.Description,
		Condition:   auctionMongo.Condition,
		Status:      auctionMongo.Status,
		Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
	}

	return auctionEntity, nil
}

func (r *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category, productName string,
) ([]*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["product_name"] = primitive.Regex{Pattern: productName, Options: "i"}
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error finding auctions", err)
		return nil, internal_error.NewInternalServerError("Error finding auctions")
	}
	defer cursor.Close(ctx)

	var auctionsMongo []AuctionEntityMongo
	err = cursor.All(ctx, &auctionsMongo)

	if err != nil {
		logger.Error("Error decoding auctions", err)
		return nil, internal_error.NewInternalServerError("Error decoding auctions")
	}

	auctions := make([]*auction_entity.Auction, 0, len(auctionsMongo))
	for _, auctionMongo := range auctionsMongo {
		auction := &auction_entity.Auction{
			ID:          auctionMongo.ID,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		}

		auctions = append(auctions, auction)
	}

	return auctions, nil
}
