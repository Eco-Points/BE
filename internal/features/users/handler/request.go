package handler

import "eco_points/internal/features/users"

type RegisterRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	ImgURL   string
}

func RegisterToUser(ur RegisterRequest) users.User {
	return users.User{
		Fullname: ur.Fullname,
		Email:    ur.Email,
		Password: ur.Password,
	}
}

func UpdateToUser(ur UpdateRequest) users.User {
	return users.User{
		Fullname: ur.Fullname,
		Email:    ur.Email,
		//	Password: ur.Password,
		Phone:   ur.Phone,
		Address: ur.Address,
		ImgURL:  ur.ImgURL,
	}
}
