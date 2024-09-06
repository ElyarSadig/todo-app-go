package http

import (
	"net/http"

	"github.com/elyarsadig/todo-app/internal/auth"
	"github.com/elyarsadig/todo-app/internal/models"
	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/logger"
	"github.com/elyarsadig/todo-app/pkg/utils"
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

func (h *authHandlers) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req RegisterRequest
	ok := utils.UnmarshalRequest(w, r, &req)
	if ok {
		token, err := h.authUC.Register(ctx, &models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			httpErrors.ReturnError(w, err)
			return
		}
		res := LoginResponse{
			Token: token,
		}
		httpErrors.ReturnSuccess(w, res)
	}
}

func (h *authHandlers) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req LoginRequest
	ok := utils.UnmarshalRequest(w, r, &req)
	if ok {
		token, err := h.authUC.Login(ctx, &models.User{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			httpErrors.ReturnError(w, err)
			return
		}
		res := LoginResponse{
			Token: token,
		}
		httpErrors.ReturnSuccess(w, res)
	}
}

func (h *authHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := ctx.Value(utils.TokenCtxKey{}).(string)
	err := h.authUC.Logout(ctx, token)
	if err != nil {
		httpErrors.ReturnError(w, err)
		return
	}
	httpErrors.ReturnSuccess(w, nil)
}
