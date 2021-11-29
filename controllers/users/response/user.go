package response

import (
	"kuiz/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.FullName,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}
