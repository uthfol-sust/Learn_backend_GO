package middleware

import (
	"fmt"
	"net/http"
)

// ✅ CORS + Preflight Middleware
func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		

		fmt.Println("I am Test 2 middleware")
		next.ServeHTTP(w, r)
	})
}