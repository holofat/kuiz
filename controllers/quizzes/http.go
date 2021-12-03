package controllers

import (
	"kuiz/app/helper"
	"kuiz/business/quizzes"
	"kuiz/controllers"
	"kuiz/controllers/quizzes/request"
	"kuiz/controllers/quizzes/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuizController struct {
	usecase quizzes.QuizUsecaseInterface
}

func NewQuizController(uc quizzes.QuizUsecaseInterface) *QuizController {
	return &QuizController{
		usecase: uc,
	}
}

func (controller *QuizController) GetQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id_quiz")

	quiz, err := controller.usecase.GetQuiz(id, ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	} else {
		controllers.SuccessResponse(c, response.FromDomain(quiz))
	}
}

func (controller *QuizController) CreateQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	var createQuiz request.CreateQuiz

	extractedUserId, err := helper.ExtractTokenAuth(c.Request)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", err)
	} else {
		if err := c.ShouldBindJSON(&createQuiz); err != nil {
			controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
		} else {
			createQuiz.AuthorID = extractedUserId
			quiz, err := controller.usecase.CreateQuiz(*createQuiz.ToDomain(), ctx)
			if err != nil {
				controllers.ErrorResponse(c, http.StatusInternalServerError, "error in body", err)
			} else {
				controllers.SuccessResponse(c, response.FromDomain(quiz))
			}
		}
	}
}

func (controller *QuizController) DeleteQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	var quiz request.DeleteQuiz
	id := c.Param("id_quiz")
	extractedUserId, errAuth := helper.ExtractTokenAuth(c.Request)
	if errAuth != nil {
		controllers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", errAuth)
		c.Abort()
	}

	_, err := controller.usecase.DeleteQuiz(id, extractedUserId, *quiz.ToDomain(), ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
		c.Abort()
	}
	controllers.SuccessResponse(c, "Success Deleted from Database")
}
