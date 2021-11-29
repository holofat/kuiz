package request

import "kuiz/business/answers"

type DeleteAnswer struct {
	Id uint
}

func (answer *DeleteAnswer) ToDomain() *answers.Domain {
	return &answers.Domain{
		Id: answer.Id,
	}
}
