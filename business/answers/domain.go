package answers

import (
	"context"
	"time"
)

type Answer struct {
	Id            uint
	QuestionId    uint
	CreatedAt     time.Time
	Answer        string
	CorrectAnswer bool
}

type AnswerUsecaseInterface interface {
	CreateAnswer(domain Answer, ctx context.Context) (Answer, error)
	DeleteAnswer(id string, domain Answer, ctx context.Context) (Answer, error)
}

type AnswerRepoInterface interface {
	CreateAnswer(domain Answer, ctx context.Context) (Answer, error)
	DeleteAnswer(id string, domain Answer, ctx context.Context) (Answer, error)
}
