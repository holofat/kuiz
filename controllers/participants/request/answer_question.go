package request

import "kuiz/business/participants"

type AnswerQuestion struct {
	UserId     int `json:"user_id"`
	QuizId     int `json:"quiz_id"`
	QuestionId int `json:"question_id"`
	AnswerId   int `json:"answer_id"`
}

func (participant *AnswerQuestion) ToDomain() *participants.Participant {
	return &participants.Participant{
		QuizId:     participant.QuizId,
		AnswerId:   participant.AnswerId,
		QuestionId: participant.QuestionId,
		UserId:     participant.UserId,
	}
}
