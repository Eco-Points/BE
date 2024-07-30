package repository

import (
	deposits "eco_points/internal/features/waste_deposits"
	"time"

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
	Status   string
	ImgURL   string
}
type WasteDeposit struct {
	gorm.Model
	TrashID    uint   `gorm:"column:trash_id"`
	UserID     uint   `gorm:"column:user_id"`
	LocationID uint   `gorm:"column:location_id"`
	Quantity   uint   `gorm:"column:quantity"`
	Point      uint   `gorm:"column:point"`
	Status     string `gorm:"column:status"`
}

type ListWasteDeposit struct {
	gorm.Model
	ID       uint      `gorm:"column:id"`
	Type     string    `gorm:"column:trash_type"`
	Point    uint      `gorm:"column:point"`
	DepoTime time.Time `gorm:"column:updated_at"`
	Quantity uint      `gorm:"column:quantity"`
	Fullname string    `gorm:"column:fullname"`
	Status   string    `gorm:"column:status"`
	TrashID  uint      `gorm:"column:trash_id"`
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

func toWasteDepositInterface(dataDepo WasteDeposit, dataTrash Trash, dataUser User) deposits.WasteDepositInterface {
	return deposits.WasteDepositInterface{
		Type:     dataTrash.TrashType,
		Point:    dataDepo.Point,
		Quantity: dataDepo.Quantity,
		DepoTime: dataDepo.CreatedAt.String(),
		ID:       dataDepo.ID,
		Status:   dataDepo.Status,
		Fullname: dataUser.Fullname,
	}
}

func toWasteDepositListInterface(data []ListWasteDeposit) deposits.ListWasteDepositInterface {
	var result deposits.ListWasteDepositInterface
	for _, v := range data {
		dataList := struct {
			Type     string
			Point    uint
			DepoTime string
			Quantity uint
			ID       uint
			Status   string
			Fullname string
		}{
			Type:     v.Type,
			Point:    v.Point,
			Quantity: v.Quantity,
			DepoTime: v.CreatedAt.String(),
			ID:       v.ID,
			Status:   v.Status,
			Fullname: v.Fullname,
		}
		result = append(result, dataList)
	}

	return result
}
