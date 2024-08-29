package auth

import (
	"context"

	"github.com/elyarsadig/todo-app/internal/models"
)

type UseCase interface {
	Login(ctx context.Context, user *models.User) (string, error)
	Logout(ctx context.Context, token string) error
	Register(ctx context.Context, user *models.User) (string, error)
}
