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

type ExHandler interface {
	AddExchange() echo.HandlerFunc
}

type ExServices interface {
	AddExchange(newExchange Exchange) error
}

type ExQuery interface {
	AddExchange(newExchange Exchange) error
}
