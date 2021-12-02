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

func (repo *QuizRepository) GetQuiz(id string, ctx context.Context) (quizzes.Domain, error) {
	var quiz Quiz

	err := repo.db.Preload("Answer").Preload("Question").Where("id = ?", id).First(&quiz).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return quizzes.Domain{}, errors.New("id is invalid")
		}
		return quizzes.Domain{}, errors.New("error in database")
	}
	return quiz.ToDomain(), nil
}

func (repo *QuizRepository) CreateQuiz(domain quizzes.Domain, ctx context.Context) (quizzes.Domain, error) {
	quizDb := FromDomain(domain)

	if err := repo.db.Create(&quizDb).Error; err != nil {
		return quizzes.Domain{}, errors.New("error in database")
	} else {
		return quizDb.ToDomain(), nil
	}
}

func (repo *QuizRepository) DeleteQuiz(id string, currentUserId uint, domain quizzes.Domain, ctx context.Context) (quizzes.Domain, error) {
	quiz := FromDomain(domain)

	err := repo.db.Where("id = ?", id).First(&quiz).Error
	if err != nil {
		return quizzes.Domain{}, errors.New("record is not found")
	}
	if quiz.AuthorId != currentUserId {
		return quizzes.Domain{}, errors.New("you can't delete other's quiz")
	}
	repo.db.Delete(&quiz)
	return quiz.ToDomain(), nil
}

func (repo *QuizRepository) UpdateQuiz(id string, domain quizzes.Domain, ctx context.Context) (quizzes.Domain, error) {
	quiz := FromDomain(domain)
	err := repo.db.Save(&quiz).Error
	if err != nil {
		return quizzes.Domain{}, errors.New("record is not found")
	}
	return quiz.ToDomain(), nil
}
