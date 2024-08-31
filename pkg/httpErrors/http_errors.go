package httpErrors

import "fmt"

type RestErr interface {
	Status() int
	Error() string
}

type RestError struct {
	ErrStatus int    `json:"status"`
	ErrError  string `json:"error"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - error: %s", e.ErrStatus, e.ErrError)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func NewRestError(status int, err string) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
	}
}
