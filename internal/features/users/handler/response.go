package handler

import "eco_points/internal/features/users"

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	ImgURL   string `json:"image_url"`
	IsAdmin  bool   `json:"is_admin"`
	Status   string `json:"status"`
	Point    uint   `json:"point"`
}

func ToLoginResponse(result users.User, token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}

func ToGetUserResponse(input users.User) UserResponse {
	return UserResponse{
		ID:       input.ID,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
		Address:  input.Address,
		ImgURL:   input.ImgURL,
		IsAdmin:  input.IsAdmin,
		Status:   input.Status,
		Point:    input.Point,
	}
}
