package repository

import (
	"eco_points/internal/features/excelize"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Exchange struct {
	gorm.Model
	PointUsed uint
	Fullname  string `gorm:"column:name"`
	Reward    string `gorm:"column:rewards"`
}

type WasteDeposit struct {
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

func ConvertDateTime(data string) string {
	find := strings.SplitAfter(data, ".")[1]
	depotime := strings.Replace(data, find, "", 1)
	find = strings.SplitAfter(depotime, ".")[1]
	depotime = strings.Replace(depotime, find, "", 1)
	return strings.Replace(depotime, ".", "", 1)
}

func toWasteDepositInterface(data []WasteDeposit) []excelize.ExcelDeposit {
	var returnValue []excelize.ExcelDeposit
	for _, v := range data {
		excelDeposit := excelize.ExcelDeposit{
			ID:       v.ID,
			Type:     v.Type,
			Point:    v.Point,
			DepoTime: ConvertDateTime(v.CreatedAt.String()),
			Quantity: v.Quantity,
			Fullname: v.Fullname,
			Status:   v.Status,
			TrashID:  v.TrashID,
		}
		returnValue = append(returnValue, excelDeposit)
	}
	return returnValue
}

func toExchangeInterface(data []Exchange) []excelize.ExcelExchange {
	var returnValue []excelize.ExcelExchange
	for _, v := range data {
		exchange := excelize.ExcelExchange{
			ID:           v.ID,
			PointUsed:    v.PointUsed,
			ExchangeTime: ConvertDateTime(v.CreatedAt.String()),
			Fullname:     v.Fullname,
			Reward:       v.Reward,
		}
		returnValue = append(returnValue, exchange)
	}
	return returnValue
}
