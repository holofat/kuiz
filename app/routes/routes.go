package routes

import (
	answerController "kuiz/controllers/answers"
	participantController "kuiz/controllers/participants"
	questionController "kuiz/controllers/questions"
	quizController "kuiz/controllers/quizzes"
	userController "kuiz/controllers/users"

	"kuiz/app/middleware"

	"github.com/gin-gonic/gin"
)

type RouteControllerList struct {
	UserController        userController.UserController
	QuestionController    questionController.QuestionController
	QuizController        quizController.QuizController
	AnswerController      answerController.AnswerController
	ParticipantController participantController.ParticipantController
}

func (controller RouteControllerList) RouteRegister(r *gin.Engine) {
	authRoute := r.Group("")
	authRoute.Use(middleware.TokenAuthMiddleware())

	// User Controller
	r.POST("/auth/register", controller.UserController.Register)
	r.POST("/auth/login", controller.UserController.Login)

	// Quiz Controller
	authRoute.POST("/quiz/create", controller.QuizController.CreateQuiz)
	authRoute.DELETE("/quiz/delete/:id", controller.QuizController.DeleteQuiz)
	authRoute.POST("/quiz/:id_quiz/question/:id_question", controller.ParticipantController.AnswerQuestion)

	// Question Controller
	authRoute.POST("/question/create", controller.QuestionController.CreateQuestion)
	authRoute.GET("/quiz/:id/questions", controller.QuestionController.GetQuestion)
	authRoute.DELETE("/question/:id/delete", controller.QuestionController.DeleteQuestion)

	// Answer Controller
	authRoute.POST("/answer/create", controller.AnswerController.CreateAnswer)
	authRoute.DELETE("/answer/:id/delete", controller.AnswerController.DeleteAnswer)
}
