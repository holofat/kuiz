package users

import (
	"context"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *UserUseCase) Register(domain Domain, ctx context.Context) (Domain, error) {
	user, err := usecase.repo.Register(domain, ctx)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (usecase *UserUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {

	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return domain, err
	}

	return user, nil
}
