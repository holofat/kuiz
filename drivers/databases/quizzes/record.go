package quizzes

import (
	"kuiz/business/quizzes"
	"time"
)

type Quiz struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	TitleQuiz string
	GivenTime uint
	AuthorId  uint
}

func (quiz Quiz) ToDomain() quizzes.Domain {
	return quizzes.Domain{
		Id:        quiz.Id,
		CreatedAt: quiz.CreatedAt,
		TitleQuiz: quiz.TitleQuiz,
		GivenTime: quiz.GivenTime,
		AuthorId:  quiz.AuthorId,
	}
}

func FromDomain(domain quizzes.Domain) Quiz {
	return Quiz{
		Id:        domain.Id,
		TitleQuiz: domain.TitleQuiz,
		GivenTime: domain.GivenTime,
		AuthorId:  domain.AuthorId,
	}
}