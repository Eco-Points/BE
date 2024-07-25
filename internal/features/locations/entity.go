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

type QueryInterface interface {
	AddLocation(data LocationInterface) error
}

type ServiceInterface interface {
	AddLocation(data LocationInterface) error
}

type HandlerInterface interface {
	AddLocation() echo.HandlerFunc
}
