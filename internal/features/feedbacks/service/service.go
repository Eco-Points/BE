package service

import (
	"eco_points/internal/features/feedbacks"
	"errors"
)

type FeedbackServices struct {
	qry feedbacks.FQuery
}

func NewFeedbackService(q feedbacks.FQuery) feedbacks.FServices {
	return &FeedbackServices{
		qry: q,
	}
}

func (fs *FeedbackServices) AddFeedback(newFeedback feedbacks.Feedback) error {
	err := fs.qry.AddFeedback(newFeedback)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambahkan feedback")
	}
	return nil
}

func (fs *FeedbackServices) UpdateFeedback(feedbackID uint, updatedFeedback feedbacks.Feedback) error {
	err := fs.qry.UpdateFeedback(feedbackID, updatedFeedback)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat memperbarui feedback")
	}
	return nil
}

func (fs *FeedbackServices) DeleteFeedback(feedbackID uint) error {
	err := fs.qry.DeleteFeedback(feedbackID)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menghapus feedback")
	}
	return nil
}

func (fs *FeedbackServices) GetFeedbackByID(feedbackID uint) (feedbacks.Feedback, error) {
	feedback, err := fs.qry.GetFeedbackByID(feedbackID)
	if err != nil {
		return feedbacks.Feedback{}, errors.New("terjadi kesalahan pada server saat mengambil feedback berdasarkan ID")
	}
	return feedback, nil
}

func (fs *FeedbackServices) GetAllFeedbacks(limit int, offset int) ([]feedbacks.Feedback, int, error) {
	feedbacks, totalItems, err := fs.qry.GetAllFeedbacks(limit, offset)
	if err != nil {
		return nil, 0, errors.New("terjadi kesalahan pada server saat mengambil semua feedback")
	}
	return feedbacks, totalItems, nil
}
