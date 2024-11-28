package auction_controller

import "github.com/KelpGF/Go-Auction/internal/usecase/auction_usecase"

type auctionController struct {
	auctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *auctionController {
	return &auctionController{
		auctionUseCase: auctionUseCase,
	}
}
