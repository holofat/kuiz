package request

import (
	"kuiz/business/answers"
)

type CreateAnswer struct {
	QuestionId    uint
	Answer        string
	CorrectAnswer bool
}

func (answer *CreateAnswer) ToDomain() *answers.Domain {
	return &answers.Domain{
		QuestionId:    answer.QuestionId,
		Answer:        answer.Answer,
		CorrectAnswer: answer.CorrectAnswer,
	}
}
