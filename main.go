package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todoapp/handler"
)


func main() {
	fmt.Println("Starting server in localhost:8088")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getTodoList", handler.GetTodoList).Methods("GET")
	log.Fatal(http.ListenAndServe(":8088", myRouter))
}
