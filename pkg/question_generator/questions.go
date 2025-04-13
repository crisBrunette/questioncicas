package question_generator

import (
	"fmt"

	"github.com/crisBrunette/questioncicas/internal/database"
)

type Question struct {
	UUID     int
	Question string
	Answer   string
}

var Questions []Question

func GenerateQuestions() {
	Questions = make([]Question, 0, len(database.QuestionAnswer))

	uuid := 0
	for question, answer := range database.QuestionAnswer {
		q := Question{
			UUID:     uuid,
			Question: question,
			Answer:   answer,
		}
		Questions = append(Questions, q)
		uuid++
	}

	for _, q := range Questions {
		fmt.Printf("UUID: %d, Question: %s, Answer: %s", q.UUID, q.Question, q.Answer) // Cambiado Printfln a Printf
	}
}
