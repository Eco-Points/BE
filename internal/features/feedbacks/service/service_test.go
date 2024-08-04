package service_test

import (
	"eco_points/internal/features/feedbacks"
	"eco_points/internal/features/feedbacks/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFeedback(t *testing.T) {
	mockQuery := new(mocks.FQuery)
	feedbackService := service.NewFeedbackService(mockQuery)

	feedbackRequest := feedbacks.Feedback{
		UserID: 1,
		Review: "This is a feedback",
		Rating: 5,
	}

	t.Run("success add feedback", func(t *testing.T) {
		mockQuery.On("AddFeedback", mock.AnythingOfType("feedbacks.Feedback")).Return(nil).Once()

		err := feedbackService.AddFeedback(feedbackRequest)
		assert.Nil(t, err)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("AddFeedback", mock.AnythingOfType("feedbacks.Feedback")).Return(errors.New("query error")).Once()

		err := feedbackService.AddFeedback(feedbackRequest)
		assert.NotNil(t, err)
		assert.Equal(t, "terjadi kesalahan pada server saat menambahkan feedback", err.Error())
	})
}

func TestUpdateFeedback(t *testing.T) {
	mockQuery := new(mocks.FQuery)
	feedbackService := service.NewFeedbackService(mockQuery)

	feedbackRequest := feedbacks.Feedback{
		UserID: 1,
		Review: "Updated feedback",
		Rating: 4,
	}

	t.Run("success update feedback", func(t *testing.T) {
		mockQuery.On("UpdateFeedback", uint(1), feedbackRequest).Return(nil).Once()

		err := feedbackService.UpdateFeedback(1, feedbackRequest)
		assert.Nil(t, err)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("UpdateFeedback", uint(1), feedbackRequest).Return(errors.New("query error")).Once()

		err := feedbackService.UpdateFeedback(1, feedbackRequest)
		assert.NotNil(t, err)
		assert.Equal(t, "terjadi kesalahan pada server saat memperbarui feedback", err.Error())
	})
}

func TestDeleteFeedback(t *testing.T) {
	mockQuery := new(mocks.FQuery)
	feedbackService := service.NewFeedbackService(mockQuery)

	t.Run("success delete feedback", func(t *testing.T) {
		mockQuery.On("DeleteFeedback", uint(1)).Return(nil).Once()

		err := feedbackService.DeleteFeedback(1)
		assert.Nil(t, err)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("DeleteFeedback", uint(1)).Return(errors.New("query error")).Once()

		err := feedbackService.DeleteFeedback(1)
		assert.NotNil(t, err)
		assert.Equal(t, "terjadi kesalahan pada server saat menghapus feedback", err.Error())
	})
}

func TestGetFeedbackByID(t *testing.T) {
	mockQuery := new(mocks.FQuery)
	feedbackService := service.NewFeedbackService(mockQuery)

	expectedFeedback := feedbacks.Feedback{
		ID:     1,
		UserID: 1,
		Review: "This is a feedback",
		Rating: 5,
	}

	t.Run("success get feedback by ID", func(t *testing.T) {
		mockQuery.On("GetFeedbackByID", uint(1)).Return(expectedFeedback, nil).Once()

		result, err := feedbackService.GetFeedbackByID(1)
		assert.Nil(t, err)
		assert.Equal(t, expectedFeedback, result)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("GetFeedbackByID", uint(1)).Return(feedbacks.Feedback{}, errors.New("query error")).Once()

		result, err := feedbackService.GetFeedbackByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, feedbacks.Feedback{}, result)
		assert.Equal(t, "terjadi kesalahan pada server saat mengambil feedback berdasarkan ID", err.Error())
	})
}

func TestGetAllFeedbacks(t *testing.T) {
	mockQuery := new(mocks.FQuery)
	feedbackService := service.NewFeedbackService(mockQuery)

	expectedFeedbacks := []feedbacks.Feedback{
		{
			ID:     1,
			UserID: 1,
			Review: "Feedback 1",
			Rating: 5,
		},
		{
			ID:     2,
			UserID: 2,
			Review: "Feedback 2",
			Rating: 4,
		},
	}
	expectedTotalItems := 2

	t.Run("success get all feedbacks", func(t *testing.T) {
		mockQuery.On("GetAllFeedbacks", 10, 0).Return(expectedFeedbacks, expectedTotalItems, nil).Once()

		results, totalItems, err := feedbackService.GetAllFeedbacks(10, 0)
		assert.Nil(t, err)
		assert.Equal(t, expectedFeedbacks, results)
		assert.Equal(t, expectedTotalItems, totalItems)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("GetAllFeedbacks", 10, 0).Return(nil, 0, errors.New("query error")).Once()

		results, totalItems, err := feedbackService.GetAllFeedbacks(10, 0)
		assert.NotNil(t, err)
		assert.Nil(t, results)
		assert.Equal(t, 0, totalItems)
		assert.Equal(t, "terjadi kesalahan pada server saat mengambil semua feedback", err.Error())
	})
}
