package server

import (
	authHttp "github.com/elyarsadig/todo-app/internal/auth/delivery/http"
	authRepository "github.com/elyarsadig/todo-app/internal/auth/repository"
	authUseCase "github.com/elyarsadig/todo-app/internal/auth/usecase"
	"github.com/elyarsadig/todo-app/internal/middleware"
	"github.com/nahojer/httprouter"
)

func (s *Server) MapHandlers(router *httprouter.Router) error {
	// init repository
	authRepo := authRepository.New(s.db, s.logger)
	// init usecase
	authUC := authUseCase.New(authRepo, s.logger)
	// init handlers
	authHandlers := authHttp.NewAuthHandlers(authUC, s.logger)
	// setup middlewares
	router.Use(middleware.CORS)

	v1 := router.Group("/apis/v1")
	authGroup := v1.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandlers)

	return nil
}
