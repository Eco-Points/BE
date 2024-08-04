package service_test

import (
	"bytes"
	"eco_points/internal/features/trashes"
	"eco_points/internal/features/trashes/service"
	"eco_points/mocks"
	"errors"
	"log"
	"mime/multipart"
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

func TestUpdateTrash(t *testing.T) {
	qry := mocks.NewQueryTrashInterface(t)
	cloud := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewTrashService(qry, cloud)

	inputTrashQry := trashes.TrashEntity{
		TrashType: "plastik",
		Name:      "plastik biasa",
		Point:     "20",
		Descrip:   "hanya plastik biasa",
		Image:     "http://",
		UserID:    1,
	}
	t.Run("success update trash", func(t *testing.T) {

		qry.On("UpdateTrash", uint(1), inputTrashQry).Return(nil).Once()
		err := srv.UpdateTrash(uint(1), inputTrashQry)
		assert.Nil(t, err)
	})

	t.Run("failed update trash", func(t *testing.T) {
		expectedError := errors.New("error on query / parameter")

		qry.On("UpdateTrash", uint(1), inputTrashQry).Return(expectedError).Once()
		err := srv.UpdateTrash(uint(1), inputTrashQry)
		assert.NotNil(t, err)
	})
}

type mockFile struct {
	*bytes.Reader
}

func (m *mockFile) Close() error {
	return nil
}

func setupMockFile() (*multipart.FileHeader, multipart.File, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", "test.png")
	if err != nil {
		log.Println("error make form file ", err)
		return nil, nil, err
	}

	_, err = part.Write([]byte("mock file content"))
	if err != nil {
		log.Println("error make mock file ", err)
		return nil, nil, err
	}

	fileContent := &mockFile{Reader: bytes.NewReader(buffer.Bytes())}

	return &multipart.FileHeader{
		Filename: "test.png",
		Size:     int64(buffer.Len()),
		Header:   make(map[string][]string),
	}, fileContent, nil
}

func TestAddTrash(t *testing.T) {
	customFileHeader, fileContent, err := setupMockFile()
	assert.NoError(t, err, "Failed to set up mock file")

	qry := mocks.NewQueryTrashInterface(t)
	cloud := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewTrashService(qry, cloud)

	trashData := trashes.TrashEntity{
		Image: "http://cloudinary/test.png",
	}

	t.Run("Success Add Trash", func(t *testing.T) {
		cloud.On("FileOpener", customFileHeader).Return(fileContent, nil).Once()
		cloud.On("UploadToCloudinary", fileContent, customFileHeader.Filename).Return("http://cloudinary/test.png", nil).Once()
		qry.On("AddTrash", trashData).Return(nil).Once()

		err := srv.AddTrash(trashData, customFileHeader)

		assert.NoError(t, err, "AddTrash should not return an error")
		assert.Equal(t, "http://cloudinary/test.png", trashData.Image, "Image URL should be set correctly")
		// cloud.AssertExpectations(t)
		// qry.AssertExpectations(t)
	})

	t.Run("FileOpener Error", func(t *testing.T) {
		cloud.On("FileOpener", customFileHeader).Return(nil, errors.New("open error")).Once()
		// cloud.On("FileOpener", mock.AnythingOfType("*multipart.FileHeader")).Return(nil, errors.New("open error")).Once()

		err := srv.AddTrash(trashData, customFileHeader)

		assert.Error(t, err, "AddTrash should return an error on file open failure")
		// cloud.AssertExpectations(t)
		// qry.AssertNotCalled(t, "AddTrash", mock.Anything) // Ensure AddTrash was not called
	})

	t.Run("UploadToCloudinary Error", func(t *testing.T) {
		cloud.On("FileOpener", customFileHeader).Return(fileContent, nil).Once()
		cloud.On("UploadToCloudinary", fileContent, customFileHeader.Filename).Return("", errors.New("upload error")).Once()
		// cloud.On("FileOpener", mock.AnythingOfType("*multipart.FileHeader")).Return(fileContent, nil).Once()
		// cloud.On("UploadToCloudinary", mock.Anything, "test.png").Return("", errors.New("upload error")).Once()

		err := srv.AddTrash(trashData, customFileHeader)

		assert.Error(t, err, "AddTrash should return an error on upload failure")
		// cloud.AssertExpectations(t)
		// qry.AssertNotCalled(t, "AddTrash", mock.Anything) // Ensure AddTrash was not called
	})

	t.Run("AddTrash Error", func(t *testing.T) {
		cloud.On("FileOpener", customFileHeader).Return(fileContent, nil).Once()
		cloud.On("UploadToCloudinary", fileContent, customFileHeader.Filename).Return("http://cloudinary/test.png", nil).Once()
		qry.On("AddTrash", trashData).Return(errors.New("database error")).Once()

		err := srv.AddTrash(trashData, customFileHeader)

		assert.Error(t, err, "AddTrash should return an error on database failure")
		// cloud.AssertExpectations(t)
		// qry.AssertExpectations(t)
	})
}
