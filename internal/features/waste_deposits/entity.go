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
	Fullname   string
}

type ListWasteDepositInterface []struct {
	Type     string
	Point    uint
	DepoTime string
	Quantity uint
	ID       uint
	Status   string
	Fullname string
}

type QueryDepoInterface interface {
	SetDbSchema(schema string)
	DepositTrash(data WasteDepositInterface) error
	UpdateWasteDepositStatus(waste_id uint, status string) (uint, int, error)
	GetUserDeposit(id uint, limit uint, offset uint, is_admin bool) (ListWasteDepositInterface, error)
	GetDepositbyId(deposit_id uint) (WasteDepositInterface, error)
	GetUserEmailData(ID uint) (string, string, error)
}

type ServiceDepoInterface interface {
	DepositTrash(data WasteDepositInterface) error
	UpdateWasteDepositStatus(waste_id uint, status string) error
	GetUserDeposit(id uint, limit uint, offset uint, is_admin bool) (ListWasteDepositInterface, error)
	GetDepositbyId(deposit_id uint) (WasteDepositInterface, error)
}

type HandlerDepoInterface interface {
	DepositTrash() echo.HandlerFunc
	UpdateWasteDepositStatus() echo.HandlerFunc
	GetUserDeposit() echo.HandlerFunc
	GetDepositbyId() echo.HandlerFunc
}
