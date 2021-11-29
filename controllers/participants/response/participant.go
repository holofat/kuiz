package response

import "kuiz/business/participants"

type ParticipantResponse struct {
	Id         int `json:"id"`
	UserId     int `json:"user_id"`
	QuestionId int `json:"question_id"`
	QuizId     int `json:"quiz_id"`
	AnswerId   int `json:"answer_id"`
}

func FromDomain(domain participants.Participant) ParticipantResponse {
	return ParticipantResponse{
		Id:         domain.Id,
		QuestionId: domain.QuestionId,
		UserId:     domain.UserId,
		AnswerId:   domain.AnswerId,
		QuizId:     domain.QuizId,
	}
}
