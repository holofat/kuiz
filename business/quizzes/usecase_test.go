package quizzes_test

import (
	"context"
	"kuiz/business/quizzes"
	"kuiz/business/quizzes/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var quizRepoInterfaceMock mocks.QuizRepoInterface
var quizUsecaseInterface quizzes.QuizUsecaseInterface
var quizDataDummy quizzes.Domain

func setup() {
	quizUsecaseInterface = quizzes.NewUsecase(&quizRepoInterfaceMock, time.Hour*1)
	quizDataDummy = quizzes.Domain{
		Id:        1,
		AuthorId:  1,
		TitleQuiz: "test",
		GivenTime: 30,
	}
}

func TestGetQuiz(t *testing.T) {
	setup()

	t.Run("Success Fetching Data", func(t *testing.T) {
		quizRepoInterfaceMock.On("GetQuiz", "1", mock.Anything).Return(quizDataDummy, nil).Once()

		domain, err := quizUsecaseInterface.GetQuiz("1", context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, quizDataDummy, domain)
	})
}

func TestCreateQuiz(t *testing.T) {
	setup()

	t.Run("Success Create A Quiz", func(t *testing.T) {
		quizRepoInterfaceMock.On("CreateQuiz", mock.AnythingOfType("quizzes.Domain"), mock.Anything).Return(quizDataDummy, nil).Once()

		var reqCreateData = quizzes.Domain{
			TitleQuiz: "test",
			GivenTime: 30,
		}
		domain, err := quizUsecaseInterface.CreateQuiz(reqCreateData, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, quizDataDummy, domain)
	})
}
