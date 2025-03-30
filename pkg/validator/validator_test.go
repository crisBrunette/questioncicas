package validator

import (
	"testing"

	"github.com/crisBrunette/questioncicas/pkg/question_generator"
	"github.com/stretchr/testify/assert"
)

func TestValidatorCorrectAnswer(t *testing.T) {

	question_generator.RandomQuestion = "¿Como me llamo? a)Cristina, b)Begoña, c)Imposible de distinguir"

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

	question_generator.RandomQuestion = ""

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
