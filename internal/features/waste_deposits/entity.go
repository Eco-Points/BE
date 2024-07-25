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

type QueryInterface interface {
	DepositTrash(data WasteDepositInterface) error
}

type ServiceInterface interface {
	DepositTrash(data WasteDepositInterface) error
}

type HandlerInterface interface {
	DepositTrash() echo.HandlerFunc
}
