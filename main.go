package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todoapp/routes"

	//"log"
	//"net/http"
	"os"
)

func main() {

	// Echo instance
	e := echo.New()
	if os.Getenv("DB_USERNAME") == "" {
		e := godotenv.Load() //Load .env file for local environment
		if e != nil {
			fmt.Print(e)
		}
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.HandleRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))

	fmt.Println("Starting server in http://localhost:" + os.Getenv("SERVER_PORT"))
	//err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), routes.HandleRoutes())
	//if err != nil {
	//	log.Fatal(err)
	//}
}
