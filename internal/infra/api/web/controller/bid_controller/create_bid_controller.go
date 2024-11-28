package bid_controller

import (
	"net/http"

	"github.com/KelpGF/Go-Auction/config/rest_err"
	"github.com/KelpGF/Go-Auction/internal/infra/api/web/validation"
	"github.com/KelpGF/Go-Auction/internal/usecase/bid_usecase"
	"github.com/gin-gonic/gin"
)

func (controller *bidController) Create(c *gin.Context) {
	var bidInputDTO *bid_usecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := controller.bidUseCase.CreateBid(c, bidInputDTO)
	if err != nil {
		restErr := rest_err.ConvertErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bid created"})
}
