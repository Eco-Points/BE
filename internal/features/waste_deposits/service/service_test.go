package service_test

import (
	deposits "eco_points/internal/features/waste_deposits"
	"eco_points/internal/features/waste_deposits/service"
	"eco_points/internal/utils"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateWasteDepositStatus(t *testing.T) {
	email := mocks.NewGomailUtilityInterface(t)
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry, email)

	t.Run("success Update Deposit Status - verified", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "verified").Return(uint(1), 100, nil).Once()
		qry.On("GetUserEmailData", uint(1)).Return("a@gmail.com", "a", nil).Once()
		email.On("SendEmail", 100, "Halo, <b>a</b><br/> Pengajuan deposit sampahmu sudah diverifikasi oleh admin keren dari tim Eco Points. <br/>Selamat kamu mendapatkan <b>100 poin</b> ", "a@gmail.com", "a").Return(nil).Once()

		err := srv.UpdateWasteDepositStatus(uint(1), "verified")

		assert.Nil(t, err)

		qry.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("Success Update Waste Deposit Status - rejected", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "rejected").Return(uint(1), 100, nil).Once()
		qry.On("GetUserEmailData", uint(1)).Return("a@gmail.com", "a", nil).Once()
		email.On("SendEmail", 100, "Halo, <b>a</b><br/> Pengajuan deposit sampahmu belum memenuhi syarat untuk mendapatkan poin", "a@gmail.com", "a").Return(nil).Once()

		err := srv.UpdateWasteDepositStatus(uint(1), "rejected")

		assert.Nil(t, err)

		qry.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("failed Update Deposit Status - unhandled status", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "unknown").Return(uint(1), 0, nil).Once()
		qry.On("GetUserEmailData", uint(1)).Return("a@gmail.com", "a", nil).Once()
		email.On("SendEmail", 0, "Halo, <b>a</b><br/> Pengajuan deposit sampahmu belum memenuhi syarat untuk mendapatkan poin", "a@gmail.com", "a").Return(nil).Once()
		err := srv.UpdateWasteDepositStatus(uint(1), "unknown")

		assert.Nil(t, err)

		qry.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("failed Update Waste Deposit Status - GetUserEmailData", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "verified").Return(uint(1), 100, nil).Once()
		expectedError := errors.New("error getting user email data")
		qry.On("GetUserEmailData", uint(1)).Return("", "", expectedError).Once()

		err := srv.UpdateWasteDepositStatus(uint(1), "verified")

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedError)
		qry.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("failed Update Waste Deposit Status - SendEmail", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "verified").Return(uint(1), 100, nil).Once()
		qry.On("GetUserEmailData", uint(1)).Return("a@gmail.com", "a", nil).Once()
		expectedError := errors.New("error sending email")
		email.On("SendEmail", 100, "Halo, <b>a</b><br/> Pengajuan deposit sampahmu sudah diverifikasi oleh admin keren dari tim Eco Points. <br/>Selamat kamu mendapatkan <b>100 poin</b> ", "a@gmail.com", "a").Return(expectedError).Once()

		err := srv.UpdateWasteDepositStatus(uint(1), "verified")

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedError)
		qry.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("failed SendEmail with formatted message", func(t *testing.T) {
		qry.On("UpdateWasteDepositStatus", uint(1), "verified").Return(uint(1), 100, nil).Once()
		qry.On("GetUserEmailData", uint(1)).Return("a@gmail.com", "a", nil).Once()
		email.On("SendEmail", 100, "Halo, <b>a</b><br/> Pengajuan deposit sampahmu sudah diverifikasi oleh admin keren dari tim Eco Points. <br/>Selamat kamu mendapatkan <b>100 poin</b> ", "a@gmail.com", "a").Return(errors.New("email send error")).Once()
		err := srv.UpdateWasteDepositStatus(uint(1), "verified")

		assert.NotNil(t, err)
		assert.Equal(t, "email send error", err.Error())

		qry.AssertExpectations(t)
		email.AssertExpectations(t)
	})

}

func TestGetDeposit(t *testing.T) {
	email := utils.NewGomailUtility()
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry, email)

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

	t.Run("failed Get Deposit", func(t *testing.T) {
		expectedError := errors.New("error get deposit data")
		qry.On("GetUserDeposit", uint(1), uint(10), uint(0), true).Return(deposits.ListWasteDepositInterface{}, expectedError).Once()
		result, err := srv.GetUserDeposit(uint(1), uint(10), uint(0), true)

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, result, deposits.ListWasteDepositInterface{})
	})
}

func TestGetDepositbyId(t *testing.T) {
	email := utils.NewGomailUtility()
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry, email)

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

	t.Run("failed Get Deposit  By ID", func(t *testing.T) {
		expectedError := errors.New("error get deposit data")
		qry.On("GetDepositbyId", uint(1)).Return(deposits.WasteDepositInterface{}, expectedError).Once()
		result, err := srv.GetDepositbyId(uint(1))

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, result, deposits.WasteDepositInterface{})
	})
}

func TestDepositTrash(t *testing.T) {
	email := utils.NewGomailUtility()
	qry := mocks.NewQueryDepoInterface(t)
	srv := service.NewDepositsService(qry, email)

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
