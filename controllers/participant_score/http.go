package controllers

import (
	"kuiz/app/helper"
	participant_score "kuiz/business/participant_score"
	"kuiz/controllers"
	"kuiz/controllers/participant_score/response"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ParticipantScoreController struct {
	usecase participant_score.ParticipantScoreUseCaseInterface
}

func NewParticipantController(usecase participant_score.ParticipantScoreUseCaseInterface) *ParticipantScoreController {
	return &ParticipantScoreController{
		usecase: usecase,
	}
}

func (controller *ParticipantScoreController) FetchAllData(c *gin.Context) {
	ctx := c.Request.Context()

	extractedUserId, _ := helper.ExtractTokenAuth(c.Request)
	idQuiz, _ := strconv.Atoi(c.Param("id_quiz"))

	listScore, err := controller.usecase.FetchAllData(idQuiz, int(extractedUserId), ctx)
	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	} else {
		controllers.SuccessResponse(c, response.FromDomainList(listScore))
	}
}

func (controller *ParticipantScoreController) FetchDataByIdParticipant(c *gin.Context) {
	ctx := c.Request.Context()

	idParticipant, _ := strconv.Atoi(c.Param("id_participant"))
	idQuiz, _ := strconv.Atoi(c.Param("id_quiz"))

	score, err := controller.usecase.FetchDataByIdParticipant(idQuiz, idParticipant, ctx)

	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	} else {
		controllers.SuccessResponse(c, response.FromDomain(score))
	}
}
