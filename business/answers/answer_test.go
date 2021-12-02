package answers_test

import (
	"context"
	"kuiz/business/answers"
	"kuiz/business/answers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var answerRepoInterfaceMock mocks.AnswerRepoInterface
var answerUsecaseInterface answers.AnswerUsecaseInterface
var answerDataDummy answers.Answer

func setup() {
	answerUsecaseInterface = answers.NewUsecase(&answerRepoInterfaceMock, time.Hour*1)
	answerDataDummy = answers.Answer{
		QuestionId:    1,
		Answer:        "test answer",
		CorrectAnswer: true,
	}
}

func TestCreateAnswer(t *testing.T) {
	setup()
	t.Run("Success Create Answer", func(t *testing.T) {
		answerRepoInterfaceMock.On("CreateAnswer", mock.AnythingOfType("answers.Answer"), mock.Anything).Return(answerDataDummy, nil).Once()

		var reqCreateAnswer = answers.Answer{
			QuestionId:    1,
			Answer:        "test answer",
			CorrectAnswer: true,
		}

		domain, err := answerUsecaseInterface.CreateAnswer(reqCreateAnswer, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, answerDataDummy, domain)
	})

	t.Run("Failed when answer is empty", func(t *testing.T) {
		var reqCreateAnswer = answers.Answer{
			QuestionId: 1,
			Answer:     "",
		}
		domain, err := answerUsecaseInterface.CreateAnswer(reqCreateAnswer, context.Background())

		assert.Equal(t, "answer must be filled", err.Error())
		assert.Equal(t, answers.Answer{}, domain)
	})
}
