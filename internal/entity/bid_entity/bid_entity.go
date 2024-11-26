package bid_entity

import (
	"context"
	"time"

	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type Bid struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntities []*Bid) *internal_error.InternalError
	FindBidByAuctionId(
		ctx context.Context, auctionID string,
	) ([]*Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(
		ctx context.Context, auctionID string,
	) (*Bid, *internal_error.InternalError)
}
