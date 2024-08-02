package exchanges

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Exchange struct {
	ID        uint `gorm:"primaryKey"`
	RewardID  uint
	UserID    uint
	PointUsed uint32
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

type ListExchangeInterface struct {
	ID           uint
	PointUsed    uint
	Fullname     string
	Reward       string
	ExchangeTime string
}
type ExHandler interface {
	AddExchange() echo.HandlerFunc
	GetExchangeHistory() echo.HandlerFunc
}

type ExServices interface {
	AddExchange(newExchange Exchange) error
	GetExchangeHistory(userid uint, isAdmin bool, limit uint) ([]ListExchangeInterface, error)
}

type ExQuery interface {
	AddExchange(newExchange Exchange) error
	GetExchangeHistory(userid uint, isAdmin bool, limit uint) ([]ListExchangeInterface, error)
	SetDbSchema(schema string)
}
