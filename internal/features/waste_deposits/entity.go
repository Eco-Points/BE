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
	Type       string
	DepoTime   string
}

type ListWasteDepositInterface []struct {
	Type     string
	Point    uint
	DepoTime string
	Quantity uint
	ID       uint
	Status   string
}

type QueryInterface interface {
	SetDbSchema(schema string)
	DepositTrash(data WasteDepositInterface) error
	UpdateWasteDepositStatus(waste_id uint, status string) error
	GetUserDeposit(id uint, limit uint, offset uint) (ListWasteDepositInterface, error)
	GetDepositbyId(deposit_id uint) (WasteDepositInterface, error)
}

type ServiceInterface interface {
	DepositTrash(data WasteDepositInterface) error
	UpdateWasteDepositStatus(waste_id uint, status string) error
	GetUserDeposit(id uint, limit uint, offset uint) (ListWasteDepositInterface, error)
	GetDepositbyId(deposit_id uint) (WasteDepositInterface, error)
}

type HandlerInterface interface {
	DepositTrash() echo.HandlerFunc
	UpdateWasteDepositStatus() echo.HandlerFunc
	GetUserDeposit() echo.HandlerFunc
	GetDepositbyId() echo.HandlerFunc
}
