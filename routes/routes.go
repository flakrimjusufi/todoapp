package routes

import (
	"github.com/gorilla/mux"
	"todoapp/controllers"
)

func HandleRoutes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.RenderTemplate)
	route.HandleFunc("/viewToDoList", controllers.RenderCustomToDoTemplate)
	route.HandleFunc("/getSampleTodoList", controllers.GetSampleToDoList).Methods("GET")
	route.HandleFunc("/toDoList", controllers.CreateToDoList).Methods("POST")
	route.HandleFunc("/toDoList", controllers.GetAllToDos).Methods("GET")
	route.HandleFunc("/toDoList/{id}", controllers.GetToDoById).Methods("GET")
	route.HandleFunc("/toDoList/{id}", controllers.UpdateToDoById).Methods("PUT")
	route.HandleFunc("/toDoList/{id}", controllers.DeleteToDo).Methods("DELETE")
	return route
}
