package auth

import (
	"context"

	"github.com/elyarsadig/todo-app/internal/models"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	Create(ctx context.Context, obj *models.User) error
	Update(ctx context.Context, obj *models.User) (models.User, error)
	DeleteUserToken(ctx context.Context, token string) error
}
