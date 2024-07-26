package deposits

import "github.com/labstack/echo/v4"

type WasteDepositInterface struct {
	ID         uint
	TrashID    uint
	UserID     uint
	LocationID uint
	Quantity   uint
	Point      uint
	Status     string
}

type ListWasteDepositInterface []struct {
	Type     string
	Point    uint
	DepoTime string
	Quantity uint
	ID       uint
}

type QueryInterface interface {
	SetDbSchema(schema string)
	DepositTrash(data WasteDepositInterface) error
	GetUserDeposit(id uint, limit uint, offset uint) (ListWasteDepositInterface, error)
}

type ServiceInterface interface {
	DepositTrash(data WasteDepositInterface) error
	GetUserDeposit(id uint, limit uint, offset uint) (ListWasteDepositInterface, error)
}

type HandlerInterface interface {
	DepositTrash() echo.HandlerFunc
	GetUserDeposit() echo.HandlerFunc
}
