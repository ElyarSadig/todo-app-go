package httpErrors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Status() int
	Error() string
	Causes() any
}

type RestError struct {
	ErrStatus int    `json:"status"`
	ErrError  string `json:"error"`
	ErrCauses any    `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - error: %s causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() any {
	return e.ErrCauses
}

func NewRestError(status int, err string, causes any) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewInternalServerError(causes any) RestErr {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError: "internal server error",
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes any) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError: "not found",
		ErrCauses: causes,
	}
}
