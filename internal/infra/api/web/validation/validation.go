package validation

import (
	"encoding/json"
	"errors"

	"github.com/KelpGF/Go-Auction/config/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewNotFoundError("invalid json body")
	}

	if errors.As(validationErr, &jsonValidation) {
		errorCauses := []rest_err.Causes{}

		for _, fieldErr := range validationErr.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   fieldErr.Field(),
				Message: fieldErr.Translate(transl),
			})
		}

		return rest_err.NewBadRequestError("Invalid field values", errorCauses...)
	}

	return rest_err.NewBadRequestError("Error trying to convert fields")
}
