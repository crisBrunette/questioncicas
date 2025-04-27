package question_generator

import (
	"math/rand"
	"net/http"
)

type QuestionGenerator struct {
	Questions        QuestionList
	IDRandomQuestion int
}

func InitQuestionGenerator(questions *QuestionList) *QuestionGenerator {
	return &QuestionGenerator{
		Questions: *questions,
	}
}

func (q *QuestionGenerator) Generator(req *http.Request) string {
	index := rand.Intn(len(q.Questions.QuestionList))
	q.IDRandomQuestion = q.Questions.QuestionList[index].ID
	return q.Questions.QuestionList[index].Question
}
