package usecase

import (
	"context"
	"io"
	"testing"

	"github.com/elyarsadig/todo-app/internal/auth/mock"
	"github.com/elyarsadig/todo-app/internal/models"
	"github.com/elyarsadig/todo-app/pkg/bcrypt"
	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := mock.NewMockRepository(ctrl)
	logger := logger.New(io.Discard)
	authUC := New(mockAuthRepo, logger)

	ctx := context.Background()
	user := models.User{
		Email:    "test@email.com",
		Password: "test123",
	}
	hashedPass, err := bcrypt.HashPassword(user.Password)
	require.NoError(t, err)
	mockUser := models.User{
		Email:    "test@email.com",
		Password: hashedPass,
	}
	mockAuthRepo.EXPECT().GetUserByEmail(ctx, user.Email).Return(mockUser, nil)
	mockAuthRepo.EXPECT().UpdateUserToken(ctx, user.Email, gomock.Any()).Return(nil)
	token, err := authUC.Login(ctx, &user)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestRegister(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthRepo := mock.NewMockRepository(ctrl)
	logger := logger.New(io.Discard)
	authUC := New(mockAuthRepo, logger)

	ctx := context.Background()
	user := models.User{
		Name:     "test",
		Password: "test123",
		Email:    "test@email.com",
	}
	mockAuthRepo.EXPECT().GetUserByEmail(ctx, "test@email.com").Return(models.User{}, httpErrors.UserNotFoundError)
	mockAuthRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)
	token, err := authUC.Register(ctx, &user)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestLogOut(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := mock.NewMockRepository(ctrl)
	logger := logger.New(io.Discard)
	authUC := New(mockAuthRepo, logger)
	ctx := context.Background()
	token := "secret-token"
	mockAuthRepo.EXPECT().DeleteUserToken(ctx, token).Return(nil)
	err := authUC.Logout(ctx, token)
	require.NoError(t, err)
}