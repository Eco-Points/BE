package service

import (
	"eco_points/internal/features/rewards"
	"eco_points/internal/utils"
	"errors"
	"mime/multipart"
)

type RewardServices struct {
	qry        rewards.RQuery
	cloudinary utils.CloudinaryUtilityInterface
}

func NewRewardService(q rewards.RQuery, c utils.CloudinaryUtilityInterface) rewards.RServices {
	return &RewardServices{
		qry:        q,
		cloudinary: c,
	}
}

func (rs *RewardServices) AddReward(newReward rewards.Reward, src multipart.File, filename string) error {
	// Upload image to Cloudinary
	imageURL, err := rs.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		return errors.New("gagal mengunggah gambar ke cloudinary")
	}
	newReward.Image = imageURL

	err = rs.qry.AddReward(newReward)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambahkan reward")
	}
	return nil
}

func (rs *RewardServices) UpdateReward(rewardID uint, updatedReward rewards.Reward, src multipart.File, filename string) error {
	if src != nil {
		// Upload image to Cloudinary
		imageURL, err := rs.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			return errors.New("gagal mengunggah gambar ke cloudinary")
		}
		updatedReward.Image = imageURL
	}

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

func (rs *RewardServices) GetRewardByID(rewardID uint) (rewards.Reward, error) {
	reward, err := rs.qry.GetRewardByID(rewardID)
	if err != nil {
		return rewards.Reward{}, errors.New("terjadi kesalahan pada server saat mengambil data reward")
	}
	return reward, nil
}

func (rs *RewardServices) GetAllRewards(limit int, offset int) ([]rewards.Reward, int, error) {
	rewards, totalItems, err := rs.qry.GetAllRewards(limit, offset)
	if err != nil {
		return nil, 0, errors.New("terjadi kesalahan pada server saat mengambil semua reward")
	}
	return rewards, totalItems, nil
}
