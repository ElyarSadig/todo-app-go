package server

import (
	authHttp "github.com/elyarsadig/todo-app/internal/auth/delivery/http"
	authRepository "github.com/elyarsadig/todo-app/internal/auth/repository"
	authUseCase "github.com/elyarsadig/todo-app/internal/auth/usecase"
	"github.com/elyarsadig/todo-app/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func (s *Server) MapHandlers(router chi.Router) error {
	// init repository
	authRepo := authRepository.New(s.db, s.logger)
	// init usecase
	authUC := authUseCase.New(authRepo, s.logger)
	// init handlers
	authHandlers := authHttp.NewAuthHandlers(authUC, s.logger)
	// setup middlewares
	router.Use(middleware.CORS)

	authHttp.MapAuthRoutes(router, authHandlers)

	router.Mount("/apis/v1", router)
	return nil
}
