package request

import (
	"kuiz/business/answers"
)

type CreateAnswer struct {
	QuestionId    uint   `json:"question_id"`
	Answer        string `json:"answer"`
	CorrectAnswer bool   `json:"correct_answer"`
}

func (answer *CreateAnswer) ToDomain() *answers.Answer {
	return &answers.Answer{
		QuestionId:    answer.QuestionId,
		Answer:        answer.Answer,
		CorrectAnswer: answer.CorrectAnswer,
	}
}
