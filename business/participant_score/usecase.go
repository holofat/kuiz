package participantscore

import (
	"context"
	"time"
)

type ParticipantScoreUsecase struct {
	repo ParticipantScoreRepoInterface
	ctx  time.Duration
}

func NewUseCase(PartipantRepo ParticipantScoreRepoInterface, contextTimeout time.Duration) ParticipantScoreUseCaseInterface {
	return &ParticipantScoreUsecase{
		repo: PartipantRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *ParticipantScoreUsecase) FetchData(idQuiz int, idAuthor int, ctx context.Context) ([]ParticipantScore, error) {
	participantScoreList, err := usecase.repo.FetchData(idQuiz, idAuthor, ctx)
	if err != nil {
		return []ParticipantScore{}, err
	}
	return participantScoreList, nil
}
