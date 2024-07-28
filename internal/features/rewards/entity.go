package rewards

import (
	"github.com/labstack/echo/v4"
)

type Reward struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Description   string
	PointRequired uint32
	Stock         uint32
	Image         string
}

type RHandler interface {
	AddReward() echo.HandlerFunc
	UpdateReward() echo.HandlerFunc
	DeleteReward() echo.HandlerFunc
	GetRewardByID() echo.HandlerFunc
}

type RServices interface {
	AddReward(newReward Reward) error
	UpdateReward(rewardID uint, updatedReward Reward) error
	DeleteReward(rewardID uint) error
	GetRewardByID(rewardID uint) (Reward, error)
}

type RQuery interface {
	AddReward(newReward Reward) error
	UpdateReward(rewardID uint, updatedReward Reward) error
	DeleteReward(rewardID uint) error
	GetRewardByID(rewardID uint) (Reward, error)
}
