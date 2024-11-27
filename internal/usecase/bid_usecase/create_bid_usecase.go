package bid_usecase

import (
	"context"

	"github.com/KelpGF/Go-Auction/config/logger"
	"github.com/KelpGF/Go-Auction/internal/entity/bid_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

var bidBatch []*bid_entity.Bid

func (uc *BidUsecase) CreateBid(ctx context.Context, bidInputDTO *BidInputDTO) *internal_error.InternalError {
	bidEntity, err := bid_entity.NewBid(
		bidInputDTO.UserID,
		bidInputDTO.AuctionID,
		bidInputDTO.Amount,
	)

	if err != nil {
		return err
	}

	uc.bidChannel <- bidEntity

	return nil
}

func (uc *BidUsecase) insertBatch(
	ctx context.Context, bids []*bid_entity.Bid,
) *internal_error.InternalError {
	err := uc.bidRepository.CreateBid(ctx, bids)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BidUsecase) triggerCreateBatchRoutine(ctx context.Context) {
	go func() {
		defer close(uc.bidChannel)

		for {
			select {
			case bidEntity, ok := <-uc.bidChannel:
				bidBatch = append(bidBatch, bidEntity)

				if !ok {
					if len(bidBatch) > 0 {
						err := uc.insertBatch(ctx, bidBatch)
						if err != nil {
							uc.logInsertBatchError(err)
						}
					}

					return
				}

				if len(bidBatch) >= uc.maxBatchSize {
					err := uc.insertBatch(ctx, bidBatch)
					if err != nil {
						uc.logInsertBatchError(err)
					}

					uc.resetBatch()
				}
			case <-uc.timer.C:
				if len(bidBatch) > 0 {
					err := uc.insertBatch(ctx, bidBatch)
					if err != nil {
						uc.logInsertBatchError(err)
					}

					uc.resetBatch()
				}
			}
		}
	}()
}

func (uc *BidUsecase) resetBatch() {
	bidBatch = nil
	uc.timer.Reset(uc.batchInsertInterval)
}

func (uc *BidUsecase) logInsertBatchError(err error) {
	logger.Error("Error inserting batch", err)
}
