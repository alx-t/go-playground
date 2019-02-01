package routers

import (
	"github.com/gorilla/mux"
	"github.com/alx-t/go-playground/controllers"
)

func SetProjectRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Root).Methods("GET")
	return router
}
