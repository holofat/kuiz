package questions

import (
	"context"
	"time"
)

type QuestionUseCase struct {
	repo QuestionRepoInterface
	ctx  time.Duration
}

func NewUsecase(QuestionRepo QuestionRepoInterface, contexTimeout time.Duration) QuestionUsecaseInterface {
	return &QuestionUseCase{
		repo: QuestionRepo,
		ctx:  contexTimeout,
	}
}

func (usecase *QuestionUseCase) CreateQuestion(domain Question, ctx context.Context) (Question, error) {
	question, err := usecase.repo.CreateQuestion(domain, ctx)

	if err != nil {
		return Question{}, err
	}

	return question, nil
}

func (usecase *QuestionUseCase) GetQuestion(id string, ctx context.Context) ([]Question, error) {
	questionList, err := usecase.repo.GetQuestion(id, ctx)

	if err != nil {
		return []Question{}, err
	}

	return questionList, nil
}

func (usecase *QuestionUseCase) DeleteQuestion(id string, ctx context.Context) error {
	err := usecase.repo.DeleteQuestion(id, ctx)

	if err != nil {
		return err
	}

	return nil
}
