package handler

import (
	"eco_points/internal/features/dashboards"
)

type UserDashboardResponse struct {
	ID       uint   `json:"id"`
	CreateAt string `json:"registered_date"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

func ToUserDashboardResponse(input dashboards.User) UserDashboardResponse {
	return UserDashboardResponse{
		ID:       input.ID,
		Fullname: input.Fullname,
		CreateAt: input.CreateAt,
		Email:    input.Email,
		Status:   input.Status,
	}
}
