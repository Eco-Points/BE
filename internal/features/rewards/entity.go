package rewards

import "github.com/labstack/echo/v4"

type Reward struct {
	ID            uint
	Name          string
	Description   string
	PointRequired uint32
	Stock         uint32
	Image         string
}

type RHandler interface {
	AddReward() echo.HandlerFunc
}

type RServices interface {
	AddReward(newReward Reward) error
}

type RQuery interface {
	AddReward(newReward Reward) error
}
