package auth

import (
	"context"

	"github.com/elyarsadig/todo-app/internal/models"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	Create(ctx context.Context, obj *models.User) error
	UpdateUserToken(ctx context.Context, email, token string) error
	DeleteUserToken(ctx context.Context, token string) error
}
