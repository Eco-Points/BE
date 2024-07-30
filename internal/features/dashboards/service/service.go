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
	var err error
	msg := "an unexpected error occurred"

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() || !isAdmin {
			msg := "not found"
			return dashboards.Dashboard{}, errors.New(msg)
		} else {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	result.UserCount, err = ds.qry.GetUserCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	result.DepositCount, err = ds.qry.GetDepositCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	result.ExchangeCount, err = ds.qry.GetExchangeCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	return result, nil
}

func (ds *DashboardServices) GetAllUsers(userID uint, nameParams string) ([]dashboards.User, error) {

	var err error
	msg := "an unexpected error occurred"

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() || !isAdmin {
			msg := "not allowed"
			return []dashboards.User{}, errors.New(msg)
		} else {
			return []dashboards.User{}, errors.New(msg)
		}
	}
	nameParams = "%" + nameParams + "%"
	result, err := ds.qry.GetAllUsers(nameParams)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg := "not found"
			return []dashboards.User{}, errors.New(msg)
		} else {
			return []dashboards.User{}, errors.New(msg)
		}
	}

	return result, nil
}

func (ds *DashboardServices) GetUser(userID uint, targetID uint) (dashboards.User, error) {

	isAdmin, err := ds.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() || !isAdmin {
			return dashboards.User{}, errors.New("not allowed")
		}
		return dashboards.User{}, errors.New("check admin query error")

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
		if err.Error() == gorm.ErrRecordNotFound.Error() || !isAdmin {
			return errors.New("not allowed")
		}
		return errors.New("check admin query error")

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
