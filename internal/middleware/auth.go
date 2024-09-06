package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/elyarsadig/todo-app/pkg/httpErrors"
	"github.com/elyarsadig/todo-app/pkg/utils"
)

func Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if len(bearerToken) == 0 {
			httpErrors.ReturnError(w, httpErrors.NewRestError(http.StatusUnauthorized, "unauthorized", nil))
			return
		}
		if !strings.Contains(bearerToken, "Bearer ") {
			httpErrors.ReturnError(w, httpErrors.NewRestError(http.StatusUnauthorized, "unauthorized", nil))
			return
		}
		token := bearerToken[len("Bearer "):]
		if len(token) == 0 {
			httpErrors.ReturnError(w, httpErrors.NewRestError(http.StatusUnauthorized, "unauthorized", nil))
			return
		}
		ctx := context.WithValue(r.Context(), utils.TokenCtxKey{}, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
