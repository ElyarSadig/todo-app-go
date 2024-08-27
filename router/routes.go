package router

import (
	"net/http"

	"github.com/elyarsadig/todo-app/errors"
	"github.com/nahojer/httprouter"
)

type Router struct {
	*httprouter.Router
}

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

func NewV1() Router {
	router := httprouter.New()
	v1 := router.Group("apis/v1")
	v1.ErrorHandler = ErrorHandler
	return Router{Router: v1}
}

func (r *Router) AddHandler(method HttpMethod, path string, handler httprouter.Handler, mw ...httprouter.Middleware) {
	r.Handle(string(method), path, handler, mw...)
}

func (r *Router) AddMiddlewares(mw ...httprouter.Middleware) {
	r.Use(mw...)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	errorHandler, ok := err.(errors.HandlerError)
	if !ok || errorHandler.Status() == http.StatusInternalServerError {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(errorHandler.Status())
	w.Write([]byte(err.Error()))
}
