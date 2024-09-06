package http

import (
	"net/http"

	"github.com/elyarsadig/todo-app/internal/auth"
	"github.com/elyarsadig/todo-app/pkg/logger"
	"github.com/elyarsadig/todo-app/pkg/utils"
	"github.com/nahojer/httprouter"
)

type authHandlers struct {
	authUC auth.UseCase
	logger logger.Logger
}

func NewAuthHandlers(authUC auth.UseCase, log logger.Logger) auth.Handler {
	return &authHandlers{
		authUC: authUC,
		logger: log,
	}
}

func (h *authHandlers) Register() httprouter.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func (h *authHandlers) Login() httprouter.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func (h *authHandlers) Logout() httprouter.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		token := ctx.Value(utils.TokenCtxKey{}).(string)
		err := h.authUC.Logout(ctx, token)
		if err != nil {
			return err
		}
		return nil
	}
}
