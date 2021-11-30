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

func (usecase *QuizUseCase) GetQuiz(id string, ctx context.Context) (Domain, error) {
	quiz, err := usecase.repo.GetQuiz(id, ctx)

	if err != nil {
		return Domain{}, err
	}

	return quiz, nil
}

func (usecase *QuizUseCase) CreateQuiz(domain Domain, ctx context.Context) (Domain, error) {
	quiz, err := usecase.repo.CreateQuiz(domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return quiz, nil
}

func (usecase *QuizUseCase) DeleteQuiz(id string, currentUserId uint, domain Domain, ctx context.Context) (Domain, error) {
	quiz, err := usecase.repo.DeleteQuiz(id, currentUserId, domain, ctx)

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
