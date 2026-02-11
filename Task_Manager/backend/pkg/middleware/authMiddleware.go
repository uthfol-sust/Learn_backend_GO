package middleware

import (
	"context"
	"net/http"
	"taskmanager/pkg/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyJWT(authHeader[len("Bearer "):])

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), utils.RoleKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
