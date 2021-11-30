package controllers

import (
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
		c.Abort()
	}

	user, err := controller.usecase.Login(*userLogin.ToDomain(), ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusUnauthorized, err.Error(), err)
		c.Abort()
	}

	var authD helper.AuthDetails
	authD.UserId = user.Id

	token, loginErr := helper.GenerateToken(authD)
	if loginErr != nil {
		controllers.ErrorResponse(c, http.StatusForbidden, loginErr.Error(), loginErr)
		c.Abort()
	}

	controllers.SuccessResponse(c, token)

}
