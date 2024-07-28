package handler

import "eco_points/internal/features/rewards"

type CreateOrUpdateRewardRequest struct {
	Name          string `form:"name"`
	Description   string `form:"description"`
	Stock         uint32 `form:"stock"`
	PointRequired uint32 `form:"point_required"`
	Image         string `form:"image"`
}

func ToRewardModel(req CreateOrUpdateRewardRequest, imageURL string) rewards.Reward {
	return rewards.Reward{
		Name:          req.Name,
		Description:   req.Description,
		PointRequired: req.PointRequired,
		Stock:         req.Stock,
		Image:         imageURL,
	}
}
