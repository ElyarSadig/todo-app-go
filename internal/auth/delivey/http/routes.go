package http

import (
	"github.com/elyarsadig/todo-app/internal/auth"
	"github.com/nahojer/httprouter"
)

func MapAuthRoutes(authGroup *httprouter.Router, h auth.Handler) {
	authGroup.Handle("post", "/register", h.Register())
	authGroup.Handle("post", "/login", h.Login())
	authGroup.Handle("post", "/logout", h.Logout())
}