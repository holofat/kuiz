package request

import (
	"kuiz/business/questions"
)

type CreateQuestion struct {
	QuizId           uint   `json:"quiz_id"`
	QuestionSentence string `json:"question"`
}

func (question *CreateQuestion) ToDomain() *questions.Question {
	return &questions.Question{
		QuizId:           question.QuizId,
		QuestionSentence: question.QuestionSentence,
	}
}
