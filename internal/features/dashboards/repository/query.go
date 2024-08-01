package repository

import (
	"eco_points/config"
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
	var table string = config.ImportSetting().Schema
	table = table + ".users"
	err := um.db.Table(table).Where("is_admin = ?", "false").Count(&result).Error

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (um *DashboardQuery) GetDepositCount() (int, error) {
	var result int64
	var table string = config.ImportSetting().Schema
	table = table + ".waste_deposits"
	err := um.db.Table(table).Where("status = ?", "verified").Count(&result).Error

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (um *DashboardQuery) GetExchangeCount() (int, error) {
	var result int64
	var table string = config.ImportSetting().Schema
	table = table + ".exchanges"
	err := um.db.Table(table).Count(&result).Error

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

func (um *DashboardQuery) GetUser(targetID uint) (dashboards.User, error) {
	var result User
	err := um.db.First(&result, targetID).Error

	if err != nil {
		return dashboards.User{}, err
	}

	return result.ToUserEntity(), nil
}

func (um *DashboardQuery) UpdateUserStatus(targetID uint, status string) error {

	qry := um.db.Model(&User{}).Where("id = ?", targetID).Update("status", status)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (um *DashboardQuery) GetRewardStat() (int, error) {
	var result int64
	var table string = config.ImportSetting().Schema
	table = table + ".exchanges"
	err := um.db.Table(table).Count(&result).Error

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (um *DashboardQuery) GetDepositStat(whereParam string) ([]dashboards.StatData, error) {
	var result []StatData
	var result2 []dashboards.StatData
	var table string = config.ImportSetting().Schema + ".waste_deposits"

	err := um.db.Table(table).Select("date(created_at) as date , count(*) as total").Group("date(created_at)").Where(whereParam).Find(&result).Error

	if err != nil {
		return []dashboards.StatData{}, err
	}
	for _, v := range result {
		result2 = append(result2, v.ToStatDataEntity())
	}
	return result2, nil
}

func (um *DashboardQuery) GetRewardStatData(whereParam string) ([]dashboards.RewardStatData, error) {
	var result []RewardStatData
	var result2 []dashboards.RewardStatData
	var table string = config.ImportSetting().Schema + ".exchanges"

	err := um.db.Debug().Table(table).Select("reward_id , count(*) as total").Group("reward_id").Where(whereParam).Find(&result).Error

	if err != nil {
		return []dashboards.RewardStatData{}, err
	}
	for _, v := range result {
		result2 = append(result2, v.ToRewardStatDataEntity())
	}
	return result2, nil
}

func (um *DashboardQuery) GetRewardNameByID(rewardID uint) (string, error) {
	var result Reward
	var table string = config.ImportSetting().Schema + ".rewards"
	err := um.db.Table(table).First(&result, rewardID).Error

	if err != nil {
		return "", err
	}

	return result.Name, nil
}
