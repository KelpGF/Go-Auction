package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type BidInputDTO struct {
	UserID    string  `json:"user_id"`
	AuctionID string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	AuctionID string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUsecaseInterface interface {
	CreateBid(
		ctx context.Context, bidInputDTO *BidInputDTO,
	) *internal_error.InternalError

	FindBidByAuctionId(
		ctx context.Context, auctionID string,
	) ([]*BidOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context, auctionID string,
	) (*BidOutputDTO, *internal_error.InternalError)
}

type BidUsecase struct {
	bidRepository bid_entity.BidRepositoryInterface

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan *bid_entity.Bid
}

func NewBidUseCase(
	bidRepository bid_entity.BidRepositoryInterface,
) BidUsecaseInterface {
	batchInsertInterval := getMaxBatchInsertInterval()
	maxBatchSize := getMaxBatchSize()

	bidUsecase := &BidUsecase{
		bidRepository:       bidRepository,
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: batchInsertInterval,
		timer:               time.NewTimer(batchInsertInterval),
		bidChannel:          make(chan *bid_entity.Bid, maxBatchSize),
	}

	bidUsecase.triggerCreateBatchRoutine(context.Background())

	return bidUsecase
}

func getMaxBatchSize() int {
	batchSize, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}

	return batchSize
}

func getMaxBatchInsertInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}

	return duration
}
