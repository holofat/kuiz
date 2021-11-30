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
	GetQuiz(id string, ctx context.Context) (Domain, error)
	CreateQuiz(domain Domain, ctx context.Context) (Domain, error)
	DeleteQuiz(id string, currentUserId uint, domain Domain, ctx context.Context) (Domain, error)
}

type QuizRepoInterface interface {
	GetQuiz(id string, ctx context.Context) (Domain, error)
	CreateQuiz(domain Domain, ctx context.Context) (Domain, error)
	DeleteQuiz(id string, currentUserId uint, domain Domain, ctx context.Context) (Domain, error)
}
