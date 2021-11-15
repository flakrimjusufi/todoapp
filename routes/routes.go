package routes

import (
	"github.com/labstack/echo/v4"
	"todoapp/controllers"
)

func HandleRoutes(route *echo.Echo) {

	route.GET("/", controllers.RenderTemplate)
	route.GET("/viewToDoList", controllers.RenderCustomToDoTemplate)
	route.GET("/getSampleTodoList", controllers.GetSampleToDoList)
	route.POST("/toDoList", controllers.CreateToDoList)
	route.GET("/toDoList", controllers.GetAllToDos)
	route.GET("/toDoList/:id", controllers.GetToDoById)
	route.PUT("/toDoList/:id", controllers.UpdateToDoById)
	route.DELETE("/toDoList/:id", controllers.DeleteToDo)
}
