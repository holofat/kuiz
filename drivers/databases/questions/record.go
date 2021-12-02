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
	Answer           []answers.Answer `gorm:"foreignKey:question_id"`
}

func (q Question) ToDomain() questions.Question {
	return questions.Question{
		Id:               q.Id,
		CreatedAt:        q.CreatedAt,
		QuizId:           q.QuizId,
		QuestionSentence: q.QuestionSentence,
		Answer:           q.Answer,
	}
}

type listQuestion []questions.Question

func (q *listQuestion) ToDomainList() []questions.Question {
	var list []questions.Question
	for _, k := range *q {
		list = append(list, k)
	}
	return list
}

func FromDomain(domain questions.Question) Question {
	return Question{
		Id:               domain.Id,
		QuizId:           domain.QuizId,
		QuestionSentence: domain.QuestionSentence,
		Answer:           domain.Answer,
	}
}
