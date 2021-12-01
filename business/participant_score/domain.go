package participantscore

import (
	"context"
)

type ParticipantScore struct {
	Id       int
	FullName string
	QuizId   int
	Score    int
}

type ParticipantScoreUseCaseInterface interface {
	FetchData(idQuiz int, idAuthor int, ctx context.Context) ([]ParticipantScore, error)
}

type ParticipantScoreRepoInterface interface {
	FetchData(idQuiz int, idAuthor int, ctx context.Context) ([]ParticipantScore, error)
}
