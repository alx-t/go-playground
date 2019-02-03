package routers

import (
	"github.com/alx-t/go-playground/controllers"
	"github.com/gorilla/mux"
)

func SetProjectRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Root).Methods("GET")

	router.HandleFunc("/projects", controllers.GetProjects).Methods("GET")
	router.HandleFunc("/projects", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", controllers.GetProjectById).Methods("GET")
	router.HandleFunc("/projects/{id}", controllers.UpdateProject).Methods("PATCH")
	router.HandleFunc("/projects/{id}", controllers.DeleteProject).Methods("DELETE")

	return router
}
