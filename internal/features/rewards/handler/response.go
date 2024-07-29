package handler

import "eco_points/internal/features/rewards"

type RewardResponse struct {
	ID            uint   `json:"reward_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Stock         uint32 `json:"stock"`
	PointRequired uint32 `json:"point_required"`
	Image         string `json:"image"`
}

func ToRewardResponse(r rewards.Reward) RewardResponse {
	return RewardResponse{
		ID:            r.ID,
		Name:          r.Name,
		Description:   r.Description,
		PointRequired: r.PointRequired,
		Stock:         r.Stock,
		Image:         r.Image,
	}
}

func ToRewardsResponse(rewards []rewards.Reward) []RewardResponse {
	responses := make([]RewardResponse, len(rewards))
	for i, reward := range rewards {
		responses[i] = ToRewardResponse(reward)
	}
	return responses
}
