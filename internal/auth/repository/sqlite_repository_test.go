package repository

import (
	"context"
	"database/sql/driver"
	"io"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/elyarsadig/todo-app/internal/models"
	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestGetUserByEmail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	authRepo := New(db, logger.New(io.Discard))

	t.Run("GetByEmail Success", func(t *testing.T) {
		email := "elyar@email.com"

		rows := sqlmock.NewRows([]string{"id", "name", "email", "token", "password"}).AddRow(1, "Elyar", email, "token", "password-hash")

		mock.ExpectPrepare(getUserByEmailQuery).ExpectQuery().WithArgs(email).WillReturnRows(rows)

		user, err := authRepo.GetUserByEmail(context.Background(), email)
		require.NoError(t, err)
		require.Equal(t, 1, user.ID)
		require.Equal(t, email, user.Email)
		require.Equal(t, "token", *user.Token)
		require.Equal(t, "password-hash", user.Password)
	})

	t.Run("GetByEmail No Rows", func(t *testing.T) {
		email := "nonexistent@email.com"

		rows := sqlmock.NewRows([]string{"id", "email", "token", "password"})

		mock.ExpectPrepare(getUserByEmailQuery).
			ExpectQuery().
			WithArgs(email).
			WillReturnRows(rows)

		user, err := authRepo.GetUserByEmail(context.Background(), email)
		errValue, ok := err.(httpErrors.RestErr)
		if !ok {
			t.Fatal("type assertion to httpErrors.RestErr failed")
		}
		require.Error(t, err)
		require.Equal(t, http.StatusNotFound, errValue.Status())
		require.Equal(t, 0, user.ID)
		require.Equal(t, "", user.Email)
	})

	t.Run("Prepare Statement Error", func(t *testing.T) {
		email := "error@preparingstmt.com"

		mock.ExpectPrepare(getUserByEmailQuery).
			WillReturnError(httpErrors.NewInternalServerError(nil))

		user, err := authRepo.GetUserByEmail(context.Background(), email)
		errValue, ok := err.(httpErrors.RestErr)
		if !ok {
			t.Fatal("type assertion to httpErrors.RestErr failed")
		}
		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, errValue.Status())
		require.Equal(t, 0, user.ID)
		require.Equal(t, "", user.Email)
	})

	t.Run("Query Execution Error", func(t *testing.T) {
		email := "error@queryexec.com"

		mock.ExpectPrepare(getUserByEmailQuery).
			ExpectQuery().
			WithArgs(email).
			WillReturnError(httpErrors.NewInternalServerError(nil))

		user, err := authRepo.GetUserByEmail(context.Background(), email)
		errValue, ok := err.(httpErrors.RestErr)
		if !ok {
			t.Fatal("type assertion to httpErrors.RestErr failed")
		}
		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, errValue.Status())
		require.Equal(t, 0, user.ID)
		require.Equal(t, "", user.Email)
	})
}

func TestCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	logger := logger.New(io.Discard)
	authRepo := New(db, logger)

	t.Run("Create Success", func(t *testing.T) {
		dummyToken := "dummy"
		obj := models.User{
			Name:     "dummy name",
			Email:    "dummy@email.com",
			Token:    &dummyToken,
			Password: "password",
		}

		mock.ExpectPrepare(createUserQuery).
			ExpectExec().
			WithArgs(obj.Name, obj.Email, obj.Token, obj.Password).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := authRepo.Create(context.Background(), &obj)
		require.NoError(t, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	})
}

func TestDeleteUserToken(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	logger := logger.New(io.Discard)
	authRepo := New(db, logger)

	t.Run("DeleteUserToken Success", func(t *testing.T) {
		mock.ExpectPrepare(deleteUserToken).ExpectExec().WillReturnResult(driver.RowsAffected(1))
		err := authRepo.DeleteUserToken(context.Background(), "token")
		require.NoError(t, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	})
}

func TestUpdateUserToken(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	authRepo := New(db, logger.New(io.Discard))

	t.Run("UpdateUserToken Success", func(t *testing.T) {
		email := "email"
		token := "token"
		mock.ExpectPrepare(updateUserToken).ExpectExec().WithArgs(token, email).WillReturnResult(sqlmock.NewResult(0, 1))
		err := authRepo.UpdateUserToken(context.Background(), email, token)
		require.NoError(t, err)
		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	})
}
