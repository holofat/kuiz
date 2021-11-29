package controllers

import (
	"fmt"
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

func (controller *QuizController) CreateQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	var createQuiz request.CreateQuiz

	if err := c.ShouldBindJSON(&createQuiz); err != nil {
		controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	} else {
		quiz, err := controller.usecase.CreateQuiz(*createQuiz.ToDomain(), ctx)
		fmt.Println(quiz)
		if err != nil {
			controllers.ErrorResponse(c, http.StatusInternalServerError, "error in body", err)
		}
		controllers.SuccessResponse(c, response.FromDomain(quiz))
	}
}

func (controller *QuizController) DeleteQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	var quiz request.DeleteQuiz
	id := c.Param("id")

	_, err := controller.usecase.DeleteQuiz(id, *quiz.ToDomain(), ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	} else {
		controllers.SuccessResponse(c, "Success Deleted from Database")
	}
}
