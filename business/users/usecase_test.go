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
var userUsecaseInterface users.UserUsecaseInterface
var userLoginDataDummy users.User

func setup() {
	userUsecaseInterface = users.NewUsecase(&userRepoInterfaceMock, time.Hour*1)
	userLoginDataDummy = users.User{
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
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(userLoginDataDummy, nil).Once()
		var requestLoginDomain = users.User{
			Email:    "john@doe.com",
			Password: "123",
		}
		domain, err := userUsecaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userLoginDataDummy, domain)
	})

	t.Run("Login with Email Empty", func(t *testing.T) {
		var requestLogin = users.User{
			Email:    "",
			Password: "123",
		}
		domain, err := userUsecaseInterface.Login(requestLogin, context.Background())

		assert.Equal(t, "email is empty", err.Error())
		assert.Equal(t, requestLogin, domain)
	})

	t.Run("Login with Password Empty", func(t *testing.T) {
		var requestLogin = users.User{
			Email:    "john@doe.com",
			Password: "",
		}
		domain, err := userUsecaseInterface.Login(requestLogin, context.Background())

		assert.Equal(t, "password is empty", err.Error())
		assert.Equal(t, requestLogin, domain)
	})

	t.Run("Login with Password and Email Empty", func(t *testing.T) {
		var requestLoginDomain = users.User{
			Email:    "",
			Password: "",
		}
		domain, err := userUsecaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "email and password must be filled", err.Error())
		assert.Equal(t, requestLoginDomain, domain)
	})
}
