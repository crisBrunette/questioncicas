package validator

import (
	"strings"

	"slices"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
)

type Validator struct {
	QuestionGenerator question_generator.QuestionGenerator
}

func InitValidator(questionGenerator *question_generator.QuestionGenerator) *Validator {
	return &Validator{
		QuestionGenerator: *questionGenerator,
	}
}

func (v Validator) IsValid(userAnswer string) bool {
	index := slices.IndexFunc(v.QuestionGenerator.Questions.QuestionList, func(n question_generator.Question) bool {
		return n.UUID == v.QuestionGenerator.UUIDRandomQuestion
	})

	if index == -1 {
		return false
	}
	println("Correct answer:", v.QuestionGenerator.Questions.QuestionList[index].Answer)
	return strings.EqualFold(userAnswer, v.QuestionGenerator.Questions.QuestionList[index].Answer)
}
