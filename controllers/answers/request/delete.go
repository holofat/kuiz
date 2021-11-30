package request

import "kuiz/business/answers"

type DeleteAnswer struct {
	Id uint
}

func (answer *DeleteAnswer) ToDomain() *answers.Answer {
	return &answers.Answer{
		Id: answer.Id,
	}
}
