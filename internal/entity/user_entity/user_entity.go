package user_entity

import (
	"context"

	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserByID(ctx context.Context, id string) (*User, *internal_error.InternalError)
}
