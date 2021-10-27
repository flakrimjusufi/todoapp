package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"todoapp/controllers"
)

func HandleRoutes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.RenderTemplate)
	route.HandleFunc("/getTodoList", controllers.GetTodoList).Methods("GET")
	route.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	return route
}
