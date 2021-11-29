package participants

import (
	"context"
	"time"
)

type ParticipantUseCase struct {
	repo ParticipantRepoInterface
	ctx  time.Duration
}

func NewUsecase(ParticipantRepo ParticipantRepoInterface, contextTimeout time.Duration) ParticipantUsecaseInterface {
	return &ParticipantUseCase{
		repo: ParticipantRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *ParticipantUseCase) AnswerQuestion(domain Participant, ctx context.Context) error {
	err := usecase.repo.AnswerQuestion(domain, ctx)
	if err != nil {
		return err
	}

	return nil
}
