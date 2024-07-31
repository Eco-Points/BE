package locations

import "github.com/labstack/echo/v4"

type LocationInterface struct {
	Address    string
	Long       string
	Lat        string
	Status     string
	Start_time string
	End_time   string
	Phone      string
	UserID     uint
	Name       string
	ID         uint
}

type QueryLocInterface interface {
	AddLocation(data LocationInterface) error
	GetLocation(limit uint) ([]LocationInterface, error)
}

type ServiceLocInterface interface {
	AddLocation(data LocationInterface) error
	GetLocation(limit uint) ([]LocationInterface, error)
}

type HandlerLocInterface interface {
	AddLocation() echo.HandlerFunc
	GetLocation() echo.HandlerFunc
}
