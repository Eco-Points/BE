package repository

import (
	"eco_points/internal/features/feedbacks"

	"gorm.io/gorm"
)

type FeedbackModel struct {
	db *gorm.DB
}

func NewFeedbackModel(connection *gorm.DB) feedbacks.FQuery {
	return &FeedbackModel{
		db: connection,
	}
}

func (fm *FeedbackModel) AddFeedback(newFeedback feedbacks.Feedback) error {
	cnvData := ToFeedbackData(newFeedback)
	return fm.db.Create(&cnvData).Error
}

func (fm *FeedbackModel) UpdateFeedback(feedbackID uint, updateFeedback feedbacks.Feedback) error {
	cnvData := ToFeedbackData(updateFeedback)
	qry := fm.db.Where("id = ?", feedbackID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (fm *FeedbackModel) DeleteFeedback(feedbackID uint) error {
	qry := fm.db.Where("id = ?", feedbackID).Delete(&Feedback{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (fm *FeedbackModel) GetFeedbackByID(feedbackID uint) (feedbacks.Feedback, error) {
	var feedback Feedback
	err := fm.db.Preload("User").First(&feedback, feedbackID).Error
	if err != nil {
		return feedbacks.Feedback{}, err
	}
	return feedback.ToFeedbackEntity(), nil
}

func (fm *FeedbackModel) GetAllFeedbacks(limit int, offset int) ([]feedbacks.Feedback, int, error) {
	var feedbacksList []Feedback
	var totalItems int64

	err := fm.db.Model(&Feedback{}).Where("deleted_at IS NULL").Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	err = fm.db.Preload("User").Limit(limit).Offset(offset).Find(&feedbacksList).Error
	if err != nil {
		return nil, 0, err
	}

	feedbacksEntities := make([]feedbacks.Feedback, len(feedbacksList))
	for i, feedback := range feedbacksList {
		feedbacksEntities[i] = feedback.ToFeedbackEntity()
	}

	return feedbacksEntities, int(totalItems), nil
}
