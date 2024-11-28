package auction_usecase

import (
	"context"
	"time"

	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
	"github.com/KelpGF/Go-Auction/internal/usecase/bid_usecase"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=1"`
	Description string           `json:"description" binding:"required,min=10"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	ID          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductCondition int
type AuctionStatus int

type WinningInfoOutputDTO struct {
	Auction *AuctionOutputDTO         `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type AuctionUseCaseInterface interface {
	CreateAuction(
		ctx context.Context, input *AuctionInputDTO,
	) *internal_error.InternalError

	FindAuctionByID(
		ctx context.Context, id string,
	) (*AuctionOutputDTO, *internal_error.InternalError)

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string,
	) ([]*AuctionOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionID(
		ctx context.Context, auctionID string,
	) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

type AuctionUseCase struct {
	AuctionRepository auction_entity.AuctionRepositoryInterface
	BidRepository     bid_entity.BidRepositoryInterface
}
