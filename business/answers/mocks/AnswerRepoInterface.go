// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	answers "kuiz/business/answers"

	mock "github.com/stretchr/testify/mock"
)

// AnswerRepoInterface is an autogenerated mock type for the AnswerRepoInterface type
type AnswerRepoInterface struct {
	mock.Mock
}

// CreateAnswer provides a mock function with given fields: domain, ctx
func (_m *AnswerRepoInterface) CreateAnswer(domain answers.Answer, ctx context.Context) (answers.Answer, error) {
	ret := _m.Called(domain, ctx)

	var r0 answers.Answer
	if rf, ok := ret.Get(0).(func(answers.Answer, context.Context) answers.Answer); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(answers.Answer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(answers.Answer, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAnswer provides a mock function with given fields: id, domain, ctx
func (_m *AnswerRepoInterface) DeleteAnswer(id string, domain answers.Answer, ctx context.Context) (answers.Answer, error) {
	ret := _m.Called(id, domain, ctx)

	var r0 answers.Answer
	if rf, ok := ret.Get(0).(func(string, answers.Answer, context.Context) answers.Answer); ok {
		r0 = rf(id, domain, ctx)
	} else {
		r0 = ret.Get(0).(answers.Answer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, answers.Answer, context.Context) error); ok {
		r1 = rf(id, domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}