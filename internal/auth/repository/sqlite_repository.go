package repository

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/elyarsadig/todo-app/internal/auth"
	"github.com/elyarsadig/todo-app/internal/models"
	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/logger"
)

type authRepo struct {
	db     *sql.DB
	logger logger.Logger
}

func New(db *sql.DB, log logger.Logger) auth.Repository {
	return &authRepo{
		db:     db,
		logger: log,
	}
}

func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	stmt, err := r.db.PrepareContext(ctx, getUserByEmailQuery)
	if err != nil {
		r.logger.Error(err)
		return models.User{}, httpErrors.NewInternalServerError(err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, email)
	var tempUser models.User
	err = row.Scan(&tempUser.ID, &tempUser.Name, &tempUser.Email, &tempUser.Token, &tempUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, httpErrors.NewNotFoundError(err)
		}
		r.logger.Error(err)
		return models.User{}, httpErrors.NewInternalServerError(err)
	}
	return tempUser, nil
}

func (r *authRepo) Create(ctx context.Context, obj *models.User) error {
	stmt, err := r.db.PrepareContext(ctx, createUserQuery)
	if err != nil {
		r.logger.Error(err)
		return httpErrors.NewInternalServerError(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, obj.Name, obj.Email, obj.Token, obj.Password)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint") {
			return httpErrors.NewRestError(http.StatusConflict, "email already in use", err)
		}
		r.logger.Error(err)
		return httpErrors.NewInternalServerError(err)
	}
	return nil
}

func (r *authRepo) DeleteUserToken(ctx context.Context, token string) error {
	stmt, err := r.db.PrepareContext(ctx, deleteUserToken)
	if err != nil {
		r.logger.Error(err)
		return httpErrors.NewInternalServerError(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, token)
	if err != nil {
		r.logger.Error(err)
		return httpErrors.NewInternalServerError(err)
	}
	return nil
}

func (r *authRepo) UpdateUserToken(ctx context.Context, email, token string) error {
	stmt, err := r.db.PrepareContext(ctx, updateUserToken)
	if err != nil {
		r.logger.Error(err)
		return httpErrors.NewInternalServerError(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, token, email)
	if err != nil {
		r.logger.Error(err)
		return httpErrors.NewInternalServerError(err)
	}
	return nil
}
