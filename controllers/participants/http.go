package controllers

import (
	"kuiz/app/helper"
	"kuiz/business/participants"
	"kuiz/controllers"
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

	answerId, _ := strconv.Atoi(c.Param("id_answer"))
	quizId, _ := strconv.Atoi(c.Param("id_quiz"))
	questionId, _ := strconv.Atoi(c.Param("id_question"))

	extractedUserId, _ := helper.ExtractTokenAuth(c.Request)
	err := controller.usecase.AnswerQuestion(int(extractedUserId), quizId, answerId, questionId, ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
		c.Abort()
	} else {
		controllers.SuccessResponse(c, "success")
	}
}
