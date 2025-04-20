package question_generator

import (
	"math/rand"
	"net/http"
)

type QuestionGenerator struct {
	Questions          QuestionList
	UUIDRandomQuestion int
}

func InitQuestionGenerator(questions *QuestionList) *QuestionGenerator {
	return &QuestionGenerator{
		Questions: *questions,
	}
}

func (q *QuestionGenerator) Generator(req *http.Request) string {
	index := rand.Intn(len(q.Questions.QuestionList))
	q.UUIDRandomQuestion = q.Questions.QuestionList[index].UUID
	return q.Questions.QuestionList[index].Question
}
