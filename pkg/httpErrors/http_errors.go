package httpErrors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Status() int
	Error() string
	Causes() any
	ErrorValue() string
}

type restError struct {
	ErrStatus int    `json:"status"`
	ErrError  string `json:"error"`
	ErrCauses any    `json:"-"`
}

func (e restError) Error() string {
	return fmt.Sprintf("status: %d - error: %s causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restError) ErrorValue() string {
	return e.ErrError
}

func (e restError) Status() int {
	return e.ErrStatus
}

func (e restError) Causes() any {
	return e.ErrCauses
}

func NewRestError(status int, err string, causes any) RestErr {
	return restError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewInternalServerError(causes any) RestErr {
	return restError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  "internal server error",
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes any) RestErr {
	return restError{
		ErrStatus: http.StatusNotFound,
		ErrError:  "not found",
		ErrCauses: causes,
	}
}
