package service_test

import (
	"eco_points/internal/features/exchanges"
	"eco_points/internal/features/exchanges/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddExchange(t *testing.T) {
	qry := mocks.NewExQuery(t)
	srv := service.NewExchangeService(qry)

	newExchange := exchanges.Exchange{
		RewardID:  1,
		UserID:    2,
		PointUsed: 100,
	}

	t.Run("success add exchange", func(t *testing.T) {
		qry.On("AddExchange", newExchange).Return(nil).Once()

		err := srv.AddExchange(newExchange)
		assert.Nil(t, err)
		qry.AssertExpectations(t)
	})

	t.Run("fail add exchange", func(t *testing.T) {
		expectedError := errors.New("failed to create exchange")
		qry.On("AddExchange", newExchange).Return(expectedError).Once()

		err := srv.AddExchange(newExchange)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to create exchange", err.Error())
		qry.AssertExpectations(t)
	})
}
