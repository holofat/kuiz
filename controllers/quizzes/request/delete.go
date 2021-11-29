package request

import (
	"kuiz/business/quizzes"
)

type DeleteQuiz struct {
	Id uint `json:"id"`
}

func (quiz *DeleteQuiz) ToDomain() *quizzes.Domain {
	return &quizzes.Domain{
		Id: quiz.Id,
	}
}
