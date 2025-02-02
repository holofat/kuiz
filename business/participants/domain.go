package participants

import (
	"context"
)

type Participant struct {
	Id         int
	UserId     int
	QuizId     int
	QuestionId int
	AnswerId   int
}

type ParticipantUsecaseInterface interface {
	AnswerQuestion(idUser int, idQuiz int, idAnswer int, idQuestion int, ctx context.Context) error
}

type ParticipantRepoInterface interface {
	AnswerQuestion(idUser int, idQuiz int, idAnswer int, idQuestion int, ctx context.Context) error
}
