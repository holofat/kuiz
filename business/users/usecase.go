package users

import (
	"context"
	"errors"
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

	if domain.Email == "" && domain.Password == "" {
		return domain, errors.New("email and password must be filled")
	} else if domain.Email == "" {
		return domain, errors.New("email is empty")
	} else if domain.Password == "" {
		return domain, errors.New("password is empty")
	}

	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return domain, err
	}

	return user, nil
}
