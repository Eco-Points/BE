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

func (rm *RewardModel) UpdateReward(rewardID uint, updateReward rewards.Reward) error {
	var reward rewards.Reward
	err := rm.db.First(&reward, rewardID).Error
	if err != nil {
		return err
	}
	return rm.db.Model(&reward).Updates(updateReward).Error
}

func (rm *RewardModel) DeleteReward(rewardID uint) error {
	return rm.db.Delete(&rewards.Reward{}, rewardID).Error
}

func (rm *RewardModel) GetRewardByID(rewardID uint) (rewards.Reward, error) {
	var reward rewards.Reward
	err := rm.db.First(&reward, rewardID).Error
	if err != nil {
		return rewards.Reward{}, err
	}
	return reward, nil
}
