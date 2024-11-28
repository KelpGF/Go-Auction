package bid_controller

import "github.com/KelpGF/Go-Auction/internal/usecase/bid_usecase"

type bidController struct {
	bidUseCase bid_usecase.BidUsecaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUsecaseInterface) *bidController {
	return &bidController{
		bidUseCase: bidUseCase,
	}
}
