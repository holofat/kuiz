package request

import "kuiz/business/participants"

type AnswerQuestion struct {
	UserId     int
	QuizId     int
	QuestionId int
	AnswerId   int
}

func (participant *AnswerQuestion) ToDomain() *participants.Participant {
	return &participants.Participant{
		QuizId:     participant.QuizId,
		AnswerId:   participant.AnswerId,
		QuestionId: participant.QuestionId,
		UserId:     participant.UserId,
	}
}
