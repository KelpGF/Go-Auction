package auction_controller

import (
	"net/http"

	"github.com/KelpGF/Go-Auction/config/rest_err"
	"github.com/KelpGF/Go-Auction/internal/infra/api/web/validation"
	"github.com/KelpGF/Go-Auction/internal/usecase/auction_usecase"
	"github.com/gin-gonic/gin"
)

func (controller *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO *auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	id, err := controller.auctionUseCase.CreateAuction(c, auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
