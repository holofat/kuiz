package request

import (
	"kuiz/business/questions"
)

type DeleteQuestion struct {
	Id uint `json:"id"`
}

func (question *DeleteQuestion) ToDomain() *questions.Question {
	return &questions.Question{
		Id: question.Id,
	}
}
