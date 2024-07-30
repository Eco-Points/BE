package service

import (
	"eco_points/internal/features/dashboards"
	"errors"

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
