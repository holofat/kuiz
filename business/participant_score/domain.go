package participantscore

import (
	"context"
)

type ParticipantScore struct {
	Id                      int
	UserId                  int
	QuizId                  int
	Score                   float64
	NumberOfCorrectedAnswer int
}

type ParticipantScoreUseCaseInterface interface {
	FetchAllData(idQuiz int, idAuthor int, ctx context.Context) ([]ParticipantScore, error)
	FetchDataByIdParticipant(idQuiz int, idParticipant int, ctx context.Context) (ParticipantScore, error)
}

type ParticipantScoreRepoInterface interface {
	FetchAllData(idQuiz int, idAuthor int, ctx context.Context) ([]ParticipantScore, error)
	FetchDataByIdParticipant(idQuiz int, idParticipant int, ctx context.Context) (ParticipantScore, error)
}
