package answers

import (
	"context"
	"errors"
	"kuiz/business/answers"

	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(gormDb *gorm.DB) answers.AnswerRepoInterface {
	return &AnswerRepository{
		db: gormDb,
	}
}

func (repo *AnswerRepository) CreateAnswer(domain answers.Domain, ctx context.Context) (answers.Domain, error) {
	answerDb := FromDomain(domain)

	if err := repo.db.Create(&answerDb).Error; err != nil {
		return answers.Domain{}, errors.New("Error in database")
	} else {
		return answerDb.ToDomain(), nil
	}
}

func (repo *AnswerRepository) DeleteAnswer(id string, domain answers.Domain, ctx context.Context) (answers.Domain, error) {
	answer := FromDomain(domain)

	err := repo.db.Where("id = ?", id).First(&answer).Delete(&answer).Error
	if err != nil {
		return answers.Domain{}, errors.New("Error in database")
	}
	return answer.ToDomain(), nil
}
