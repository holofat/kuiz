package users

import (
	"context"
	"errors"
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

func (repo *UserRepository) Register(domain users.User, ctx context.Context) (users.User, error) {
	userDb := FromDomain(domain)

	// Check if email is exist
	var countEmail int64
	repo.db.Table("users").Where("email = ?", userDb.Email).Count(&countEmail)
	if countEmail > 0 {
		return users.User{}, errors.New("email is already exist")
	}

	// Make encrypted password
	encryptedPassword, errorEncrypted := helper.HashPassword(userDb.Password)
	if errorEncrypted != nil {
		return users.User{}, errors.New("error when encrypting")
	}
	userDb.Password = encryptedPassword

	// Add a new user to database
	err := repo.db.Create(&userDb).Error

	if err != nil {
		return users.User{}, errors.New(err.Error())
	}

	return userDb.ToDomain(), nil
}

func (repo *UserRepository) Login(domain users.User, ctx context.Context) (users.User, error) {

	// Define data from input and database
	userInput := FromDomain(domain)
	userDb := FromDomain(domain)

	// Check if email is exist in database or not
	err := repo.db.Where("email = ?", userInput.Email).First(&userDb).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userInput.ToDomain(), errors.New("email not found")
		}
		return userInput.ToDomain(), errors.New(err.Error())
	}

	// If email exist, then verify password
	foundUserPassword := userDb.Password
	passwordIsValid := helper.VerifyPassword(userInput.Password, foundUserPassword)

	if passwordIsValid {
		return userDb.ToDomain(), nil
	}
	return userDb.ToDomain(), errors.New("password is incorrect")

}
