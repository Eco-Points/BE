package service_test

import (
	"bytes"
	"eco_points/internal/features/trashes"
	"eco_points/internal/features/trashes/service"
	"eco_points/mocks"
	"errors"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

// MockMultipartFile implements the multipart.File interface
type MockMultipartFile struct {
	*bytes.Reader
}

func (m *MockMultipartFile) Close() error {
	return nil
}

type CustomMultipartFileHeader struct {
	*multipart.FileHeader
	MockFile multipart.File
}

func (f *CustomMultipartFileHeader) Open() (multipart.File, error) {
	return f.MockFile, nil
}

func TestAddTrash(t *testing.T) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	defer w.Close()

	// Create a mock file header
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="test.png"`)
	h.Set("Content-Type", "image/png")

	// Write the mock file content
	fileWriter, err := w.CreateFormFile("file", "test.png")
	assert.NoError(t, err)
	_, err = fileWriter.Write([]byte("fake image content"))
	assert.NoError(t, err)

	// Create a mock file
	mockFile := &MockMultipartFile{bytes.NewReader(b.Bytes())}

	// Create a custom file header
	customFileHeader := &CustomMultipartFileHeader{
		FileHeader: &multipart.FileHeader{
			Filename: "test.png",
			Header:   h,
			Size:     123,
		},
		MockFile: mockFile,
	}

	qry := mocks.NewQueryTrashInterface(t)
	cloud := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewTrashService(qry, cloud)

	trashData := trashes.TrashEntity{
		Image: "http://cloudinary/test.png",
	}
	t.Run("Success Add Trash", func(t *testing.T) {
		cloud.On("UploadToCloudinary", mock.AnythingOfType("*bytes.Reader"), "test.png").Return("http://cloudinary/test.png", nil)
		qry.On("AddTrash", mock.Anything).Return(nil)

		err := srv.AddTrash(trashData, customFileHeader.FileHeader)

		assert.NoError(t, err)
		assert.Equal(t, "http://cloudinary/test.png", trashData.Image)
		cloud.AssertExpectations(t)
		qry.AssertExpectations(t)
	})
}
