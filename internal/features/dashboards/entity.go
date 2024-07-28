package dashboards

import "github.com/labstack/echo/v4"

type Dashboard struct {
	UserCount     int `json:"user_count"`
	DepositCount  int `json:"deposit_count"`
	ExchangeCount int `json:"exchange_count"`
}

type Handler interface {
	GetDashboard() echo.HandlerFunc
}

type Service interface {
	GetDashboard(userID uint) (Dashboard, error)
}

type Query interface {
	GetUserCount() (int, error)
	GetExchangeCount() (int, error)
	GetDepositCount() (int, error)
	CheckIsAdmin(userID uint) (bool, error)
}
