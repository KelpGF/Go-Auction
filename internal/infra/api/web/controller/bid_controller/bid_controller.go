package bid_controller

import "github.com/KelpGF/Go-Auction/internal/usecase/bid_usecase"

type BidController struct {
	bidUseCase bid_usecase.BidUsecaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUsecaseInterface) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
	}
}
