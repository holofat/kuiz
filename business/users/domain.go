package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
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
	Login(domain User, ctx context.Context) (User, error)
	Register(domain User, ctx context.Context) (User, error)
}

type UserRepoInterface interface {
	Login(domain User, ctx context.Context) (User, error)
	Register(domain User, ctx context.Context) (User, error)
}
