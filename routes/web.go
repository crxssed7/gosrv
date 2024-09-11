package routes

import (
	"gosrv/handlers"
	"gosrv/handlers/person_resource"
	"gosrv/middleware"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	// Define routes
	router.HandleFunc("/", handlers.HomeHandler)

	person_resource.Register(router)

	return router
}
