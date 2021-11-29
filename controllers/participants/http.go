package controllers

import (
	"kuiz/business/participants"
	"kuiz/controllers"
	"kuiz/controllers/participants/request"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ParticipantController struct {
	usecase participants.ParticipantUsecaseInterface
}

func NewParticipantController(usecase participants.ParticipantUsecaseInterface) *ParticipantController {
	return &ParticipantController{
		usecase: usecase,
	}
}

func (controller *ParticipantController) AnswerQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	var newAnswered request.AnswerQuestion
	quizId, _ := strconv.Atoi(c.Param("id_quiz"))
	questionId, _ := strconv.Atoi(c.Param("id_question"))
	newAnswered.QuestionId = questionId
	newAnswered.QuizId = quizId

	if err := c.ShouldBindJSON(&newAnswered); err != nil {
		controllers.ErrorResponse(c, http.StatusBadRequest, "error binding", err)
	} else {
		err := controller.usecase.AnswerQuestion(*newAnswered.ToDomain(), ctx)
		if err != nil {
			controllers.ErrorResponse(c, http.StatusInternalServerError, "error in body", err)
		} else {
			controllers.SuccessResponse(c, nil)
		}
	}
}
