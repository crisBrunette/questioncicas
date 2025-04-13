package question_generator

import (
	"math/rand"
	"net/http"
)

var UUIDRandomQuestion int

func Generator(req *http.Request) string {
	index := rand.Intn(len(Questions))
	UUIDRandomQuestion = Questions[index].UUID
	return Questions[index].Question
}
