package questions

import (
	"context"
	"errors"
	"fmt"
	"kuiz/business/questions"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(gormDb *gorm.DB) questions.QuestionRepoInterface {
	return &QuestionRepository{
		db: gormDb,
	}
}

func (repo *QuestionRepository) CreateQuestion(domain questions.Question, ctx context.Context) (questions.Question, error) {
	questionDb := FromDomain(domain)

	if err := repo.db.Create(&questionDb).Error; err != nil {
		return questions.Question{}, errors.New("Error in database")
	} else {
		return questionDb.ToDomain(), nil
	}
}

func (repo *QuestionRepository) GetQuestion(idQuiz string, ctx context.Context) ([]questions.Question, error) {
	var questionList listQuestion

	err := repo.db.Preload("Answer").Where("quiz_id  = ?", idQuiz).Find(&questionList).Error
	fmt.Println(questionList)
	if err != nil {
		return []questions.Question{}, errors.New("Error in database")
	} else {
		return questionList.ToDomainList(), nil
	}
}

func (repo *QuestionRepository) DeleteQuestion(id string, ctx context.Context) error {
	var question Question

	err := repo.db.Where("id = ?", id).First(&question).Delete(&question).Error
	if err != nil {
		return errors.New("Error in database")
	}
	return nil
}
