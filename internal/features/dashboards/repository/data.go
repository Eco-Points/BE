package repository

import (
	"eco_points/internal/features/dashboards"
	lochQuery "eco_points/internal/features/locations/repository"
	trashQuery "eco_points/internal/features/trashes/repository"
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
	IsAdmin  bool
	Point    uint
	ImgURL   string
	Status   string
	Trash    []trashQuery.Trash       `gorm:"foreignKey:UserID"`
	Location []lochQuery.Location     `gorm:"foreignKey:UserID"`
	Deposit  []depoQuery.WasteDeposit `gorm:"foreignKey:UserID"`
}

// dari database di pindah ke entity
func (u *User) ToUserEntity() dashboards.User {
	return dashboards.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		CreateAt: u.CreatedAt.Format("02-01-2006"),
		Email:    u.Email,
		Phone:    u.Phone,
		Address:  u.Address,
		Status:   u.Status,
	}
}

// dari entity pindah ke database
func ToUserQuery(input dashboards.User) User {
	return User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
		Address:  input.Address,
		Status:   input.Status,
	}
}
