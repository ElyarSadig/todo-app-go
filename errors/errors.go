package errors

type HandlerError struct {
	err    error
	status int
}

func New(err error, httpStatus int) *HandlerError {
	return &HandlerError{
		err:    err,
		status: httpStatus,
	}
}

func (h HandlerError) Error() string {
	return h.err.Error()
}

func (h HandlerError) Status() int {
	return h.status
}
