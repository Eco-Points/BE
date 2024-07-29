package repository

import (
	"eco_points/internal/features/dashboards"

	"gorm.io/gorm"
)

type DashboardQuery struct {
	db *gorm.DB
}

func NewDashboardQuery(connect *gorm.DB) dashboards.DshQuery {
	return &DashboardQuery{
		db: connect,
	}
}

func (um *DashboardQuery) GetUserCount() (int, error) {
	var result int64
	err := um.db.Table("users").Where("is_admin = ?", "false").Count(&result).Error

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (um *DashboardQuery) GetDepositCount() (int, error) {
	var result int64
	err := um.db.Table("waste_deposits").Where("status = ?", "verified").Count(&result).Error

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (um *DashboardQuery) GetExchangeCount() (int, error) {
	var result int64
	err := um.db.Table("user_reward_exchanges").Count(&result).Error

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (um *DashboardQuery) CheckIsAdmin(userID uint) (bool, error) {
	var result User
	err := um.db.First(&result, userID).Error

	if err != nil {
		return false, err
	}

	return result.IsAdmin, nil
}

func (um *DashboardQuery) GetAllUsers(nameParams string) ([]dashboards.User, error) {
	var result []User
	var result2 []dashboards.User
	err := um.db.Where("lower(fullname) LIKE lower(?) AND is_admin = ?", nameParams, "false").Find(&result).Error

	if err != nil {
		return []dashboards.User{}, err
	}

	for _, v := range result {
		result2 = append(result2, v.ToUserEntity())
	}
	return result2, nil
}
