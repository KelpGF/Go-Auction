package bid_repository

import (
	"context"
	"time"

	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *BidRepository) FindBidByAuctionId(
	ctx context.Context, auctionID string,
) ([]*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error finding bids by auction ID "+auctionID, err)
		return nil, internal_error.NewInternalServerError("Error finding bids by auction ID " + auctionID)
	}

	var bidsMongo []*BidEntityMongo
	if err = cursor.All(ctx, &bidsMongo); err != nil {
		logger.Error("Error decoding bids by auction ID "+auctionID, err)
		return nil, internal_error.NewInternalServerError("Error decoding bids by auction ID " + auctionID)
	}

	bids := make([]*bid_entity.Bid, 0, len(bidsMongo))
	for _, bidMongo := range bidsMongo {
		bid := &bid_entity.Bid{
			ID:        bidMongo.ID,
			UserID:    bidMongo.UserID,
			AuctionID: bidMongo.AuctionID,
			Amount:    bidMongo.Amount,
			Timestamp: time.Unix(bidMongo.Timestamp, 0),
		}
		bids = append(bids, bid)
	}

	return bids, nil
}

func (r *BidRepository) FindWinningBidByAuctionId(
	ctx context.Context, auctionID string,
) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionID}
	opt := options.FindOne().SetSort(bson.M{"amount": -1})

	var bidMongo *BidEntityMongo

	err := r.Collection.FindOne(ctx, filter, opt).Decode(&bidMongo)
	if err != nil {
		logger.Error("Error finding winning bid by auction ID "+auctionID, err)
		return nil, internal_error.NewInternalServerError("Error finding winning bid by auction ID " + auctionID)
	}

	bid := &bid_entity.Bid{
		ID:        bidMongo.ID,
		UserID:    bidMongo.UserID,
		AuctionID: bidMongo.AuctionID,
		Amount:    bidMongo.Amount,
		Timestamp: time.Unix(bidMongo.Timestamp, 0),
	}

	return bid, nil
}
