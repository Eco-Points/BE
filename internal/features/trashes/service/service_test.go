package service_test

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/features/trashes/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTrash(t *testing.T) {
	qry := mocks.NewQueryTrashInterface(t)
	cloud := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewTrashService(qry, cloud)

	returnTrash := trashes.ListTrashEntity{}
	dataList := struct {
		TrashType string
		Name      string
		Point     uint
		Descrip   string
		Image     string
		UserID    uint
		ID        uint
	}{
		TrashType: "plastik",
		Name:      "plastik biasa",
		Point:     20,
		Descrip:   "hanya plastik biasa",
		Image:     "http://",
		UserID:    1,
	}

	returnTrash = append(returnTrash, dataList)

	t.Run("success get trash by type", func(t *testing.T) {
		inputTrashQry := trashes.ListTrashEntity{}
		dataList := struct {
			TrashType string
			Name      string
			Point     uint
			Descrip   string
			Image     string
			UserID    uint
			ID        uint
		}{
			TrashType: "plastik",
			Name:      "plastik biasa",
			Point:     20,
			Descrip:   "hanya plastik biasa",
			Image:     "http://",
			UserID:    1,
		}

		inputTrashQry = append(inputTrashQry, dataList)

		qry.On("GetTrashbyType", "plastik").Return(returnTrash, nil).Once()
		returnTrash, err := srv.GetTrash("plastik")

		assert.Nil(t, err)
		assert.Equal(t, returnTrash, inputTrashQry)
	})

	t.Run("success all get trash ", func(t *testing.T) {
		inputTrashQry := trashes.ListTrashEntity{}
		dataList := struct {
			TrashType string
			Name      string
			Point     uint
			Descrip   string
			Image     string
			UserID    uint
			ID        uint
		}{
			TrashType: "plastik",
			Name:      "plastik biasa",
			Point:     20,
			Descrip:   "hanya plastik biasa",
			Image:     "http://",
			UserID:    1,
		}

		inputTrashQry = append(inputTrashQry, dataList)

		qry.On("GetTrashLimit").Return(returnTrash, nil).Once()
		returnTrash, err := srv.GetTrash("")

		assert.Nil(t, err)
		assert.Equal(t, returnTrash, inputTrashQry)
	})

	t.Run("failed get trash by type", func(t *testing.T) {
		// inputTrashQry := trashes.ListTrashEntity{}
		expectedError := errors.New("error on query / parameter")

		qry.On("GetTrashbyType", "not plastik").Return(trashes.ListTrashEntity{}, expectedError).Once()
		inputTrashQry, err := srv.GetTrash("not plastik")

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, trashes.ListTrashEntity{}, inputTrashQry)
	})

	t.Run("failed get all trash ", func(t *testing.T) {
		expectedError := errors.New("error on query / parameter")

		qry.On("GetTrashLimit").Return(trashes.ListTrashEntity{}, expectedError).Once()
		inputTrashQry, err := srv.GetTrash("")

		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, trashes.ListTrashEntity{}, inputTrashQry)
	})
}

func TestDeleteTrash(t *testing.T) {
	qry := mocks.NewQueryTrashInterface(t)
	cloud := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewTrashService(qry, cloud)
	t.Run("success delete trash", func(t *testing.T) {
		qry.On("DeleteTrash", uint(1)).Return(nil).Once()
		err := srv.DeleteTrash(uint(1))
		assert.Nil(t, err)
	})

	t.Run("failed delete trash", func(t *testing.T) {
		expectedError := errors.New("error on query / parameter")
		qry.On("DeleteTrash", uint(1)).Return(expectedError).Once()
		err := srv.DeleteTrash(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
	})

}
