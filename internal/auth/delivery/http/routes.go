package http

import (
	"github.com/elyarsadig/todo-app/internal/auth"
	"github.com/elyarsadig/todo-app/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func MapAuthRoutes(router chi.Router, h auth.Handler) {
	router.Post("/register", h.Register)
	router.Post("/login", h.Login)
	router.With(middleware.Protected).Post("/logout", h.Logout)
}
