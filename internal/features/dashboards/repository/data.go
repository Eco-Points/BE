package repository

import (
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
	Trash    []trashQuery.Trash       `gorm:"foreignKey:UserID"`
	Location []lochQuery.Location     `gorm:"foreignKey:UserID"`
	Deposit  []depoQuery.WasteDeposit `gorm:"foreignKey:UserID"`
}
