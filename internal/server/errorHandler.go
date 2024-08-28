package server

import (
	"net/http"

	"github.com/elyarsadig/todo-app/pkg/httpErrors"
)

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	errorHandler, ok := err.(httpErrors.RestErr)
	if !ok || errorHandler.Status() == http.StatusInternalServerError {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(errorHandler.Status())
	w.Write([]byte(err.Error()))
}
