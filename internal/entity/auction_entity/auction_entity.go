package auction_entity

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

func NewAuction(
	productName, category, description string,
	condition ProductCondition,
) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		ID:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	err := auction.Validate()
	if err != nil {
		return nil, err
	}

	return auction, nil
}

func (a *Auction) Validate() *internal_error.InternalError {
	if len(a.ProductName) <= 1 {
		return internal_error.NewBadRequestError("product name must be longer than 1 character")
	}

	if len(a.Category) <= 1 {
		return internal_error.NewBadRequestError("category must be longer than 1 character")
	}

	if len(a.Description) <= 10 {
		return internal_error.NewBadRequestError("description must be longer than 10 character")
	}

	if a.Condition < New || a.Condition > Refurbished {
		return internal_error.NewBadRequestError("invalid condition")
	}

	return nil
}

type ProductCondition int

const (
	New = iota
	Used
	Refurbished
)

type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string,
	) ([]*Auction, *internal_error.InternalError)
}
