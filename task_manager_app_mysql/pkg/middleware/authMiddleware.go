package middleware

import (
	"fmt"
	"net/http"
	"taskmanager/pkg/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("This 2 midleware\n")
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		sys_role,err := utils.VerifyJWT(authHeader[len("Bearer "):])

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		 userRole := utils.UserRole(sys_role)
		 fmt.Println(userRole)
		 next.ServeHTTP(w, r)
	})
}
