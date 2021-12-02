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

func (usecase *ParticipantUseCase) AnswerQuestion(idUser int, idQuiz int, idAnswer int, idQuestion int, ctx context.Context) error {
	err := usecase.repo.AnswerQuestion(idUser, idQuiz, idAnswer, idQuestion, ctx)
	if err != nil {
		return err
	}

	return nil
}
