package repository

import (
	deposits "eco_points/internal/features/waste_deposits"
	"fmt"
	"time"

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

type ListWasteDeposit []struct {
	gorm.Model
	ID       uint      `gorm:"column:id"`
	Type     string    `gorm:"column:trash_type"`
	Point    uint      `gorm:"column:point"`
	DepoTime time.Time `gorm:"column:updated_at"`
	Quantity uint      `gorm:"column:quantity"`
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

func toWasteDepositInterface(data WasteDeposit) deposits.WasteDepositInterface {
	return deposits.WasteDepositInterface{
		TrashID:    data.TrashID,
		UserID:     data.UserID,
		LocationID: data.LocationID,
		Quantity:   data.Quantity,
		Point:      (data.Point * data.Quantity),
		Status:     data.Status,
	}
}

func toWasteDepositListInterface(data ListWasteDeposit) deposits.ListWasteDepositInterface {
	var result deposits.ListWasteDepositInterface
	for _, v := range data {
		dataList := struct {
			Type     string
			Point    uint
			DepoTime string
			Quantity uint
			ID       uint
			Status   string
		}{
			Type:     v.Type,
			Point:    v.Point,
			Quantity: v.Quantity,
			DepoTime: fmt.Sprintf("%v", v.DepoTime),
			ID:       v.ID,
		}
		result = append(result, dataList)
	}

	return result
}
