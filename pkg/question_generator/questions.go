package question_generator

import (
	"github.com/crisBrunette/questioncicas/internal/database"
)

type Question struct {
	ID       int
	Question string
	Answer   string
}

type QuestionList struct {
	QuestionList []Question
}

func IntiQuestionList() QuestionList {
	return QuestionList{
		QuestionList: make([]Question, 0, len(database.QuestionAnswer)),
	}
}

func (questionList *QuestionList) PopulateQuestions() {
	id := 0
	for question, answer := range database.QuestionAnswer {
		q := Question{
			ID:       id,
			Question: question,
			Answer:   answer,
		}
		questionList.QuestionList = append(questionList.QuestionList, q)
		id++
	}
}
