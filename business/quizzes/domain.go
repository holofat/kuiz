package quizzes

import (
	"context"
	"kuiz/business/questions"
	"time"
)

type Domain struct {
	Id        uint
	AuthorId  uint
	TitleQuiz string
	CreatedAt time.Time
	GivenTime uint
	Status    bool
	Question  []questions.Question
}

type QuizUsecaseInterface interface {
	CreateQuiz(domain Domain, ctx context.Context) (Domain, error)
	DeleteQuiz(id string, domain Domain, ctx context.Context) (Domain, error)
}

type QuizRepoInterface interface {
	CreateQuiz(domain Domain, ctx context.Context) (Domain, error)
	DeleteQuiz(id string, domain Domain, ctx context.Context) (Domain, error)
}
