package participants

import (
	"context"
)

type Participant struct {
	Id         int
	UserId     int
	QuizId     int
	QuestionId int
	AnswerId   int
}

type ParticipantUsecaseInterface interface {
	AnswerQuestion(domain Participant, ctx context.Context) error
}

type ParticipantRepoInterface interface {
	AnswerQuestion(domain Participant, ctx context.Context) error
}
