package auction_usecase

import (
	"context"

	"github.com/KelpGF/Go-Auction/internal/entity/auction_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

func (uc *AuctionUseCase) FindAuctionByID(
	ctx context.Context, id string,
) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := uc.AuctionRepository.FindAuctionById(ctx, id)
	if auctionEntity == nil {
		return nil, err
	}

	auctionOutputDTO := &AuctionOutputDTO{
		ID:          auctionEntity.ID,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}
	return auctionOutputDTO, nil
}

func (uc *AuctionUseCase) FindAuctions(
	ctx context.Context,
	status AuctionStatus,
	category, productName string,
) ([]*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntities, err := uc.AuctionRepository.FindAuctions(
		ctx,
		auction_entity.AuctionStatus(status),
		category,
		productName,
	)
	if err != nil {
		return nil, err
	}

	auctionOutputDTOs := make([]*AuctionOutputDTO, 0, len(auctionEntities))
	for _, auctionEntity := range auctionEntities {
		auctionOutputDTO := &AuctionOutputDTO{
			ID:          auctionEntity.ID,
			ProductName: auctionEntity.ProductName,
			Category:    auctionEntity.Category,
			Description: auctionEntity.Description,
			Condition:   ProductCondition(auctionEntity.Condition),
			Status:      AuctionStatus(auctionEntity.Status),
			Timestamp:   auctionEntity.Timestamp,
		}
		auctionOutputDTOs = append(auctionOutputDTOs, auctionOutputDTO)
	}

	return auctionOutputDTOs, nil
}
