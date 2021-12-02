package questions_test

import (
	"context"
	"kuiz/business/questions"
	"kuiz/business/questions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var questionRepoInterfaceMock mocks.QuestionRepoInterface
var questionUsecaseInterface questions.QuestionUsecaseInterface
var questionDataDummy questions.Question

func setup() {
	questionUsecaseInterface = questions.NewUsecase(&questionRepoInterfaceMock, time.Hour*1)
	questionDataDummy = questions.Question{
		QuizId:           1,
		QuestionSentence: "test question",
	}
}

func TestCreateQuestion(t *testing.T) {
	setup()

	t.Run("Success Create Question", func(t *testing.T) {
		questionRepoInterfaceMock.On("CreateQuestion", mock.AnythingOfType("questions.Question"), mock.Anything).Return(questionDataDummy, nil).Once()
		var reqCreateQuestion = questions.Question{
			QuizId:           1,
			QuestionSentence: "test question",
		}

		domain, err := questionUsecaseInterface.CreateQuestion(reqCreateQuestion, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, reqCreateQuestion, domain)
	})
}

func TestDeleteQuestion(t *testing.T) {
	setup()

	t.Run("Success Delete Question", func(t *testing.T) {
		idDummy := "1"
		questionRepoInterfaceMock.On("DeleteQuestion", idDummy, mock.Anything).Return(nil)

		err := questionUsecaseInterface.DeleteQuestion(idDummy, context.Background())
		assert.Equal(t, nil, err)
	})
}
