package feedbacks

import (
	"eco_points/internal/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Feedback struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	UserFullname string
	User         users.User `gorm:"foreignKey:UserID"`
	Rating       uint
	Review       string
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}

type FHandler interface {
	AddFeedback() echo.HandlerFunc
	UpdateFeedback() echo.HandlerFunc
	DeleteFeedback() echo.HandlerFunc
	GetFeedbackByID() echo.HandlerFunc
	GetAllFeedbacks() echo.HandlerFunc
}

type FServices interface {
	AddFeedback(newFeedback Feedback) error
	UpdateFeedback(feedbackID uint, updatedFeedback Feedback) error
	DeleteFeedback(feedbackID uint) error
	GetFeedbackByID(feedbackID uint) (Feedback, error)
	GetAllFeedbacks(limit int, offset int) ([]Feedback, int, error)
}

type FQuery interface {
	AddFeedback(newFeedback Feedback) error
	UpdateFeedback(feedbackID uint, updatedFeedback Feedback) error
	DeleteFeedback(feedbackID uint) error
	GetFeedbackByID(feedbackID uint) (Feedback, error)
	GetAllFeedbacks(limit int, offset int) ([]Feedback, int, error)
}
