package question_generator

import (
	"github.com/crisBrunette/questioncicas/internal/database"
)

type Question struct {
	UUID     int
	Question string
	Answer   string
}

type QuestionList struct {
	QuestionList []Question
}

func (q Question) GetUUID() int {
	return q.UUID
}
func (q Question) GetQuestion() string {
	return q.Question
}
func (q Question) GetAnswer() string {
	return q.Answer
}

func IntiQuestionList() QuestionList {
	return QuestionList{
		QuestionList: make([]Question, 0, len(database.QuestionAnswer)),
	}
}

func (questionList *QuestionList) PopulateQuestions() {
	uuid := 0
	for question, answer := range database.QuestionAnswer {
		q := Question{
			UUID:     uuid,
			Question: question,
			Answer:   answer,
		}
		questionList.QuestionList = append(questionList.QuestionList, q)
		uuid++
	}
}
