package controllers

import (
	"kuiz/business/answers"
	"kuiz/controllers"
	"kuiz/controllers/answers/request"
	"kuiz/controllers/answers/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AnswerController struct {
	uc answers.AnswerUsecaseInterface
}

func NewAnswerController(uc answers.AnswerUsecaseInterface) *AnswerController {
	return &AnswerController{
		uc: uc,
	}
}

func (controller *AnswerController) CreateAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	var newAnswer request.CreateAnswer

	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	} else {
		Answer, err := controller.uc.CreateAnswer(*newAnswer.ToDomain(), ctx)
		if err != nil {
			controllers.ErrorResponse(c, http.StatusInternalServerError, "error in body", err)
		} else {
			controllers.SuccessResponse(c, response.FromDomain(Answer))
		}
	}
}

func (controller *AnswerController) DeleteAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	var Answer request.DeleteAnswer
	id := c.Param("id")

	_, err := controller.uc.DeleteAnswer(id, *Answer.ToDomain(), ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	} else {
		controllers.SuccessResponse(c, "Success deleted from database")
	}
}
