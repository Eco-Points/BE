package service_test

import (
	deposits "eco_points/internal/features/waste_deposits"
	"eco_points/internal/features/waste_deposits/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateDepositStatus(t *testing.T) {
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry)
	t.Run("success Update Deposit Status", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "accepted").Return(nil).Once()
		err := srv.UpdateWasteDepositStatus(uint(1), "accepted")

		assert.Nil(t, err)
	})

	t.Run("failed Update Deposit Status", func(t *testing.T) {
		expectedError := errors.New("error update deposit status")
		qry.On("UpdateWasteDepositStatus", uint(1), "accepted").Return(expectedError).Once()
		err := srv.UpdateWasteDepositStatus(uint(1), "accepted")

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
	})
}

func TestGetDeposit(t *testing.T) {
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry)

	expectedOutput := deposits.ListWasteDepositInterface{}
	dataList := struct {
		Type     string
		Point    uint
		DepoTime string
		Quantity uint
		ID       uint
		Status   string
		Fullname string
	}{
		Type:     "plastik",
		Point:    200,
		Quantity: 20,
		DepoTime: "2024-07-29 12:57:41",
		ID:       1,
		Status:   "active",
		Fullname: "anggi",
	}
	expectedOutput = append(expectedOutput, dataList)

	t.Run("success Get Deposit", func(t *testing.T) {
		qry.On("GetUserDeposit", uint(1), uint(10), uint(0), true).Return(expectedOutput, nil).Once()
		result, err := srv.GetUserDeposit(uint(1), uint(10), uint(0), true)

		assert.Nil(t, err)
		assert.Equal(t, result, expectedOutput)
	})

	t.Run("failed Update Deposit", func(t *testing.T) {
		expectedError := errors.New("error get deposit data")
		qry.On("GetUserDeposit", uint(1), uint(10), uint(0), true).Return(deposits.ListWasteDepositInterface{}, expectedError).Once()
		result, err := srv.GetUserDeposit(uint(1), uint(10), uint(0), true)

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, result, deposits.ListWasteDepositInterface{})
	})
}

func TestGetDepositbyId(t *testing.T) {
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry)

	expectedOutput := deposits.WasteDepositInterface{
		Type:     "plastik",
		Point:    200,
		Quantity: 20,
		DepoTime: "2024-07-29 12:57:41",
		ID:       1,
		Status:   "active",
		Fullname: "anggi",
	}

	t.Run("success Get Deposit By ID", func(t *testing.T) {
		qry.On("GetDepositbyId", uint(1)).Return(expectedOutput, nil).Once()
		result, err := srv.GetDepositbyId(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, result, expectedOutput)
	})

	t.Run("failed Update Deposit  By ID", func(t *testing.T) {
		expectedError := errors.New("error get deposit data")
		qry.On("GetDepositbyId", uint(1)).Return(deposits.WasteDepositInterface{}, expectedError).Once()
		result, err := srv.GetDepositbyId(uint(1))

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, result, deposits.WasteDepositInterface{})
	})
}

func TestDepositTrash(t *testing.T) {
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry)

	input := deposits.WasteDepositInterface{
		Type:       "plastik",
		Point:      200,
		Quantity:   20,
		DepoTime:   "2024-07-29 12:57:41",
		ID:         1,
		Status:     "pending",
		Fullname:   "anggi",
		TrashID:    1,
		UserID:     1,
		LocationID: 1,
	}

	t.Run("success Deposit Trash", func(t *testing.T) {
		qry.On("DepositTrash", input).Return(nil).Once()
		err := srv.DepositTrash(input)

		assert.Nil(t, err)
	})

	t.Run("failed Deposit Trash", func(t *testing.T) {
		expectedError := errors.New("error deposit data")
		qry.On("DepositTrash", input).Return(expectedError).Once()
		err := srv.DepositTrash(input)

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
	})
}