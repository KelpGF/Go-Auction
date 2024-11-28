package rest_err

import (
	"net/http"

	"github.com/KelpGF/Go-Auction/internal/internal_error"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *RestErr) Error() string {
	return e.Message
}

func NewBadRequestError(message string, causes ...Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}

func ConvertErr(internal_error *internal_error.InternalError) *RestErr {
	switch internal_error.Err {
	case "bad_request":
		return NewBadRequestError(internal_error.Error())
	case "not_found":
		return NewNotFoundError(internal_error.Error())
	default:
		return NewInternalServerError(internal_error.Error(), internal_error)
	}
}
