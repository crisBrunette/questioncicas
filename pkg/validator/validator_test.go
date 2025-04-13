package validator

import (
	"testing"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
	"github.com/stretchr/testify/assert"
)

func TestValidatorCorrectAnswer(t *testing.T) {

	questionsTest := []question_generator.Question{
		{UUID: 0, Question: "question1", Answer: "a"},
		{UUID: 1, Question: "question2", Answer: "a"},
	}

	question_generator.Questions = questionsTest

	question_generator.UUIDRandomQuestion = questionsTest[0].UUID

	var tests = []struct {
		userAnswer string
		want       bool
	}{
		{"a", true},
		{"b", false},
		{"c", false},
		{"A", true},
	}

	for _, tt := range tests {
		t.Run(tt.userAnswer, func(t *testing.T) {
			result := Validator(tt.userAnswer)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestValidatorEmptyRandomQuestionAnswer(t *testing.T) {

	question_generator.UUIDRandomQuestion = -1

	var tests = []struct {
		userAnswer string
	}{
		{"a"},
		{"b"},
		{"c"},
		{"A"},
	}

	for _, tt := range tests {
		t.Run(tt.userAnswer, func(t *testing.T) {
			result := Validator(tt.userAnswer)
			assert.False(t, result)
		})
	}
}
