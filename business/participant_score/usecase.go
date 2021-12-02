package participantscore

import (
	"context"
	"time"
)

type ParticipantScoreUsecase struct {
	repo ParticipantScoreRepoInterface
	ctx  time.Duration
}

func NewUseCase(ParticipantRepo ParticipantScoreRepoInterface, contextTimeout time.Duration) ParticipantScoreUseCaseInterface {
	return &ParticipantScoreUsecase{
		repo: ParticipantRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *ParticipantScoreUsecase) FetchAllData(idQuiz int, idAuthor int, ctx context.Context) ([]ParticipantScore, error) {
	participantScoreList, err := usecase.repo.FetchAllData(idQuiz, idAuthor, ctx)
	if err != nil {
		return []ParticipantScore{}, err
	}
	return participantScoreList, nil
}

func (usecase *ParticipantScoreUsecase) FetchDataByIdParticipant(idQuiz int, idParticipant int, ctx context.Context) (ParticipantScore, error) {
	participantScore, err := usecase.repo.FetchDataByIdParticipant(idQuiz, idParticipant, ctx)

	if err != nil {
		return ParticipantScore{}, err
	}

	return participantScore, nil
}
