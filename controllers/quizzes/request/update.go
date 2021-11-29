package request

import (
	"kuiz/business/quizzes"
)

type UpdateQuiz struct {
	GivenTime uint   `json:"given_time"`
	TitleQuiz string `json:"title"`
}

func (quiz *UpdateQuiz) ToDomain() *quizzes.Domain {
	return &quizzes.Domain{
		GivenTime: quiz.GivenTime,
		TitleQuiz: quiz.TitleQuiz,
	}
}
