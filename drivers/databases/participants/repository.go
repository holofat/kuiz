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

func (repo *ParticipantRepository) AnswerQuestion(domain participants.Participant, ctx context.Context) error {
	answered := FromDomain(domain)

	err := repo.db.Create(&answered).Error
	if err != nil {
		return errors.New("Error in database")
	}
	return nil
}
