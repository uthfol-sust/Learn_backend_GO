package routes

import (
	"net/http"
	"taskmanager/pkg/controllers"
	"taskmanager/pkg/middleware"
)

func HandleUserRoutes(router *http.ServeMux, m *middleware.Manager) {

	router.Handle("POST /signup",
		m.With(http.HandlerFunc(controllers.SignUp)),
	)

	router.Handle("POST /login",
		m.With(http.HandlerFunc(controllers.Login)),
	)

	router.Handle("GET /verify",
		m.With(http.HandlerFunc(controllers.VerifyEmail)),
	)

	router.Handle("GET /users",
		m.With(http.HandlerFunc(controllers.GetAllUsers), middleware.AuthMiddleware),
	)

	router.Handle("GET /users/{id}",
		m.With(http.HandlerFunc(controllers.GetUserByID), middleware.AuthMiddleware),
	)

	router.Handle("PUT /users/{id}",
		m.With(http.HandlerFunc(controllers.UpdateUser), middleware.AuthMiddleware),
	)

	router.Handle("DELETE /users/{id}",
		m.With(http.HandlerFunc(controllers.DeleteUser), middleware.AuthMiddleware),
	)

	router.Handle("POST /users/{id}",
		m.With(http.HandlerFunc(controllers.UpdateUserPassword), middleware.AuthMiddleware),
	)

	router.Handle("PUT /passwordcode/{id}",
		m.With(http.HandlerFunc(controllers.ResetCodeVerification), middleware.AuthMiddleware),
	)
}
