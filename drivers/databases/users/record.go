package users

import (
	"kuiz/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FullName  string
	Email     string
	Password  string
	Token     string
}

func (user User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Token:     user.Token,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		FullName:  domain.FullName,
		Email:     domain.Email,
		Password:  domain.Password,
		Token:     domain.Token,
	}
}
