package routes

import (
	"net/http"
	"taskmanager/pkg/controllers"
	"taskmanager/pkg/middleware"
)

func RegisterRoutes(router *http.ServeMux) {

	manger := middleware.Manger{}
	manger.Use(middleware.CorsMiddleware)

	//user handling
	router.Handle("POST /signup",
		manger.With(
			http.HandlerFunc(controllers.SignUp),
		))
	router.Handle("POST /login",
		manger.With(
			http.HandlerFunc(controllers.Login),
		))
	router.Handle("GET /verify",
		manger.With(
			http.HandlerFunc(controllers.VerifyEmail),
		))

	router.Handle("GET /users",
		manger.With(
			http.HandlerFunc(controllers.GetAllUsers),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /users/{id}",
		manger.With(
			http.HandlerFunc(controllers.GetUserByID),
			middleware.AuthMiddleware,
		))
	router.Handle("PUT /users/{id}",
		manger.With(
			http.HandlerFunc(controllers.UpdateUser),
			middleware.AuthMiddleware,
		))
	router.Handle("DELETE /users/{id}",
		manger.With(
			http.HandlerFunc(controllers.DeleteUser),
			middleware.AuthMiddleware,
		))

	//task handing
	router.Handle("POST /tasks",
		manger.With(
			http.HandlerFunc(controllers.CreateTask),
			middleware.AuthMiddleware,
		))
	router.Handle("GET /tasks",
		manger.With(
			http.HandlerFunc(controllers.GetAllTasks),
			middleware.AuthMiddleware,
		))
	router.Handle("GET /tasks/{id}", manger.With(
		http.HandlerFunc(controllers.GetTaskByID),
		middleware.AuthMiddleware,
	))
	router.Handle("PUT /tasks/{id}",
		manger.With(http.HandlerFunc(controllers.UpdateTask),
			middleware.AuthMiddleware,
		))
	router.Handle("DELETE /tasks/{id}",
		manger.With(http.HandlerFunc(controllers.DeleteTask),
			middleware.AuthMiddleware,
		))
}
