package repository

import (
	"eco_points/internal/features/rewards"
	"eco_points/internal/features/users"

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
	cnvData := ToRewardData(newReward)
	return rm.db.Create(&cnvData).Error
}

func (rm *RewardModel) UpdateReward(rewardID uint, updateReward rewards.Reward) error {
	// Konversi data update ke struct yang sesuai
	cnvData := ToRewardData(updateReward)

	// Update data reward
	qry := rm.db.Where("id = ?", rewardID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (rm *RewardModel) DeleteReward(rewardID uint) error {
	qry := rm.db.Where("id = ?", rewardID).Delete(&Reward{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (rm *RewardModel) GetRewardByID(rewardID uint) (rewards.Reward, error) {
	var reward rewards.Reward
	err := rm.db.First(&reward, rewardID).Error
	if err != nil {
		return rewards.Reward{}, err
	}
	return reward, nil
}

func (rm *RewardModel) GetAllRewards(limit int, offset int, search string) ([]rewards.Reward, int, error) {
	var rewardsList []Reward
	var totalItems int64

	// Hitung total item termasuk yang sudah soft deleted
	err := rm.db.Model(&Reward{}).Where("name LIKE ?", "%"+search+"%").Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	// Ambil data reward dengan batas limit dan offset, tidak termasuk yang sudah soft deleted
	err = rm.db.Where("name LIKE ?", "%"+search+"%").Limit(limit).Offset(offset).Find(&rewardsList).Error
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

func (rm *RewardModel) IsAdmin(userID uint) (bool, error) {
	var user users.User
	if err := rm.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}
