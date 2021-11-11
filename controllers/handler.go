package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	db "todoapp/database"
	"todoapp/models"
	"todoapp/samples"
)

var database = db.Connect().Debug()

type Template struct {
	templates map[string]*template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: make(map[string]*template.Template),
	}
}

func (tmpl *Template) Render(w io.Writer, htmlName string, data interface{}, c echo.Context) error {
	if t, exist := tmpl.templates[htmlName]; exist { //Check if template exists
		return t.Execute(w, data)
	} else {
		return errors.New("There is no " + htmlName + " in Template map.")
	}
}

// Add Set template on htmlName so that the 'tmpl' will be okay in 'RenderTemplate' function.
func (tmpl *Template) Add(htmlName string, template *template.Template) {
	tmpl.templates[htmlName] = template
}

func GetSampleToDoList(c echo.Context) error {
	sampleData := samples.Sample{}

	got, err := sampleData.GetSampleData()
	if err != nil {
		log.Fatalln(err)
	}
	return c.JSON(http.StatusOK, got)
}

func RenderTemplate(c echo.Context) error {

	sampleData := samples.Sample{}

	got, _ := sampleData.GetSampleData()

	renderHtml := NewTemplate()
	renderHtml.Add("templates/view.html", template.Must(template.ParseGlob("templates/*.html")))

	c.Echo().Renderer = renderHtml

	return c.Render(http.StatusOK, "templates/view.html", map[string]interface{}{
		"toDoList": got,
	})
}

func CreateToDoList(c echo.Context) error {

	todo := &models.Todo{}

	if err := c.Bind(todo); err != nil {
		return err
	}
	database.NewRecord(todo)
	database.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func GetAllToDos(c echo.Context) error {
	todos := getAllTodos()
	return c.JSON(http.StatusOK, todos)
}

func GetToDoById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var todo models.Todo
	result := database.Where("id = ?", id).Find(&todo)

	if result.RecordNotFound() {
		return c.JSON(http.StatusNotFound, "ToDoList not found!")
	}
	return c.JSON(http.StatusOK, result.Value)
}

func UpdateToDoById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	userInput := new(models.Todo)
	if handleError := c.Bind(userInput); handleError != nil {
		return handleError
	}

	var todoList models.Todo
	result := database.Where("id =?", id).Find(&todoList)
	if result.RecordNotFound() {
		return c.JSON(http.StatusNotFound, "ToDoList not found!")
	}

	todoList.Labels = userInput.Labels
	todoList.Comments = userInput.Comments
	todoList.DueDate = userInput.DueDate

	database.Save(&todoList)
	return c.JSON(http.StatusOK, todoList)
}

func DeleteToDo(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var todo models.Todo
	result := database.Where("id =?", id).Find(&todo)
	if result.RecordNotFound() {
		return c.JSON(http.StatusNotFound, "ToDoList not found!")
	}
	database.Delete(&todo)
	return c.NoContent(http.StatusNoContent)
}

func RenderCustomToDoTemplate(c echo.Context) error {

	toDoList := getAllTodos()
	data := make(map[string]interface{})
	data["toDoList"] = toDoList

	renderHtml := NewTemplate()
	renderHtml.Add("templates/view.html", template.Must(template.ParseGlob("templates/*.html")))

	c.Echo().Renderer = renderHtml

	return c.Render(http.StatusOK, "templates/view.html", map[string]interface{}{
		"toDoList": toDoList,
	})
}

func getAllTodos() []models.Todo {
	var toDoList []models.Todo
	database.Find(&toDoList)
	return toDoList
}
