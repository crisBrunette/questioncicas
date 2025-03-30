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

func validateAnswerA(w http.ResponseWriter, req *http.Request) {
	if validator.Validator("a") {
		fmt.Fprintln(w, "Correct answer!")
	} else {
		fmt.Fprintln(w, "Oh! Incorrect answer... :( Try again!")
	}
}

func validateAnswerB(w http.ResponseWriter, req *http.Request) {
	if validator.Validator("b") {
		fmt.Fprintln(w, "Correct answer!")
	} else {
		fmt.Fprintln(w, "Oh! Incorrect answer... :( Try again!")
	}
}

func validateAnswerC(w http.ResponseWriter, req *http.Request) {
	if validator.Validator("c") {
		fmt.Fprintln(w, "Correct answer!")
	} else {
		fmt.Fprintln(w, "Oh! Incorrect answer... :( Try again!")
	}
}

func main() {
	fmt.Println("Questioncicas is ready...")

	http.HandleFunc("/beAsked", getQuestion)

	http.HandleFunc("/a", validateAnswerA) // http.HandleFunc("/answer(a)", validateAnswer) // parametros en http getrequest
	http.HandleFunc("/b", validateAnswerB)
	http.HandleFunc("/c", validateAnswerC)

	http.ListenAndServe(":8070", nil)
}
