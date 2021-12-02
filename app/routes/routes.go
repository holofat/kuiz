package routes

import (
	answerController "kuiz/controllers/answers"
	participantScoreController "kuiz/controllers/participant_score"
	participantController "kuiz/controllers/participants"
	questionController "kuiz/controllers/questions"
	quizController "kuiz/controllers/quizzes"
	userController "kuiz/controllers/users"

	"kuiz/app/middleware"

	"github.com/gin-gonic/gin"
)

type RouteControllerList struct {
	UserController             userController.UserController
	QuestionController         questionController.QuestionController
	QuizController             quizController.QuizController
	AnswerController           answerController.AnswerController
	ParticipantController      participantController.ParticipantController
	ParticipantScoreController participantScoreController.ParticipantScoreController
}

func (controller RouteControllerList) RouteRegister(r *gin.Engine) {
	authRoute := r.Group("")
	authRoute.Use(middleware.TokenAuthMiddleware())

	// User Controller
	r.POST("/auth/register", controller.UserController.Register)
	r.POST("/auth/login", controller.UserController.Login)
	authRoute.GET("/quiz/:id_quiz/participantScore", controller.ParticipantScoreController.FetchAllData)
	authRoute.GET("/quiz/:id_quiz/score/:id_participant", controller.ParticipantScoreController.FetchDataByIdParticipant)

	// Quiz Controller
	r.GET("/quiz/:id_quiz", controller.QuizController.GetQuiz)
	authRoute.POST("/quiz/create", controller.QuizController.CreateQuiz)
	authRoute.DELETE("/quiz/delete/:id", controller.QuizController.DeleteQuiz)
	authRoute.POST("/quiz/:id_quiz/question/:id_question/answer/:id_answer", controller.ParticipantController.AnswerQuestion)

	// Question Controller
	authRoute.POST("/question/create", controller.QuestionController.CreateQuestion)
	authRoute.GET("/quiz/:id_quiz/questions", controller.QuestionController.GetQuestion)
	authRoute.DELETE("/question/:id_question/delete", controller.QuestionController.DeleteQuestion)

	// Answer Controller
	authRoute.POST("/answer/create", controller.AnswerController.CreateAnswer)
	authRoute.DELETE("/answer/:id_answer/delete", controller.AnswerController.DeleteAnswer)
}
