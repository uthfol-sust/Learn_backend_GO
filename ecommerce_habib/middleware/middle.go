package middleware

import (
	"fmt"
	"net/http"
)

// ✅ CORS + Preflight Middleware
func MiddleTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am Test 3 middleware")

		next.ServeHTTP(w, r)
	})
}
