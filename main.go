package main

import (
	"fmt"
	"net/http"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
	"github.com/crisBrunette/questioncicas/pkg/validator"
)

func getQuestion(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, question_generator.Generator(req))
}

func validateAnswer(w http.ResponseWriter, req *http.Request) {
	value := req.FormValue("answer")
	if validator.Validator(value) {
		fmt.Fprintln(w, "Correct answer!")
	} else {
		fmt.Fprintf(w, "Ouch! your answer %s is incorrect... :( Try again!\n", value)
	}
}

func main() {
	fmt.Println("Questioncicas is ready...")

	http.HandleFunc("/beAsked", getQuestion)

	http.HandleFunc("/validate", validateAnswer)

	http.ListenAndServe(":9010", nil)
}

// preguntar a oscar si paso un slice en la query. Example localhost:9050/validate?answer=[a,b,c]&test=3?
