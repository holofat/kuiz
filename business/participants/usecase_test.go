package participants_test

import (
	"context"
	"kuiz/business/participants"
	"kuiz/business/participants/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var participantRepoInterfaceMock mocks.ParticipantRepoInterface
var participantUsecaseInterface participants.ParticipantUsecaseInterface
var participantDataDummy participants.Participant

func setup() {
	participantUsecaseInterface = participants.NewUsecase(&participantRepoInterfaceMock, time.Hour*1)
	participantDataDummy = participants.Participant{
		UserId:     1,
		QuizId:     1,
		QuestionId: 1,
		AnswerId:   1,
	}
}

func TestAnswerQuestion(t *testing.T) {
	setup()

	t.Run("Success answered", func(t *testing.T) {
		userDummyId := 1
		quizDummyId := 1
		questionDummyId := 1
		answerDummyId := 1
		participantRepoInterfaceMock.On("AnswerQuestion", userDummyId, quizDummyId, questionDummyId, answerDummyId, mock.Anything).Return(nil).Once()

		err := participantUsecaseInterface.AnswerQuestion(userDummyId, quizDummyId, questionDummyId, answerDummyId, context.Background())
		assert.Equal(t, nil, err)
	})
}
