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

func TestGetExchange(t *testing.T) {
	qry := mocks.NewExQuery(t)
	srv := service.NewExchangeService(qry)

	returnValue := []exchanges.ListExchangeInterface{}
	exchange := struct {
		ID           uint
		PointUsed    uint
		Fullname     string
		Reward       string
		ExchangeTime string
	}{
		ID:           1,
		PointUsed:    100,
		Fullname:     "anggi eko",
		Reward:       "Voucher Telkomsel",
		ExchangeTime: "2020-12-12 00:00:00",
	}
	returnValue = append(returnValue, exchange)

	t.Run("success Get exchange", func(t *testing.T) {
		qry.On("GetExchangeHistory", uint(1), true, uint(10)).Return(returnValue, nil).Once()
		result, err := srv.GetExchangeHistory(uint(1), true, uint(10))
		assert.Nil(t, err)
		assert.Equal(t, result, returnValue)
	})

	t.Run("fail Get exchange", func(t *testing.T) {
		expectedError := errors.New("failed to create exchange")
		qry.On("GetExchangeHistory", uint(1), true, uint(10)).Return([]exchanges.ListExchangeInterface{}, expectedError).Once()

		result, err := srv.GetExchangeHistory(uint(1), true, uint(10))
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, []exchanges.ListExchangeInterface{}, result)

	})
}
