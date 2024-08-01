package excelize

import (
	"github.com/labstack/echo/v4"
)

type ExcelDeposit struct {
	ID         uint
	TrashID    uint
	UserID     uint
	LocationID uint
	Quantity   uint
	Point      uint
	Status     string
	Type       string
	DepoTime   string
	Fullname   string
}

type ExcelExchange struct {
	ID           uint
	PointUsed    uint
	Fullname     string
	Reward       string
	ExchangeTime string
}

// type DataExcelInterface interface{}

type ServiceMakeExcelInterface interface {
	MakeExcel(date string, table string, userid uint, isVerif bool, limit uint) (string, error)
}

type HandlerMakeExcelInterface interface {
	MakeExcel() echo.HandlerFunc
}

type QueryMakeExcelInterface interface {
	SetDbSchema(schema string)
	MakeExcelFromDeposit(table string, date string, userid uint, isVerif bool, limit uint) (string, string, error)
}
