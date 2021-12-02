package request

import (
	"kuiz/business/users"
)

type UserRegister struct {
	FullName string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (user *UserRegister) ToDomain() *users.User {
	return &users.User{
		FullName: user.FullName,
		Password: user.Password,
		Email:    user.Email,
	}
}
