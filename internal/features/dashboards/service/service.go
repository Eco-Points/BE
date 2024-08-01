package service

import (
	"eco_points/internal/features/dashboards"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type DashboardServices struct {
	qry dashboards.DshQuery
}

func NewDashboardService(q dashboards.DshQuery) dashboards.DshService {
	return &DashboardServices{
		qry: q,
	}
}

func (ds *DashboardServices) GetDashboard(userID uint) (dashboards.Dashboard, error) {
	var result dashboards.Dashboard

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {

			return dashboards.Dashboard{}, errors.New("not allowed")
		}
		return dashboards.Dashboard{}, errors.New("query error")
	}

	if !isAdmin {
		return dashboards.Dashboard{}, errors.New("not allowed")
	}

	result.UserCount, err = ds.qry.GetUserCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New("an unexpected error occurred")
		}
	}

	result.DepositCount, err = ds.qry.GetDepositCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New("an unexpected error occurred")
		}
	}

	result.ExchangeCount, err = ds.qry.GetExchangeCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New("an unexpected error occurred")
		}
	}

	return result, nil
}

func (ds *DashboardServices) GetAllUsers(userID uint, nameParams string) ([]dashboards.User, error) {

	var err error
	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {

			return []dashboards.User{}, errors.New("not allowed")
		}
		return []dashboards.User{}, errors.New("query error")
	}

	if !isAdmin {
		return []dashboards.User{}, errors.New("not allowed")
	}

	nameParams = "%" + nameParams + "%"
	result, err := ds.qry.GetAllUsers(nameParams)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []dashboards.User{}, errors.New("not found")
		}
		return []dashboards.User{}, errors.New("query error")

	}

	return result, nil
}

func (ds *DashboardServices) GetUser(userID uint, targetID uint) (dashboards.User, error) {

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return dashboards.User{}, errors.New("not allowed")
		}
		return dashboards.User{}, errors.New("check admin query error")
	}
	if !isAdmin {
		return dashboards.User{}, errors.New("not allowed")
	}

	result, err := ds.qry.GetUser(targetID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return dashboards.User{}, errors.New("not found")
		}
		return dashboards.User{}, errors.New("query error")

	}

	return result, nil
}

func (ds *DashboardServices) UpdateUserStatus(userID uint, targetID uint, status string) error {

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return errors.New("not allowed")
		}
		return errors.New("check admin query error")

	}
	if !isAdmin {
		return errors.New("not allowed")
	}

	err = ds.qry.UpdateUserStatus(targetID, status)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return errors.New("not found")
		}
		return errors.New("query error")

	}

	return nil
}

func (ds *DashboardServices) GetDepositStat(userID uint, trashParam, locParam, startDate, endDate string) ([]dashboards.StatData, error) {

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []dashboards.StatData{}, errors.New("not allowed")
		}
		return []dashboards.StatData{}, errors.New("check admin query error")
	}
	if !isAdmin {
		return []dashboards.StatData{}, errors.New("not allowed")
	}

	var whereParam string = "status = 'verified' AND "
	if trashParam != "" {
		whereParam = whereParam + "trash_id = " + trashParam + " AND "
	}
	if locParam != "" {
		whereParam = whereParam + "location_id = " + locParam + " AND "
	}

	if startDate != "" {
		start, _ := time.Parse("02-01-2006", startDate)
		end, _ := time.Parse("02-01-2006", endDate)

		whereParam = whereParam + " created_at BETWEEN '" + start.Format("2006-01-02 15:04:05") + "' AND date '" + end.Format("2006-01-02 15:04:05") + "' + interval '1 day' AND "
	}

	if whereParam != "" {
		x := len(whereParam)
		whereParam = strings.TrimSpace(whereParam[0 : x-5])
	}
	result, err := ds.qry.GetDepositStat(whereParam)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []dashboards.StatData{}, errors.New("not found")
		}
		return []dashboards.StatData{}, errors.New("query error")

	}

	return result, nil
}

func (ds *DashboardServices) GetRewardStatData(userID uint, startDate, endDate string) ([]dashboards.RewardStatData, error) {

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []dashboards.RewardStatData{}, errors.New("not allowed")
		}
		return []dashboards.RewardStatData{}, errors.New("check admin query error")
	}
	if !isAdmin {
		return []dashboards.RewardStatData{}, errors.New("not allowed")
	}

	var whereParam string = ""

	if startDate != "" {
		start, _ := time.Parse("02-01-2006", startDate)
		end, _ := time.Parse("02-01-2006", endDate)

		whereParam = whereParam + " created_at BETWEEN date '" + start.Format("2006-01-02 15:04:05") + "' - interval '1 day'AND date '" + end.Format("2006-01-02 15:04:05") + "' + INTERVAL '1 DAY' AND "
	}

	if whereParam != "" {
		x := len(whereParam)
		whereParam = strings.TrimSpace(whereParam[0 : x-5])
	}

	result, err := ds.qry.GetRewardStatData(whereParam)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []dashboards.RewardStatData{}, errors.New("not found")
		}
		return []dashboards.RewardStatData{}, errors.New("query error")

	}

	for i, v := range result {
		result[i].RewardName, err = ds.qry.GetRewardNameByID(v.RewardID)
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				return []dashboards.RewardStatData{}, errors.New("not found")
			}
			return []dashboards.RewardStatData{}, errors.New("query error")

		}

	}

	return result, nil
}
