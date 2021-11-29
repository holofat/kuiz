package quizzes

import (
	"context"
	"errors"
	"kuiz/business/quizzes"

	"gorm.io/gorm"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(gormDb *gorm.DB) quizzes.QuizRepoInterface {
	return &QuizRepository{
		db: gormDb,
	}
}

func (repo *QuizRepository) CreateQuiz(domain quizzes.Domain, ctx context.Context) (quizzes.Domain, error) {
	quizDb := FromDomain(domain)

	if err := repo.db.Create(&quizDb).Error; err != nil {
		return quizzes.Domain{}, errors.New("Error in database")
	} else {
		return quizDb.ToDomain(), nil
	}
}

func (repo *QuizRepository) DeleteQuiz(id string, domain quizzes.Domain, ctx context.Context) (quizzes.Domain, error) {
	quiz := FromDomain(domain)

	err := repo.db.Where("id = ?", id).First(&quiz).Error
	if err != nil {
		return quizzes.Domain{}, errors.New("Record is not found")
	}
	repo.db.Delete(&quiz)
	return quiz.ToDomain(), nil
}

func (repo *QuizRepository) UpdateQuiz(id string, domain quizzes.Domain, ctx context.Context) (quizzes.Domain, error) {
	quiz := FromDomain(domain)
	err := repo.db.Save(&quiz).Error
	if err != nil {
		return quizzes.Domain{}, errors.New("Record is not found")
	}
	return quiz.ToDomain(), nil
}
