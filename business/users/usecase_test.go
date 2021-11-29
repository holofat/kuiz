package users_test

import (
	"context"
	"kuiz/business/users"
	"kuiz/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepoInterfaceMock mocks.UserRepoInterface
var UserUsecaseInterface users.UserUsecaseInterface
var userLoginDataDummy users.Domain

func setup() {
	UserUsecaseInterface = users.NewUsecase(&userRepoInterfaceMock, time.Hour*1)
	userLoginDataDummy = users.Domain{
		Id:       1,
		FullName: "john doe",
		Email:    "john@doe.com",
		Password: "123",
		Username: "john",
	}
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Success Login", func(t *testing.T) {
		status := "Success Login"
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.Domain"), mock.Anything).Return(status, nil).Once()

		var requestLoginDomain = users.Domain{
			Email:    "john@doe.com",
			Password: "123",
		}
		domain, err := UserUsecaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, status, domain)
	})

	t.Run("Login with Email Empty", func(t *testing.T) {
		var requestLoginDomain = users.Domain{
			Email:    "",
			Password: "123",
		}
		domain, err := UserUsecaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Email is empty", err.Error())
		assert.Equal(t, "error", domain)
	})

	t.Run("Login with Password Empty", func(t *testing.T) {
		var requestLoginDomain = users.Domain{
			Email:    "john@doe.com",
			Password: "",
		}
		domain, err := UserUsecaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Password is empty", err.Error())
		assert.Equal(t, "error", domain)
	})

	t.Run("Login with Password and Email Empty", func(t *testing.T) {
		var requestLoginDomain = users.Domain{
			Email:    "",
			Password: "",
		}
		domain, err := UserUsecaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Email and Password must be filled", err.Error())
		assert.Equal(t, "error", domain)
	})
}
