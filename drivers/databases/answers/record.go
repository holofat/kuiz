package answers

import (
	"kuiz/business/answers"
	"time"
)

type Answer struct {
	Id            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	QuestionId    uint
	Answer        string
	CorrectAnswer bool
}

func (answer Answer) ToDomain() answers.Answer {
	return answers.Answer{
		Id:            answer.Id,
		CreatedAt:     answer.CreatedAt,
		Answer:        answer.Answer,
		CorrectAnswer: answer.CorrectAnswer,
		QuestionId:    answer.QuestionId,
	}
}

func FromDomain(domain answers.Answer) Answer {
	return Answer{
		Id:            domain.Id,
		QuestionId:    domain.QuestionId,
		Answer:        domain.Answer,
		CorrectAnswer: domain.CorrectAnswer,
	}
}
