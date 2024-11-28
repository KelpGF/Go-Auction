package bid_controller

import (
	"net/http"

	"github.com/KelpGF/Go-Auction/config/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller *BidController) FindByAuctionId(c *gin.Context) {
	auctionId := c.Param("auction_id")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid auction id", rest_err.Causes{
			Field:   "id",
			Message: "auction id must be a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	bids, err := controller.bidUseCase.FindBidByAuctionId(c, auctionId)
	if err != nil {
		errRest := rest_err.ConvertErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bids)
}
