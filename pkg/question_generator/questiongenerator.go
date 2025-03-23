package question_generator

import (
	"math/rand"
	"net/http"

	"github.com/crisBrunette/questioncicas/internal/database"
)

var RamdonQuestion string

func Generator(req *http.Request) string {
	keys := make([]string, 0, len(database.QuestionAnswer)) // len = 0, cap = 5(database.QuestionAnswer)
	for k := range database.QuestionAnswer {
		keys = append(keys, k)
	}
	RamdonQuestion = keys[rand.Intn(len(keys))]
	return RamdonQuestion
}
