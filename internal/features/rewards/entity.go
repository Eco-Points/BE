package rewards

import (
	"mime/multipart"

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
	GetAllRewards() echo.HandlerFunc
}

type RServices interface {
	AddReward(userID uint, newReward Reward, src multipart.File, filename string) error
	UpdateReward(userID uint, rewardID uint, updatedReward Reward, src multipart.File, filename string) error
	DeleteReward(userID uint, rewardID uint) error
	GetRewardByID(rewardID uint) (Reward, error)
	GetAllRewards(limit int, offset int, search string) ([]Reward, int, error)
}

type RQuery interface {
	AddReward(newReward Reward) error
	UpdateReward(rewardID uint, updatedReward Reward) error
	DeleteReward(rewardID uint) error
	GetRewardByID(rewardID uint) (Reward, error)
	GetAllRewards(limit int, offset int, search string) ([]Reward, int, error)
	IsAdmin(userID uint) (bool, error)
}
