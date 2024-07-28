package service

import (
	"eco_points/internal/features/rewards"
	"errors"
)

type RewardServices struct {
	qry rewards.RServices
}

func NewRewardService(q rewards.RServices) rewards.RServices {
	return &RewardServices{
		qry: q,
	}
}

func (rs *RewardServices) AddReward(newReward rewards.Reward) error {
	err := rs.qry.AddReward(newReward)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambahkan reward")
	}
	return nil
}

func (rs *RewardServices) UpdateReward(rewardID uint, updatedReward rewards.Reward) error {
	err := rs.qry.UpdateReward(rewardID, updatedReward)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat memperbarui reward")
	}
	return nil
}

func (rs *RewardServices) DeleteReward(rewardID uint) error {
	err := rs.qry.DeleteReward(rewardID)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menghapus reward")
	}
	return nil
}
