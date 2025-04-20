package main

import (
	"fmt"
	"net/http"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
	"github.com/crisBrunette/questioncicas/pkg/validator"
)

var QuestionGenerator *question_generator.QuestionGenerator
var Validator *validator.Validator

func getQuestion(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, QuestionGenerator.Generator(req))
}

func validateAnswer(w http.ResponseWriter, req *http.Request) {
	value := req.FormValue("answer")
	if Validator.IsValid(value) {
		fmt.Fprintln(w, "Correct answer!")
	} else {
		fmt.Fprintf(w, "Ouch! your answer %s is incorrect... :( Try again!\n", value)
	}
}

func setUp() {
	myQuestionList := question_generator.IntiQuestionList()
	myQuestionList.PopulateQuestions()

	QuestionGenerator = question_generator.InitQuestionGenerator(&myQuestionList)
	Validator = validator.InitValidator(QuestionGenerator)
}

func main() {

	setUp()

	fmt.Println("Questioncicas is ready...")

	http.HandleFunc("/beAsked", getQuestion)

	http.HandleFunc("/validate", validateAnswer)

	http.ListenAndServe(":8090", nil)
}
