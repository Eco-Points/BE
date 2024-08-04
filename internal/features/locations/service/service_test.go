package service_test

import (
	"eco_points/internal/features/locations"
	"eco_points/internal/features/locations/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLocation(t *testing.T) {
	qry := mocks.NewQueryLocInterface(t)
	srv := service.NewLocService(qry)

	inputLocQry := locations.LocationInterface{
		Address:    "Location A",
		Long:       "106.916666",
		Lat:        "-6.100000",
		Status:     "active",
		Start_time: "08:00:00",
		End_time:   "10:00:00",
		Phone:      "123456789",
		UserID:     1,
	}

	t.Run("success Add Location", func(t *testing.T) {
		qry.On("AddLocation", inputLocQry).Return(nil).Once()
		err := srv.AddLocation(inputLocQry)
		assert.Nil(t, err)
	})

	t.Run("failed Add Location", func(t *testing.T) {
		expectedError := errors.New("error on query / parameter")
		qry.On("AddLocation", inputLocQry).Return(expectedError).Once()
		err := srv.AddLocation(inputLocQry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedError)
	})
}

func TestGetLocation(t *testing.T) {
	qry := mocks.NewQueryLocInterface(t)
	srv := service.NewLocService(qry)
	expectedOutput := []locations.LocationInterface{
		{
			Address:    "Location A",
			Long:       "106.916666",
			Lat:        "-6.100000",
			Status:     "active",
			Start_time: "08:00:00",
			End_time:   "10:00:00",
			Phone:      "123456789",
			UserID:     1,
		},
	}
	t.Run("success Get Location", func(t *testing.T) {
		qry.On("GetLocation", uint(1)).Return(expectedOutput, nil).Once()
		result, err := srv.GetLocation(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, expectedOutput, result)
		qry.AssertExpectations(t)
	})

	t.Run("failed Get Location", func(t *testing.T) {
		expectedError := errors.New("error on query / parameter")
		qry.On("GetLocation", uint(1)).Return([]locations.LocationInterface{}, expectedError).Once()
		result, err := srv.GetLocation(uint(1))
		assert.NotNil(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, expectedError)
		qry.AssertExpectations(t)
	})
}
