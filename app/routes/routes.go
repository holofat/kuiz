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
	r.Use(middleware.TokenAuthMiddleware())
	// User Controller
	r.POST("/auth/register", controller.UserController.Register)
	r.POST("/auth/login", controller.UserController.Login)

	// Quiz Controller
	r.POST("/quiz/create", controller.QuizController.CreateQuiz)
	r.DELETE("/quiz/delete/:id", controller.QuizController.DeleteQuiz)
	r.POST("/quiz/:id_quiz/question/:id_question", controller.ParticipantController.AnswerQuestion)

	// Question Controller
	r.POST("/question/create", controller.QuestionController.CreateQuestion)
	r.GET("/quiz/:id/questions", controller.QuestionController.GetQuestion)
	r.DELETE("/question/:id/delete", controller.QuestionController.DeleteQuestion)

	// Answer Controller
	r.POST("/answer/create", controller.AnswerController.CreateAnswer)
	r.DELETE("/answer/:id/delete", controller.AnswerController.DeleteAnswer)
}
