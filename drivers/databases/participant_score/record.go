package participantscore

import (
	ps "kuiz/business/participant_score"
)

type ParticipantScore struct {
	Id                      int `gorm:"primaryKey"`
	UserId                  int
	QuizId                  int
	Score                   float64
	NumberOfCorrectedAnswer int
}

func (participantScore ParticipantScore) ToDomain() ps.ParticipantScore {
	return ps.ParticipantScore{
		Id:                      participantScore.Id,
		UserId:                  participantScore.UserId,
		QuizId:                  participantScore.QuizId,
		Score:                   participantScore.Score,
		NumberOfCorrectedAnswer: participantScore.NumberOfCorrectedAnswer,
	}
}

type listParticipantScore []ps.ParticipantScore

func (listScore *listParticipantScore) ToDomainList() []ps.ParticipantScore {
	var list []ps.ParticipantScore

	for _, k := range *listScore {
		list = append(list, k)
	}
	return list
}

func FromDomain(domain ps.ParticipantScore) ParticipantScore {
	return ParticipantScore{
		Id:                      domain.Id,
		UserId:                  domain.UserId,
		QuizId:                  domain.QuizId,
		Score:                   domain.Score,
		NumberOfCorrectedAnswer: domain.NumberOfCorrectedAnswer,
	}
}
