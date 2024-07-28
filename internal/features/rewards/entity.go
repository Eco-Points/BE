package rewards

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Reward struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Description   string
	PointRequired uint32
	Stock         uint32
	Image         string
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type RHandler interface {
	AddReward() echo.HandlerFunc
	UpdateReward() echo.HandlerFunc
	DeleteReward() echo.HandlerFunc
}

type RServices interface {
	AddReward(newReward Reward) error
	UpdateReward(rewardID uint, updatedReward Reward) error
	DeleteReward(rewardID uint) error
}

type RQuery interface {
	AddReward(newReward Reward) error
	UpdateReward(rewardID uint, updatedReward Reward) error
	DeleteReward(rewardID uint) error
}
