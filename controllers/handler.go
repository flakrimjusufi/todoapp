package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	db "todoapp/database"
	"todoapp/models"
	"todoapp/samples"
)

func GetSampleToDoList(res http.ResponseWriter, req *http.Request) {
	sampleData := samples.Sample{}

	got, err := sampleData.GetSampleData()
	if err != nil {
		log.Fatalln(err)
	}

	handleError := json.NewEncoder(res).Encode(got)
	if handleError != nil {
		log.Fatalln(handleError)
	}
}

func RenderTemplate(res http.ResponseWriter, req *http.Request) {

	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	t, err := template.ParseFiles("templates/view.html")
	if err != nil {
		log.Fatalln(err)
	}

	sampleData := samples.Sample{}

	got, err := sampleData.GetSampleData()
	if err != nil {
		log.Fatalln(err)
	}

	data := make(map[string]interface{})

	data["toDoList"] = got

	if err != nil {
		log.Println(err)
	}
	handleError := t.Execute(res, data)
	if handleError != nil {
		log.Fatalln(err)
	}
}

func CreateToDoList(res http.ResponseWriter, req *http.Request) {

	var database = db.Connect().Debug()

	todo := models.Todo{}

	err := json.NewDecoder(req.Body).Decode(&todo) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
	}

	database.NewRecord(todo)
	database.Create(&todo)
	database.Close()

	handleError := json.NewEncoder(res).Encode(todo)
	if handleError != nil {
		log.Fatalln(handleError)
	}
}

func GetAllToDos(res http.ResponseWriter, req *http.Request) {
	todos := getAllTodos()
	err := json.NewEncoder(res).Encode(todos)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetToDoById(res http.ResponseWriter, req *http.Request) {

	var database = db.Connect().Debug()

	vars := mux.Vars(req)
	idFromMux := vars["id"]

	var todo models.Todo
	database.Where("id = ?", idFromMux).Find(&todo)
	database.Close()

	err := json.NewEncoder(res).Encode(todo)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateToDoById(w http.ResponseWriter, r *http.Request) {

	var database = db.Connect().Debug()

	todo := models.Todo{}
	vars := mux.Vars(r)
	id := vars["id"]

	database.Where("id =?", id).Find(&todo)

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		panic(err)
	}

	database.Save(&todo)
	database.Close()

	handleError := json.NewEncoder(w).Encode(todo)
	if handleError != nil {
		log.Fatalln(handleError)
	}
}

func DeleteToDo(res http.ResponseWriter, req *http.Request) {

	var database = db.Connect().Debug()

	todo := models.Todo{}
	vars := mux.Vars(req)
	id := vars["id"]

	database.Where("id =?", id).Find(&todo)
	database.Delete(&todo)
	database.Close()

	err := json.NewEncoder(res).Encode("ToDo deleted successfully")
	if err != nil {
		log.Fatalln(err)
	}
}

func RenderCustomToDoTemplate(res http.ResponseWriter, req *http.Request) {

	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	t, err := template.ParseFiles("templates/view.html")
	if err != nil {
		log.Fatalln(err)
	}

	toDoList := getAllTodos()

	data := make(map[string]interface{})

	data["toDoList"] = toDoList

	if err != nil {
		log.Println(err)
	}
	handleError := t.Execute(res, data)
	if handleError != nil {
		log.Fatalln(err)
	}
}

func getAllTodos() []models.Todo {
	var database = db.Connect().Debug()

	var toDoList []models.Todo
	database.Find(&toDoList)
	database.Close()
	return toDoList
}
