package middleware

import (
	"context"
	"net/http"

	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/utils"
	"github.com/nahojer/httprouter"
)

func Protected(next httprouter.Handler) httprouter.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		bearerToken := r.Header.Get("Authorization")
		if len(bearerToken) == 0 {
			httpErrors.ErrorHandler(w, r, httpErrors.NewRestError(http.StatusUnauthorized, "unauthorized", nil))
			return nil
		}
		token := bearerToken[len("Bearer "):]
		if len(token) == 0 {
			httpErrors.ErrorHandler(w, r, httpErrors.NewRestError(http.StatusUnauthorized, "unauthorized", nil))
			return nil
		}
		ctx := context.WithValue(r.Context(), utils.TokenCtxKey{}, token)
		r = r.WithContext(ctx)
		next(w, r)
		return nil
	}
}
