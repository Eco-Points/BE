package repository

import (
	lochQuery "eco_points/internal/features/locations/repository"
	trashQuery "eco_points/internal/features/trashes/repository"
	"eco_points/internal/features/users"
	depoQuery "eco_points/internal/features/waste_deposits/repository"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string
	Password string
	Phone    string
	Address  string
	ImgURL   string
	Trash    []trashQuery.Trash       `gorm:"foreignKey:UserID"`
	Location []lochQuery.Location     `gorm:"foreignKey:UserID"`
	Deposit  []depoQuery.WasteDeposit `gorm:"foreignKey:UserID"`
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
		ImgURL:   input.ImgURL,
	}
}
