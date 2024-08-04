package handler

import "eco_points/internal/features/feedbacks"

type CreateOrUpdateFeedbackRequest struct {
	Rating uint   `json:"rating" form:"rating"`
	Review string `json:"review" form:"review"`
}

func ToFeedbackModel(req CreateOrUpdateFeedbackRequest) feedbacks.Feedback {
	return feedbacks.Feedback{
		Rating: req.Rating,
		Review: req.Review,
	}
}
