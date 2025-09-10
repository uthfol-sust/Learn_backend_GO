package routes

import (
	"net/http"
	"taskmanager/pkg/middleware"
)

func RegisterRoutes(router *http.ServeMux) {
	manager := &middleware.Manager{}
	manager.Use(middleware.CorsMiddleware)

	HandleUserRoutes(router, manager)
	HandleTaskRoutes(router, manager)
}
