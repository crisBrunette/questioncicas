package validator

import (
	"strings"

	"slices"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
)

func Validator(userAnswer string) bool {
	index := slices.IndexFunc(question_generator.Questions, func(n question_generator.Question) bool {
		return n.UUID == question_generator.UUIDRandomQuestion
	})

	if index == -1 {
		return false
	}

	return strings.EqualFold(userAnswer, question_generator.Questions[index].Answer)
}
