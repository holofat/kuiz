package response

import (
	"kuiz/business/questions"
	"kuiz/business/quizzes"
	"time"
)

type QuizResponse struct {
	Id        uint                 `json:"id"`
	AuthorId  uint                 `json:"author_id"`
	TitleQuiz string               `json:"title_quiz"`
	CreatedAt time.Time            `json:"created_at"`
	GivenTime uint                 `json:"given_time"`
	ExpiredAt time.Time            `json:"expired_at"`
	Question  []questions.Question `json:"questiont_list"`
}

func FromDomain(domain quizzes.Domain) QuizResponse {
	return QuizResponse{
		Id:        domain.Id,
		AuthorId:  domain.AuthorId,
		TitleQuiz: domain.TitleQuiz,
		CreatedAt: domain.CreatedAt,
		GivenTime: domain.GivenTime,
		Question:  domain.Question,
	}
}
