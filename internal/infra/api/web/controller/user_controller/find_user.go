package user_controller

import (
	"net/http"

	"github.com/KelpGF/Go-Auction/config/rest_err"
	"github.com/KelpGF/Go-Auction/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (controller *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("id")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid user id", rest_err.Causes{
			Field:   "id",
			Message: "user id must be a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	user, err := controller.userUseCase.FindUserById(c, userId)
	if err != nil {
		errRest := rest_err.ConvertErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, user)
}
