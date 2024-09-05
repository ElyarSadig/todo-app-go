package usecase

import (
	"context"
	"net/http"

	"github.com/elyarsadig/todo-app/internal/auth"
	"github.com/elyarsadig/todo-app/internal/models"
	"github.com/elyarsadig/todo-app/pkg/bcrypt"
	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/logger"
	"github.com/elyarsadig/todo-app/pkg/utils"
)

type authUC struct {
	tokenLength int
	authRepo    auth.Repository
	logger      logger.Logger
}

func New(authRepo auth.Repository, log logger.Logger) auth.UseCase {
	return &authUC{
		tokenLength: 10,
		authRepo:    authRepo,
		logger:      log,
	}
}

func (u *authUC) Login(ctx context.Context, user *models.User) (string, error) {
	tempUser, err := u.authRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		errValue, _ := err.(httpErrors.RestErr)
		if errValue.Status() != http.StatusNotFound {
			return "", err
		}
	}
	if !bcrypt.CheckPasswordHash(user.Password, tempUser.Password) {
		return "", httpErrors.NewRestError(http.StatusBadRequest, "invalid credential try again!")
	}
	token, err := utils.GenerateSecureToken(u.tokenLength)
	if err != nil {
		u.logger.Error(err.Error())
		return "", httpErrors.NewRestError(http.StatusInternalServerError, "something went wrong!")
	}
	err = u.authRepo.UpdateUserToken(ctx, tempUser.Email, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUC) Register(ctx context.Context, user *models.User) (string, error) {
	_, err := u.authRepo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return "", httpErrors.NewRestError(http.StatusConflict, "email already in use")
	}
	hashedPassword, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		u.logger.Error(err)
		return "", httpErrors.NewRestError(http.StatusInternalServerError, "something went wrong")
	}
	token, err := utils.GenerateSecureToken(u.tokenLength)
	if err != nil {
		u.logger.Error(err)
		return "", httpErrors.NewRestError(http.StatusInternalServerError, "something went wrong")
	}
	err = u.authRepo.Create(ctx, &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Token:    user.Token,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUC) Logout(ctx context.Context, token string) error {
	return u.authRepo.DeleteUserToken(ctx, token)
}
