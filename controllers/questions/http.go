package controllers

import (
	"kuiz/business/questions"
	"kuiz/controllers"
	"kuiz/controllers/questions/request"
	"kuiz/controllers/questions/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	usecase questions.QuestionUsecaseInterface
}

func NewQuestionController(usecase questions.QuestionUsecaseInterface) *QuestionController {
	return &QuestionController{
		usecase: usecase,
	}
}

func (controller *QuestionController) CreateQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	var newQuestion request.CreateQuestion

	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	} else {
		question, err := controller.usecase.CreateQuestion(*newQuestion.ToDomain(), ctx)
		if err != nil {
			controllers.ErrorResponse(c, http.StatusInternalServerError, "error in body", err)
		} else {
			controllers.SuccessResponse(c, response.FromDomain(question))
		}
	}
}

func (controller *QuestionController) DeleteQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id_question")

	err := controller.usecase.DeleteQuestion(id, ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	} else {
		controllers.SuccessResponse(c, "Success deleted from database")
	}
}

func (controller *QuestionController) GetQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id_quiz")

	questionList, err := controller.usecase.GetQuestion(id, ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	} else {
		controllers.SuccessResponse(c, response.FromDomainList(questionList))
	}
}
