package repository

import (
	"eco_points/internal/features/rewards"

	"gorm.io/gorm"
)

type RewardModel struct {
	db *gorm.DB
}

func NewRewardModel(connection *gorm.DB) rewards.RQuery {
	return &RewardModel{
		db: connection,
	}
}

func (rm *RewardModel) AddReward(newReward rewards.Reward) error {
	return rm.db.Create(&newReward).Error
}
