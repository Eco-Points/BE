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

func (rm *RewardModel) GetAllRewards(limit int, offset int) ([]rewards.Reward, int, error) {
	var rewardsList []Reward
	var totalItems int64

	// Hitung total item termasuk yang sudah soft deleted
	err := rm.db.Model(&Reward{}).Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	// Ambil data reward dengan batas limit dan offset, tidak termasuk yang sudah soft deleted
	err = rm.db.Limit(limit).Offset(offset).Find(&rewardsList).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi data reward ke entity
	rewardsEntities := make([]rewards.Reward, len(rewardsList))
	for i, reward := range rewardsList {
		rewardsEntities[i] = reward.ToRewardEntity()
	}

	return rewardsEntities, int(totalItems), nil
}
