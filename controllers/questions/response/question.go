package response

import (
	"kuiz/business/questions"
	"time"
)

type QuestionResponse struct {
	Id               uint      `json:"id"`
	QuestionSentence string    `json:"question"`
	QuizId           uint      `json:"quiz_id"`
	CreatedAt        time.Time `json:"created_at"`
}

func FromDomain(domain questions.Question) QuestionResponse {
	return QuestionResponse{
		Id:               domain.Id,
		QuestionSentence: domain.QuestionSentence,
		QuizId:           domain.QuizId,
		CreatedAt:        domain.CreatedAt,
	}
}

func FromDomainList(domain []questions.Question) []QuestionResponse {
	var list []QuestionResponse
	for _, k := range domain {
		tempResp := QuestionResponse{
			Id:               k.Id,
			QuestionSentence: k.QuestionSentence,
			QuizId:           k.QuizId,
			CreatedAt:        k.CreatedAt,
		}
		list = append(list, tempResp)
	}
	return list
}
