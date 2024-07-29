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
}

type QueryLocInterface interface {
	AddLocation(data LocationInterface) error
	GetLocation(id uint) ([]LocationInterface, error)
	GetAllLocation() ([]LocationInterface, error)
}

type ServiceLocInterface interface {
	AddLocation(data LocationInterface) error
	GetLocation(id uint) ([]LocationInterface, error)
}

type HandlerLocInterface interface {
	AddLocation() echo.HandlerFunc
	GetLocation() echo.HandlerFunc
}
