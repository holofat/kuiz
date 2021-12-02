package main

import (
	"kuiz/app/routes"

	userUseCase "kuiz/business/users"
	userController "kuiz/controllers/users"
	userRepo "kuiz/drivers/databases/users"

	quizUseCase "kuiz/business/quizzes"
	quizController "kuiz/controllers/quizzes"
	quizRepo "kuiz/drivers/databases/quizzes"

	questionUseCase "kuiz/business/questions"
	questionController "kuiz/controllers/questions"
	questionRepo "kuiz/drivers/databases/questions"

	answerUseCase "kuiz/business/answers"
	answerController "kuiz/controllers/answers"
	answerRepo "kuiz/drivers/databases/answers"

	participantUseCase "kuiz/business/participants"
	participantController "kuiz/controllers/participants"
	participantRepo "kuiz/drivers/databases/participants"

	participantScoreUseCase "kuiz/business/participant_score"
	participantScoreController "kuiz/controllers/participant_score"
	participantScoreRepo "kuiz/drivers/databases/participant_score"

	"kuiz/drivers/mysql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{}, &quizRepo.Quiz{}, &questionRepo.Question{}, &answerRepo.Answer{}, &participantRepo.Participant{}, &participantScoreRepo.ParticipantScore{})
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitialDb()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	r := gin.Default()

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUseCase.NewUsecase(userRepoInterface, timeoutContext)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	quizRepoInterface := quizRepo.NewQuizRepository(db)
	quizUseCaseInterface := quizUseCase.NewUsecase(quizRepoInterface, timeoutContext)
	quizControllerInterface := quizController.NewQuizController(quizUseCaseInterface)

	questionRepoInterface := questionRepo.NewQuestionRepository(db)
	questionUseCaseInterface := questionUseCase.NewUsecase(questionRepoInterface, timeoutContext)
	questionControllerInterface := questionController.NewQuestionController(questionUseCaseInterface)

	answerRepoInterface := answerRepo.NewAnswerRepository(db)
	answerUseCaseInterface := answerUseCase.NewUsecase(answerRepoInterface, timeoutContext)
	answerControllerInterface := answerController.NewAnswerController(answerUseCaseInterface)

	participantRepoInterface := participantRepo.NewParticipantRepository(db)
	participantUseUseCaseInterface := participantUseCase.NewUsecase(participantRepoInterface, timeoutContext)
	participantControllerInterface := participantController.NewParticipantController(participantUseUseCaseInterface)

	participantScoreRepoInterface := participantScoreRepo.NewParticipantRepository(db)
	participantScoreUseCaseInterface := participantScoreUseCase.NewUseCase(participantScoreRepoInterface, timeoutContext)
	participantScoreControllerInterface := participantScoreController.NewParticipantController(participantScoreUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:             *userControllerInterface,
		QuizController:             *quizControllerInterface,
		QuestionController:         *questionControllerInterface,
		AnswerController:           *answerControllerInterface,
		ParticipantController:      *participantControllerInterface,
		ParticipantScoreController: *participantScoreControllerInterface,
	}

	routesInit.RouteRegister(r)
	log.Fatal(r.Run())

}
