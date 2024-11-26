package bid_usecase

import (
	"context"

	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

func (uc *BidUsecase) FindBidByAuctionId(
	ctx context.Context, auctionID string,
) ([]*BidOutputDTO, *internal_error.InternalError) {
	bids, err := uc.bidRepository.FindBidByAuctionId(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	bidOutputDTOs := make([]*BidOutputDTO, 0, len(bids))
	for _, bid := range bids {
		bidOutputDTOs = append(bidOutputDTOs, &BidOutputDTO{
			ID:        bid.ID,
			UserID:    bid.UserID,
			AuctionID: bid.AuctionID,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidOutputDTOs, nil
}

func (uc *BidUsecase) FindWinningBidByAuctionId(
	ctx context.Context, auctionID string,
) (*BidOutputDTO, *internal_error.InternalError) {
	bid, err := uc.bidRepository.FindWinningBidByAuctionId(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	return &BidOutputDTO{
		ID:        bid.ID,
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}, nil
}
