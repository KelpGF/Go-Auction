package auction_usecase

import (
	"context"

	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

func (uc *AuctionUseCase) CreateAuction(
	ctx context.Context, input *AuctionInputDTO,
) (string, *internal_error.InternalError) {
	auction, err := auction_entity.NewAuction(
		input.ProductName, input.Category, input.Description, auction_entity.ProductCondition(input.Condition),
	)
	if err != nil {
		return "", err
	}

	err = uc.AuctionRepository.CreateAuction(ctx, auction)
	if err != nil {
		return "", err
	}

	return auction.ID, nil
}
