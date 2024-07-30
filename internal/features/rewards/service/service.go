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

func (rs *RewardServices) AddReward(userID uint, newReward rewards.Reward, src multipart.File, filename string) error {
	isAdmin, err := rs.qry.IsAdmin(userID)
	if err != nil || !isAdmin {
		return errors.New("forbidden: user is not an admin")
	}

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

func (rs *RewardServices) UpdateReward(userID uint, rewardID uint, updatedReward rewards.Reward, src multipart.File, filename string) error {
	isAdmin, err := rs.qry.IsAdmin(userID)
	if err != nil || !isAdmin {
		return errors.New("forbidden: user is not an admin")
	}

	if src != nil {
		// Upload image to Cloudinary
		imageURL, err := rs.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			return errors.New("failed to upload image to Cloudinary")
		}
		updatedReward.Image = imageURL
	}

	err = rs.qry.UpdateReward(rewardID, updatedReward)
	if err != nil {
		return errors.New("server error while updating reward")
	}
	return nil
}

func (rs *RewardServices) DeleteReward(userID uint, rewardID uint) error {
	isAdmin, err := rs.qry.IsAdmin(userID)
	if err != nil || !isAdmin {
		return errors.New("forbidden: user is not an admin")
	}

	err = rs.qry.DeleteReward(rewardID)
	if err != nil {
		return errors.New("server error while deleting reward")
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

func (rs *RewardServices) GetAllRewards(limit int, offset int, search string) ([]rewards.Reward, int, error) {
	rewards, totalItems, err := rs.qry.GetAllRewards(limit, offset, search)
	if err != nil {
		return nil, 0, errors.New("terjadi kesalahan pada server saat mengambil semua reward")
	}
	return rewards, totalItems, nil
}
