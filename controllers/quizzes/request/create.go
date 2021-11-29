package request

import (
	"kuiz/business/quizzes"
)

type CreateQuiz struct {
	AuthorID  uint   `json:"author_id"`
	TitleQuiz string `json:"title"`
	GivenTime uint   `json:"given_time"`
}

func (quiz *CreateQuiz) ToDomain() *quizzes.Domain {
	return &quizzes.Domain{
		AuthorId:  quiz.AuthorID,
		TitleQuiz: quiz.TitleQuiz,
		GivenTime: quiz.GivenTime,
	}
}
