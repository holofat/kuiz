package quizzes

import (
	"context"
	"time"
)

type QuizUseCase struct {
	repo QuizRepoInterface
	ctx  time.Duration
}

func NewUsecase(quizRepo QuizRepoInterface, contextTimeout time.Duration) QuizUsecaseInterface {
	return &QuizUseCase{
		repo: quizRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *QuizUseCase) CreateQuiz(domain Domain, ctx context.Context) (Domain, error) {
	quiz, err := usecase.repo.CreateQuiz(domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return quiz, nil
}

func (usecase *QuizUseCase) DeleteQuiz(id string, domain Domain, ctx context.Context) (Domain, error) {
	quiz, err := usecase.repo.DeleteQuiz(id, domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return quiz, nil
}

func (usecase *QuizUseCase) UpdateQuiz(id string, domain Domain, ctx context.Context) (Domain, error) {
	quiz, err := usecase.UpdateQuiz(id, domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return quiz, nil
}
