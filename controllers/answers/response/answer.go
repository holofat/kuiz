package response

import "kuiz/business/answers"

type AnswerResponse struct {
	Id         uint   `json:"id"`
	Answer     string `json:"answer"`
	QuestionId uint   `json:"question_id"`
}

func FromDomain(domain answers.Answer) AnswerResponse {
	return AnswerResponse{
		Id:         domain.Id,
		QuestionId: domain.QuestionId,
		Answer:     domain.Answer,
	}
}
