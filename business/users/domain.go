package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FullName  string
	Username  string
	Password  string
	Email     string
	Token     string
}

type UserUsecaseInterface interface {
	Login(domain Domain, ctx context.Context) (string, error)
	Register(domain Domain, ctx context.Context) (Domain, error)
}

type UserRepoInterface interface {
	Login(domain Domain, ctx context.Context) (string, error)
	Register(domain Domain, ctx context.Context) (Domain, error)
}
