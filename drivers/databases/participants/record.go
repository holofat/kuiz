package participants

import (
	"kuiz/business/participants"
	"time"
)

type Participant struct {
	Id         int `gorm:"primaryKey"`
	CreatedAt  time.Time
	UserId     int
	QuizId     int
	QuestionId int
	AnswerId   int
}

func (participant Participant) ToDomain() participants.Participant {
	return participants.Participant{
		Id:         participant.Id,
		UserId:     participant.UserId,
		QuizId:     participant.QuizId,
		QuestionId: participant.QuestionId,
		AnswerId:   participant.AnswerId,
	}
}

func FromDomain(domain participants.Participant) Participant {
	return Participant{
		Id:         domain.Id,
		QuizId:     domain.QuizId,
		QuestionId: domain.QuestionId,
		AnswerId:   domain.AnswerId,
		UserId:     domain.UserId,
	}
}
