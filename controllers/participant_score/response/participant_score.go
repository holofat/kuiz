package response

import (
	participant_score "kuiz/business/participant_score"
)

type ParticipantScoreResponse struct {
	UserId int     `json:"participant_id"`
	Score  float64 `json:"score"`
}

func FromDomain(domain participant_score.ParticipantScore) ParticipantScoreResponse {
	return ParticipantScoreResponse{
		UserId: domain.UserId,
		Score:  domain.Score,
	}
}

func FromDomainList(domain []participant_score.ParticipantScore) []ParticipantScoreResponse {
	var list []ParticipantScoreResponse

	for _, k := range domain {
		list = append(list, ParticipantScoreResponse{
			UserId: k.UserId,
			Score:  k.Score,
		})
	}
	return list
}
