package service

import (
	"eco_points/internal/features/dashboards"
	"eco_points/internal/utils"
	"errors"

	"gorm.io/gorm"
)

type DashboardServices struct {
	qry dashboards.Query
}

func NewUserService(q dashboards.Query, p utils.PassUtilInterface, j utils.JwtUtilityInterface, c utils.CloudinaryUtilityInterface) dashboards.Service {
	return &DashboardServices{
		qry: q,
	}
}

func (us *DashboardServices) GetDashboard(userID uint) (dashboards.Dashboard, error) {
	var result dashboards.Dashboard
	var err error
	msg := "unexpected error occurred"

	isAdmin, err := us.qry.CheckIsAdmin(userID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() || !isAdmin {
			msg := "not found"
			return dashboards.Dashboard{}, errors.New(msg)
		} else {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	result.UserCount, err = us.qry.GetUserCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	result.DepositCount, err = us.qry.GetDepositCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	result.ExchangeCount, err = us.qry.GetExchangeCount()
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return dashboards.Dashboard{}, errors.New(msg)
		}
	}

	return result, nil
}
