package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type TokenCtxKey struct{}

type errors struct {
	Errors map[string]string `json:"errors"`
}

func UnmarshalRequest(w http.ResponseWriter, r *http.Request, payload any) bool {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		httpErrors.ReturnError(w, httpErrors.NewRestError(http.StatusBadRequest, "could not unmarshal values", err))
		return false
	}
	err = validate.Struct(payload)
	if err != nil {
		res := errors{
			Errors: translateValidationErrors(err),
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return false
	}
	return true
}

func translateValidationErrors(err error) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	errors := make(map[string]string)

	for _, fieldError := range validationErrors {
		var message string

		switch fieldError.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", strings.ToLower(fieldError.Field()))
		case "email":
			message = fmt.Sprintf("%s must be a valid email address", strings.ToLower(fieldError.Field()))
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters long", strings.ToLower(fieldError.Field()), fieldError.Param())
		default:
			message = fmt.Sprintf("%s is invalid", strings.ToLower(fieldError.Field()))
		}

		errors[strings.ToLower(fieldError.Field())] = message
	}

	return errors
}
