package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todoapp/routes"
)

func main() {
	fmt.Println("Starting server in http://localhost:" + os.Getenv("SERVER_PORT"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), routes.HandleRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
