package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"todoapp/samples"
)

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	sampleData := samples.Sample{}

	got, err := sampleData.GetSampleData()
	if err != nil {
		log.Fatalln(err)
	}

	handleError := json.NewEncoder(w).Encode(got)
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
	_ = t.Execute(res, data)
}
