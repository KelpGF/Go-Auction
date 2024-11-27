package bid_entity

import (
	"context"
	"time"

	"github.com/KelpGF/Go-Auction/internal/internal_error"
	"github.com/google/uuid"
)

type Bid struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}

func NewBid(
	userID, auctionID string,
	amount float64,
) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		ID:        uuid.New().String(),
		UserID:    userID,
		AuctionID: auctionID,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	err := bid.Validate()
	if err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	if err := uuid.Validate(b.UserID); err != nil {
		return internal_error.NewBadRequestError("UserId id is not a valid ID")
	}

	if err := uuid.Validate(b.AuctionID); err != nil {
		return internal_error.NewBadRequestError("AuctionID id is not a valid ID")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Amount must be greater than 0")
	}

	return nil
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
