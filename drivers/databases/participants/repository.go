package participants

import (
	"context"
	"errors"
	"fmt"
	"kuiz/business/participants"
	"kuiz/drivers/databases/answers"
	participantscore "kuiz/drivers/databases/participant_score"

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

	var quizCount, questionCount, answerCount, checkExistingScore, countQuestion int64
	countCorrectedAnswer := 0

	// Check if quiz id is valid
	repo.db.Table("quizzes").Where("id = ? ", idQuiz).Count(&quizCount)
	if quizCount < 1 {
		return errors.New("quiz id invalid")
	}

	// Check if the question from existing quiz
	repo.db.Table("questions").Where("id = ? AND quiz_id = ?", idQuestion, idQuiz).Count(&questionCount)
	if questionCount < 1 {
		return errors.New("question id invalid")
	}

	// Check if the answer from existing question
	repo.db.Table("answers").Where("id = ? AND question_id", idAnswer, idQuestion).Count(&answerCount)
	if answerCount < 1 {
		return errors.New("answer id invalid")
	}

	// Save participant's answer to the database
	if err := repo.db.Create(&answered).Error; err != nil {
		return errors.New(err.Error())
	}

	// Check if a participant's score record is exist
	repo.db.Table("participant_scores").Where("quiz_id = ? AND user_id = ?", idQuiz, idUser).Count(&checkExistingScore)

	// Count the questions
	repo.db.Table("questions").Where("quiz_id = ?", idQuiz).Count(&countQuestion)

	// Check the answer
	var answer answers.Answer
	repo.db.Where("id = ?", idAnswer).First(&answer)
	checkAnswer := answer.CorrectAnswer
	if checkAnswer {
		countCorrectedAnswer++
	}
	// Count the number of score
	tempScore := float64(countCorrectedAnswer) / float64(countQuestion) * 100
	fmt.Println(tempScore)
	newScore := participantscore.ParticipantScore{
		UserId:                  idUser,
		QuizId:                  idQuiz,
		Score:                   tempScore,
		NumberOfCorrectedAnswer: countCorrectedAnswer,
	}

	// If participant's score record is not exist, then create a new record
	if checkExistingScore < 1 {
		if err := repo.db.Create(&newScore).Error; err != nil {
			return errors.New(err.Error())
		}
	} else {
		var updatedScore participantscore.ParticipantScore

		// Save existing record
		repo.db.Where("user_id = ? AND quiz_id", idUser, idQuiz).First(&updatedScore)

		// Add the number of corrected answer
		updatedScore.NumberOfCorrectedAnswer += countCorrectedAnswer

		// Count the final score
		updatedScore.Score = (float64(updatedScore.NumberOfCorrectedAnswer) / float64(countQuestion)) * 100
		if err := repo.db.Save(&updatedScore).Error; err != nil {
			return errors.New(err.Error())
		}
	}
	return nil
}
