package questions

import (
	"context"
	"kuiz/business/answers"
	"time"
)

type Question struct {
	Id               uint
	CreatedAt        time.Time
	QuizId           uint
	QuestionSentence string
	Answer           []answers.Answer
}

type QuestionUsecaseInterface interface {
	CreateQuestion(domain Question, ctx context.Context) (Question, error)
	GetQuestion(id string, ctx context.Context) ([]Question, error)
	DeleteQuestion(id string, ctx context.Context) error
}

type QuestionRepoInterface interface {
	CreateQuestion(domain Question, ctx context.Context) (Question, error)
	GetQuestion(id string, ctx context.Context) ([]Question, error)
	DeleteQuestion(id string, ctx context.Context) error
}
