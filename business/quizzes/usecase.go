package quizzes

import (
	"context"
	"errors"
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
	if domain.GivenTime == 0 && domain.TitleQuiz == "" {
		return Domain{}, errors.New("you must be fill the requirement column")
	} else if domain.TitleQuiz == "" {
		return Domain{}, errors.New("title quiz is empty")
	} else if domain.GivenTime == 0 {
		return Domain{}, errors.New("given time is empty")
	}
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
