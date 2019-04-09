package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/dugiahuy/qna/models"
	"github.com/gorilla/mux"
)

// Index ..
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, my first RESTful API server!")
}

// GetQuestions ..
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	question := []models.Question{}

	question = append(question, models.Question{ID: 1, Content: "What your name? 1"})
	question = append(question, models.Question{ID: 2, Content: "What your name? 2"})
	question = append(question, models.Question{ID: 3, Content: "What your name? 3"})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(question); err != nil {
		panic(err)
	}
}

// GetQuestionID ..
func GetQuestionID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionID := vars["question_id"]
	fmt.Fprintln(w, "Question Content!", questionID)
}
