package routes

import (
	"net/http"
	"taskmanager/pkg/controllers"
	"taskmanager/pkg/middleware"
)

func HandleTaskRoutes(router *http.ServeMux, m *middleware.Manager) {
	router.Handle("POST /tasks",
		m.With(http.HandlerFunc(controllers.CreateTask), middleware.AuthMiddleware),
	)

	router.Handle("GET /tasks",
		m.With(http.HandlerFunc(controllers.GetAllTasks), middleware.AuthMiddleware),
	)

	router.Handle("GET /tasks/{id}",
		m.With(http.HandlerFunc(controllers.GetTaskByID), middleware.AuthMiddleware),
	)

	router.Handle("PUT /tasks/{id}",
		m.With(http.HandlerFunc(controllers.UpdateTask), middleware.AuthMiddleware),
	)

	router.Handle("DELETE /tasks/{id}",
		m.With(http.HandlerFunc(controllers.DeleteTask), middleware.AuthMiddleware),
	)
}
