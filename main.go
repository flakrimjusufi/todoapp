package main

import (
	"fmt"
	"log"
	"net/http"
	"todoapp/routes"
)

func main() {
	fmt.Println("Starting server in http://localhost:8088")
	err := http.ListenAndServe(":8088", routes.HandleRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
