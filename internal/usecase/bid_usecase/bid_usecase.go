package bid_usecase

import (
	"context"
	"time"

	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type BidOutputDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	AuctionID string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUsecaseInterface interface {
	FindBidByAuctionId(
		ctx context.Context, auctionID string,
	) ([]*BidOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context, auctionID string,
	) (*BidOutputDTO, *internal_error.InternalError)
}

type BidUsecase struct {
	bidRepository bid_entity.BidRepositoryInterface
}
