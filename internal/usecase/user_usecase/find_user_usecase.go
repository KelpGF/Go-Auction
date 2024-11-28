package user_usecase

import (
	"context"

	"github.com/KelpGF/Go-Auction/internal/entity/user_entity"
	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type UserOutputDTO struct {
	ID   string
	Name string
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	user, err := u.UserRepository.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
