package request

import (
	"kuiz/business/users"
)

type UserRegister struct {
	FullName string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (user *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		FullName: user.FullName,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}
