package auction_controller

import (
	"net/http"
	"strconv"

	"github.com/KelpGF/Go-Auction/config/rest_err"
	"github.com/KelpGF/Go-Auction/internal/usecase/auction_usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller *auctionController) FindById(c *gin.Context) {
	auctionId := c.Param("id")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid auction id", rest_err.Causes{
			Field:   "id",
			Message: "auction id must be a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auction, err := controller.auctionUseCase.FindAuctionByID(c, auctionId)
	if err != nil {
		errRest := rest_err.ConvertErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (controller *auctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("product_name")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("invalid status", rest_err.Causes{
			Field:   "status",
			Message: "status must be a number",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := controller.auctionUseCase.FindAuctions(
		c,
		auction_usecase.AuctionStatus(statusNumber),
		category,
		productName,
	)
	if err != nil {
		errRest := rest_err.ConvertErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (controller *auctionController) FindWinningBidByAuctionID(c *gin.Context) {
	auctionId := c.Param("id")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid auction id", rest_err.Causes{
			Field:   "id",
			Message: "auction id must be a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auction, err := controller.auctionUseCase.FindWinningBidByAuctionID(c, auctionId)
	if err != nil {
		errRest := rest_err.ConvertErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auction)
}
