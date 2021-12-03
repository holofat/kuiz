package response

import (
	"kuiz/business/answers"
	"kuiz/business/questions"
	"time"
)

type QuestionResponse struct {
	Id               uint             `json:"id"`
	QuestionSentence string           `json:"question"`
	QuizId           uint             `json:"quiz_id"`
	CreatedAt        time.Time        `json:"created_at"`
	Answer           []answers.Answer `json:"answers"`
}

func FromDomain(domain questions.Question) QuestionResponse {
	return QuestionResponse{
		Id:               domain.Id,
		QuestionSentence: domain.QuestionSentence,
		QuizId:           domain.QuizId,
		CreatedAt:        domain.CreatedAt,
		Answer:           domain.Answer,
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
			Answer:           k.Answer,
		}
		list = append(list, tempResp)
	}
	return list
}
