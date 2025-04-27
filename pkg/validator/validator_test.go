package validator

import (
	"testing"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
	"github.com/stretchr/testify/assert"
)

func TestValidatorCorrectAnswer(t *testing.T) {

	questionsTest := []question_generator.Question{
		{ID: 0, Question: "question1", Answer: "a"},
		{ID: 1, Question: "question2", Answer: "a"},
	}

	question_generator.Questions = questionsTest

	question_generator.IDRandomQuestion = questionsTest[0].ID

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

	question_generator.IDRandomQuestion = -1

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
