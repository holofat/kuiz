package answers

import (
	"context"
	"time"
)

type Domain struct {
	Id            uint
	QuestionId    uint
	CreatedAt     time.Time
	Answer        string
	CorrectAnswer bool
}

type AnswerUsecaseInterface interface {
	CreateAnswer(domain Domain, ctx context.Context) (Domain, error)
	DeleteAnswer(id string, domain Domain, ctx context.Context) (Domain, error)
}

type AnswerRepoInterface interface {
	CreateAnswer(domain Domain, ctx context.Context) (Domain, error)
	DeleteAnswer(id string, domain Domain, ctx context.Context) (Domain, error)
}
