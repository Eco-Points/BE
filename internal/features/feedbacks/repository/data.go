package repository

import (
	"eco_points/internal/features/feedbacks"
	"eco_points/internal/features/users"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	UserID uint
	Rating uint
	Review string
	User   users.User `gorm:"foreignKey:UserID"`
}

func (f *Feedback) ToFeedbackEntity() feedbacks.Feedback {
	return feedbacks.Feedback{
		ID:           f.ID,
		UserID:       f.UserID,
		UserFullname: f.User.Fullname,
		Rating:       f.Rating,
		Review:       f.Review,
		CreatedAt:    f.CreatedAt,
		UpdatedAt:    f.UpdatedAt,
	}
}

func ToFeedbackData(input feedbacks.Feedback) Feedback {
	return Feedback{
		UserID: input.UserID,
		Rating: input.Rating,
		Review: input.Review,
	}
}
