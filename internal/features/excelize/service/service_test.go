package service_test

import (
	"eco_points/internal/features/excelize/service"
	"eco_points/mocks"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeExcel(t *testing.T) {
	qry := mocks.NewQueryMakeExcelInterface(t)
	srv := service.NewExcelMake(qry)

	t.Run("success make Excel", func(t *testing.T) {
		expectedName := "temp"
		expectedLink := "http://data.test"

		qry.On("MakeExcelFromDeposit", "deposit", "2024_8_1", uint(1), true, uint(10)).Return(expectedName, expectedLink, nil).Once()
		link, err := srv.MakeExcel("2024_8_1", "deposit", uint(1), true, uint(10))

		assert.Nil(t, err)
		assert.Equal(t, expectedLink, link)
	})

	t.Run("error in make Excel", func(t *testing.T) {
		expectedError := errors.New("some error")
		qry.On("MakeExcelFromDeposit", "deposit", "2024_8_1", uint(1), true, uint(10)).Return("", "", expectedError).Once()
		link, err := srv.MakeExcel("2024_8_1", "deposit", uint(1), true, uint(10))

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, "", link)

	})

	t.Run("error deleting file", func(t *testing.T) {
		nonExistentFilePath := "./nonexistentfile.txt"

		err := srv.DeleteFile(nonExistentFilePath)
		assert.NotNil(t, err)
		assert.True(t, os.IsNotExist(err))
	})
}
