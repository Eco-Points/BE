package repository

import (
	deposits "eco_points/internal/features/waste_deposits"

	"gorm.io/gorm"
)

type WasteDeposit struct {
	gorm.Model
	TrashID    uint
	UserID     uint
	LocationID uint
	Quantity   uint
	Point      uint
	Status     string
}

type Trash struct {
	gorm.Model
	TrashType string `gorm:"column:trash_type"`
	Name      string `gorm:"column:trash_name"`
	Point     uint
	Descrip   string `gorm:"column:description"`
	Image     string
	UserID    uint
}

func toWasteDeposit(data deposits.WasteDepositInterface) WasteDeposit {
	return WasteDeposit{
		TrashID:    data.TrashID,
		UserID:     data.UserID,
		LocationID: data.LocationID,
		Quantity:   data.Quantity,
		Point:      (data.Point * data.Quantity),
		Status:     "pending",
	}
}
