package bid_repository

import (
	"context"
	"sync"

	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

func (r *BidRepository) CreateBid(ctx context.Context, bidEntities []*bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)
		go func(bid *bid_entity.Bid) {
			defer wg.Done()

			auction, err := r.AuctionRepository.FindAuctionById(ctx, bid.AuctionID)
			if err != nil {
				logger.Error("Error finding auction with ID: "+bid.AuctionID, err)
				return
			}

			if auction.Status != auction_entity.Active {
				return
			}

			bidMongo := BidEntityMongo{
				ID:        bid.ID,
				UserID:    bid.UserID,
				AuctionID: bid.AuctionID,
				Amount:    bid.Amount,
				Timestamp: bid.Timestamp.Unix(),
			}

			_, errMongo := r.Collection.InsertOne(ctx, bidMongo)
			if errMongo != nil {
				logger.Error("Error inserting bid", errMongo)
				return
			}
		}(bid)
	}

	wg.Wait()

	return nil
}
