package validator

import (
	"strings"

	"github.com/crisBrunette/questioncicas/internal/database"
	"github.com/crisBrunette/questioncicas/pkg/question_generator"
)

func Validator(userAnswer string) bool {
	correctAnswer, _ := database.QuestionAnswer[question_generator.RandomQuestion]
	return strings.EqualFold(userAnswer, correctAnswer)
}
