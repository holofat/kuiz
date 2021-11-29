package controllers

import (
	"fmt"
	"kuiz/app/helper"
	"kuiz/business/users"
	"kuiz/controllers"
	"kuiz/controllers/users/request"
	"kuiz/controllers/users/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase users.UserUsecaseInterface
}

func NewUserController(uc users.UserUsecaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var userRegister request.UserRegister
	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	} else {
		user, err := controller.usecase.Register(*userRegister.ToDomain(), ctx)

		if err != nil {
			controllers.ErrorResponse(c, http.StatusInternalServerError, "error in body", err)
		}
		controllers.SuccessResponse(c, response.FromDomain(user))
	}
}

func (controller *UserController) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var userLogin request.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	} else {
		_, err := controller.usecase.Login(*userLogin.ToDomain(), ctx)
		fmt.Println(userLogin)
		if err != nil {
			controllers.ErrorResponse(c, http.StatusUnauthorized, err.Error(), err)
		} else {
			var authD helper.AuthDetails
			authD.Email = userLogin.Email

			token, loginErr := helper.GenerateToken(authD)
			if loginErr != nil {
				controllers.ErrorResponse(c, http.StatusForbidden, loginErr.Error(), loginErr)
			} else {
				controllers.SuccessResponse(c, token)
			}
		}
	}

}
