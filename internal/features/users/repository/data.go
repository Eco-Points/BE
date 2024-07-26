package repository

import (
	trashQuery "eco_points/internal/features/trashes/repository"
	"eco_points/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string
	Password string
	Phone    string
	Address  string
	IsAdmin  bool
	Point    uint
	ImgURL   string
	UserID   []trashQuery.Trash `gorm:"foreignKey:UserID"`
}

// dari database di pindah ke entity
func (u *User) ToUserEntity() users.User {
	return users.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Address:  u.Address,
		IsAdmin:  u.IsAdmin,
		Point:    u.Point,
		ImgURL:   u.ImgURL,
	}
}

// dari entity pindah ke database
func ToUserQuery(input users.User) User {
	return User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Address:  input.Address,
		IsAdmin:  input.IsAdmin,
		Point:    input.Point,
		ImgURL:   input.ImgURL,
	}
}
