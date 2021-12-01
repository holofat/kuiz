package participants

import (
	"context"
	"errors"
	"kuiz/business/participants"

	"gorm.io/gorm"
)

type ParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(gormDb *gorm.DB) participants.ParticipantRepoInterface {
	return &ParticipantRepository{
		db: gormDb,
	}
}

func (repo *ParticipantRepository) AnswerQuestion(idUser int, idQuiz int, idAnswer int, idQuestion int, ctx context.Context) error {
	answered := FromDomain(participants.Participant{
		AnswerId:   idAnswer,
		QuizId:     idQuiz,
		QuestionId: idQuestion,
		UserId:     idUser,
	})
	var quizCount, questionCount, answerCount int64

	repo.db.Table("quizzes").Where("id = ? ", idQuiz).Count(&quizCount)
	if quizCount < 1 {
		return errors.New("quiz id invalid")
	}

	repo.db.Table("questions").Where("id = ?", idQuestion).Count(&questionCount)
	if questionCount < 1 {
		return errors.New("question id invalid")
	}

	repo.db.Table("answers").Where("id = ?", idAnswer).Count(&answerCount)
	if answerCount < 1 {
		return errors.New("answer id invalid")
	}

	err := repo.db.Create(&answered).Error
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
