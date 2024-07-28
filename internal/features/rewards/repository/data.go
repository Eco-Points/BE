package repository

import (
	"eco_points/internal/features/rewards"

	"gorm.io/gorm"
)

type Reward struct {
	gorm.Model
	Name          string
	Description   string
	PointRequired uint32
	Stock         uint32
	Image         string
}

func (r *Reward) ToRewardEntity() rewards.Reward {
	return rewards.Reward{
		ID:            r.ID,
		Name:          r.Name,
		Description:   r.Description,
		PointRequired: r.PointRequired,
		Stock:         r.Stock,
		Image:         r.Image,
	}
}

func ToRewardData(input rewards.Reward) Reward {
	return Reward{
		Name:          input.Name,
		Description:   input.Description,
		PointRequired: input.PointRequired,
		Stock:         input.Stock,
		Image:         input.Image,
	}
}
