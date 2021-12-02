// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	participantscore "kuiz/business/participant_score"

	mock "github.com/stretchr/testify/mock"
)

// ParticipantScoreRepoInterface is an autogenerated mock type for the ParticipantScoreRepoInterface type
type ParticipantScoreRepoInterface struct {
	mock.Mock
}

// FetchAllData provides a mock function with given fields: idQuiz, idAuthor, ctx
func (_m *ParticipantScoreRepoInterface) FetchAllData(idQuiz int, idAuthor int, ctx context.Context) ([]participantscore.ParticipantScore, error) {
	ret := _m.Called(idQuiz, idAuthor, ctx)

	var r0 []participantscore.ParticipantScore
	if rf, ok := ret.Get(0).(func(int, int, context.Context) []participantscore.ParticipantScore); ok {
		r0 = rf(idQuiz, idAuthor, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]participantscore.ParticipantScore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, context.Context) error); ok {
		r1 = rf(idQuiz, idAuthor, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchDataByIdParticipant provides a mock function with given fields: idQuiz, idParticipant, ctx
func (_m *ParticipantScoreRepoInterface) FetchDataByIdParticipant(idQuiz int, idParticipant int, ctx context.Context) (participantscore.ParticipantScore, error) {
	ret := _m.Called(idQuiz, idParticipant, ctx)

	var r0 participantscore.ParticipantScore
	if rf, ok := ret.Get(0).(func(int, int, context.Context) participantscore.ParticipantScore); ok {
		r0 = rf(idQuiz, idParticipant, ctx)
	} else {
		r0 = ret.Get(0).(participantscore.ParticipantScore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, context.Context) error); ok {
		r1 = rf(idQuiz, idParticipant, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}