package service_test

import (
	"eco_points/internal/features/rewards"
	"eco_points/internal/features/rewards/service"
	"eco_points/mocks"
	"errors"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddReward(t *testing.T) {
	mockQuery := new(mocks.RQuery)
	mockCloudinary := new(mocks.CloudinaryUtilityInterface)
	rewardService := service.NewRewardService(mockQuery, mockCloudinary)

	rewardRequest := rewards.Reward{
		Name:          "Reward 1",
		Description:   "Description 1",
		PointRequired: 100,
		Stock:         10,
	}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="test.png"`)
	h.Set("Content-Type", "image/png")

	file := &multipart.FileHeader{
		Filename: "test.png",
		Header:   h,
		Size:     123,
	}

	src, _ := file.Open()
	defer src.Close()

	t.Run("success add reward", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockCloudinary.On("UploadToCloudinary", src, file.Filename).Return("http://image.url", nil).Once()
		mockQuery.On("AddReward", mock.AnythingOfType("rewards.Reward")).Return(nil).Once()

		err := rewardService.AddReward(1, rewardRequest, src, file.Filename)
		assert.Nil(t, err)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(false, nil).Once()

		err := rewardService.AddReward(1, rewardRequest, src, file.Filename)
		assert.NotNil(t, err)
		assert.Equal(t, "forbidden: user is not an admin", err.Error())
	})

	t.Run("failure - cloudinary upload error", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockCloudinary.On("UploadToCloudinary", src, file.Filename).Return("", errors.New("upload error")).Once()

		err := rewardService.AddReward(1, rewardRequest, src, file.Filename)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal mengunggah gambar ke cloudinary", err.Error())
	})

	t.Run("failure - invalid file format", func(t *testing.T) {
		invalidFile := &multipart.FileHeader{
			Filename: "invalid_file.txt",
			Header:   h,
			Size:     123,
		}
		src, _ := invalidFile.Open()
		defer src.Close()

		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockCloudinary.On("UploadToCloudinary", src, invalidFile.Filename).Return("", errors.New("invalid file format")).Once()

		err := rewardService.AddReward(1, rewardRequest, src, invalidFile.Filename)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal mengunggah gambar ke cloudinary", err.Error())
	})

}

func TestUpdateReward(t *testing.T) {
	mockQuery := new(mocks.RQuery)
	mockCloudinary := new(mocks.CloudinaryUtilityInterface)
	rewardService := service.NewRewardService(mockQuery, mockCloudinary)

	rewardRequest := rewards.Reward{
		Name:          "Updated Reward",
		Description:   "Updated Description",
		PointRequired: 200,
		Stock:         20,
	}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="test.png"`)
	h.Set("Content-Type", "image/png")

	file := &multipart.FileHeader{
		Filename: "test.png",
		Header:   h,
		Size:     123,
	}

	src, _ := file.Open()
	defer src.Close()

	t.Run("success update reward", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockCloudinary.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
		mockQuery.On("UpdateReward", uint(1), rewardRequest).Return(nil).Once()

		err := rewardService.UpdateReward(1, 1, rewardRequest, src, file.Filename)
		assert.Nil(t, err)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(false, nil).Once()

		err := rewardService.UpdateReward(1, 1, rewardRequest, src, file.Filename)
		assert.NotNil(t, err)
		assert.Equal(t, "forbidden: user is not an admin", err.Error())
	})

	t.Run("failure - cloudinary upload error", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockCloudinary.On("UploadToCloudinary", src, file.Filename).Return("", errors.New("upload error")).Once()

		err := rewardService.UpdateReward(1, 1, rewardRequest, src, file.Filename)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to upload image to Cloudinary", err.Error())
	})

	t.Run("failure - invalid file format", func(t *testing.T) {
		invalidFile := &multipart.FileHeader{
			Filename: "invalid_file.txt",
			Header:   h,
			Size:     123,
		}
		src, _ := invalidFile.Open()
		defer src.Close()

		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockCloudinary.On("UploadToCloudinary", src, invalidFile.Filename).Return("", errors.New("invalid file format")).Once()

		err := rewardService.UpdateReward(1, 1, rewardRequest, src, invalidFile.Filename)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to upload image to Cloudinary", err.Error())
	})
}

func TestDeleteReward(t *testing.T) {
	mockQuery := new(mocks.RQuery)
	mockCloudinary := new(mocks.CloudinaryUtilityInterface)
	rewardService := service.NewRewardService(mockQuery, mockCloudinary)

	t.Run("success delete reward", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("DeleteReward", uint(1)).Return(nil).Once()

		err := rewardService.DeleteReward(1, 1)
		assert.Nil(t, err)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(false, nil).Once()

		err := rewardService.DeleteReward(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "forbidden: user is not an admin", err.Error())
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("IsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("DeleteReward", uint(1)).Return(errors.New("query error")).Once()

		err := rewardService.DeleteReward(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "server error while deleting reward", err.Error())
	})
}

func TestGetRewardByID(t *testing.T) {
	mockQuery := new(mocks.RQuery)
	mockCloudinary := new(mocks.CloudinaryUtilityInterface)
	rewardService := service.NewRewardService(mockQuery, mockCloudinary)

	expectedReward := rewards.Reward{
		ID:            1,
		Name:          "Reward 1",
		Description:   "Description 1",
		PointRequired: 100,
		Stock:         10,
		Image:         "http://image.url",
	}

	t.Run("success get reward by ID", func(t *testing.T) {
		mockQuery.On("GetRewardByID", uint(1)).Return(expectedReward, nil).Once()

		result, err := rewardService.GetRewardByID(1)
		assert.Nil(t, err)
		assert.Equal(t, expectedReward, result)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("GetRewardByID", uint(1)).Return(rewards.Reward{}, errors.New("query error")).Once()

		result, err := rewardService.GetRewardByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, rewards.Reward{}, result)
		assert.Equal(t, "terjadi kesalahan pada server saat mengambil data reward", err.Error())
	})
}

func TestGetAllRewards(t *testing.T) {
	mockQuery := new(mocks.RQuery)
	mockCloudinary := new(mocks.CloudinaryUtilityInterface)
	rewardService := service.NewRewardService(mockQuery, mockCloudinary)

	expectedRewards := []rewards.Reward{
		{
			ID:            1,
			Name:          "Reward 1",
			Description:   "Description 1",
			PointRequired: 100,
			Stock:         10,
			Image:         "http://image.url",
		},
		{
			ID:            2,
			Name:          "Reward 2",
			Description:   "Description 2",
			PointRequired: 200,
			Stock:         20,
			Image:         "http://image2.url",
		},
	}
	expectedTotalItems := 2

	t.Run("success get all rewards", func(t *testing.T) {
		mockQuery.On("GetAllRewards", 10, 0, "").Return(expectedRewards, expectedTotalItems, nil).Once()

		results, totalItems, err := rewardService.GetAllRewards(10, 0, "")
		assert.Nil(t, err)
		assert.Equal(t, expectedRewards, results)
		assert.Equal(t, expectedTotalItems, totalItems)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("GetAllRewards", 10, 0, "").Return(nil, 0, errors.New("query error")).Once()

		results, totalItems, err := rewardService.GetAllRewards(10, 0, "")
		assert.NotNil(t, err)
		assert.Nil(t, results)
		assert.Equal(t, 0, totalItems)
		assert.Equal(t, "terjadi kesalahan pada server saat mengambil semua reward", err.Error())
	})
}
