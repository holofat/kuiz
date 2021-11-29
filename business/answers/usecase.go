package answers

import (
	"context"
	"time"
)

type AnswerUseCase struct {
	repo AnswerRepoInterface
	ctx  time.Duration
}

func NewUsecase(answerRepo AnswerRepoInterface, contextTimeout time.Duration) AnswerUsecaseInterface {
	return &AnswerUseCase{
		repo: answerRepo,
		ctx:  contextTimeout,
	}
}

func (uc *AnswerUseCase) CreateAnswer(domain Domain, ctx context.Context) (Domain, error) {
	answer, err := uc.repo.CreateAnswer(domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return answer, nil
}

func (uc *AnswerUseCase) DeleteAnswer(id string, domain Domain, ctx context.Context) (Domain, error) {
	answer, err := uc.repo.DeleteAnswer(id, domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return answer, nil
}
