package handler

import (
	"eco_points/internal/features/feedbacks"
	"time"
)

type FeedbackResponse struct {
	ID           uint      `json:"feedback_id"`
	UserID       uint      `json:"user_id"`
	UserFullname string    `json:"user_fullname"`
	Rating       uint      `json:"rating"`
	Review       string    `json:"review"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToFeedbackResponse(feedback feedbacks.Feedback) FeedbackResponse {
	return FeedbackResponse{
		ID:           feedback.ID,
		UserID:       feedback.UserID,
		UserFullname: feedback.UserFullname,
		Rating:       feedback.Rating,
		Review:       feedback.Review,
		CreatedAt:    feedback.CreatedAt,
		UpdatedAt:    feedback.UpdatedAt,
	}
}

func ToFeedbacksResponse(feedbacks []feedbacks.Feedback) []FeedbackResponse {
	responses := make([]FeedbackResponse, len(feedbacks))
	for i, feedback := range feedbacks {
		responses[i] = ToFeedbackResponse(feedback)
	}
	return responses
}
