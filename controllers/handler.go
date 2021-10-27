package controllers

import (
	"encoding/json"
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
