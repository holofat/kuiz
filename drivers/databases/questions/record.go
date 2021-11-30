package questions

import (
	"kuiz/business/answers"
	"kuiz/business/questions"
	"time"
)

type Question struct {
	Id               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	QuizId           uint
	QuestionSentence string
	Answers          []answers.Answer `gorm:"foreignKey:QuestionId"`
}

func (q Question) ToDomain() questions.Question {
	return questions.Question{
		Id:               q.Id,
		CreatedAt:        q.CreatedAt,
		QuizId:           q.QuizId,
		QuestionSentence: q.QuestionSentence,
	}
}

type listQuestion []questions.Question

func (q *listQuestion) ToDomainList() []questions.Question {
	var list []questions.Question
	for _, k := range *q {
		tempResp := questions.Question{
			Id:               k.Id,
			QuestionSentence: k.QuestionSentence,
			QuizId:           k.QuizId,
			CreatedAt:        k.CreatedAt,
		}
		list = append(list, tempResp)
	}
	return list
}

func FromDomain(domain questions.Question) Question {
	return Question{
		Id:               domain.Id,
		QuizId:           domain.QuizId,
		QuestionSentence: domain.QuestionSentence,
	}
}
