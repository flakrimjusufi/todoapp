package routes

import (
	"github.com/gorilla/mux"
	"todoapp/controllers"
)

func HandleRoutes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/getTodoList", controllers.GetTodoList).Methods("GET")
	//route.HandleFunc("/", controllers.Show)

	return route
}
