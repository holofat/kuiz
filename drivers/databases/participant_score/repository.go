package participantscore

import (
	"context"
	"errors"
	"fmt"
	participant_score "kuiz/business/participant_score"

	"gorm.io/gorm"
)

type ParticipantScoreRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(gormDb *gorm.DB) participant_score.ParticipantScoreRepoInterface {
	return &ParticipantScoreRepository{
		db: gormDb,
	}
}

func (repo *ParticipantScoreRepository) FetchAllData(idQuiz int, idAuthor int, ctx context.Context) ([]participant_score.ParticipantScore, error) {
	var list listParticipantScore

	err := repo.db.Where("quiz_id = ?", idQuiz).Find(&list).Error

	if err != nil {
		return []participant_score.ParticipantScore{}, errors.New(err.Error())
	}
	return list.ToDomainList(), nil
}

func (repo *ParticipantScoreRepository) FetchDataByIdParticipant(idQuiz int, idParticipant int, ctx context.Context) (participant_score.ParticipantScore, error) {
	score := FromDomain(participant_score.ParticipantScore{})

	err := repo.db.Where("quiz_id = ? AND user_id = ?", idQuiz, idParticipant).First(&score).Error
	fmt.Println(score)
	if err != nil {
		return participant_score.ParticipantScore{}, err
	}

	return score.ToDomain(), nil
}
