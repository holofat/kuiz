package users

import (
	"context"
	"errors"
	"fmt"
	"kuiz/app/helper"
	"kuiz/business/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) Register(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(domain)
	encryptedPassword, errorEncrypted := helper.HashPassword(userDb.Password)
	if errorEncrypted != nil {
		return users.Domain{}, errors.New("error when encrypting")
	}
	userDb.Password = encryptedPassword

	err := repo.db.Create(&userDb).Error

	if err != nil {
		return users.Domain{}, errors.New("error in database")
	}

	return userDb.ToDomain(), nil
}

func (repo *UserRepository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userInput := FromDomain(domain)
	userDb := FromDomain(domain)

	err := repo.db.Where("email = ?", userInput.Email).First(&userDb).Error
	fmt.Println(userInput.Email, userDb)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userInput.ToDomain(), errors.New("email not found")
		}
		return userInput.ToDomain(), errors.New("error in database")
	}
	foundUserPassword := userDb.Password
	passwordIsValid := helper.VerifyPassword(userInput.Password, foundUserPassword)

	if passwordIsValid {
		return userDb.ToDomain(), nil
	}
	return userDb.ToDomain(), errors.New("password is incorrect")

}
