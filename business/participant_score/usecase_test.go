package participantscore_test

import (
	"context"
	participantscore "kuiz/business/participant_score"
	"kuiz/business/participant_score/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var participantscoreRepoInterfaceMock mocks.ParticipantScoreRepoInterface
var participantscoreUsecaseInterface participantscore.ParticipantScoreRepoInterface
var participantScoreDataDummy participantscore.ParticipantScore

func setup() {
	participantscoreUsecaseInterface = participantscore.NewUseCase(&participantscoreRepoInterfaceMock, time.Hour*1)

	participantScoreDataDummy = participantscore.ParticipantScore{
		UserId:                  1,
		QuizId:                  1,
		Score:                   99.9,
		NumberOfCorrectedAnswer: 4,
	}
}

func TestFetchDataByIdParticipant(t *testing.T) {
	setup()
	t.Run("Success Fetch Data", func(t *testing.T) {
		idQuiz := 1
		idParticipant := 1
		participantscoreRepoInterfaceMock.On("FetchDataByIdParticipant", idQuiz, idParticipant, mock.Anything).Return(participantScoreDataDummy, nil).Once()

		domain, err := participantscoreUsecaseInterface.FetchDataByIdParticipant(idQuiz, idParticipant, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, participantScoreDataDummy, domain)
	})
}

func TestFetchAllData(t *testing.T) {
	setup()
	var listScoreDataDummy []participantscore.ParticipantScore

	listScoreDataDummy = append(listScoreDataDummy, participantScoreDataDummy)
	listScoreDataDummy = append(listScoreDataDummy, participantscore.ParticipantScore{
		UserId:                  2,
		QuizId:                  1,
		Score:                   87.2,
		NumberOfCorrectedAnswer: 4,
	})

	t.Run("Success Fetch All Data", func(t *testing.T) {
		idQuiz := 1
		idAuthor := 1
		participantscoreRepoInterfaceMock.On("FetchAllData", idQuiz, idAuthor, mock.Anything).Return(listScoreDataDummy, nil).Once()

		domain, err := participantscoreUsecaseInterface.FetchAllData(idQuiz, idAuthor, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, listScoreDataDummy, domain)
	})
}
