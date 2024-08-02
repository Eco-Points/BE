// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	dashboards "eco_points/internal/features/dashboards"

	mock "github.com/stretchr/testify/mock"
)

// DshService is an autogenerated mock type for the DshService type
type DshService struct {
	mock.Mock
}

// DeleteUserByAdmin provides a mock function with given fields: userID, targetID
func (_m *DshService) DeleteUserByAdmin(userID uint, targetID uint) error {
	ret := _m.Called(userID, targetID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserByAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, targetID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields: userID, nameParams
func (_m *DshService) GetAllUsers(userID uint, nameParams string) ([]dashboards.User, error) {
	ret := _m.Called(userID, nameParams)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []dashboards.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) ([]dashboards.User, error)); ok {
		return rf(userID, nameParams)
	}
	if rf, ok := ret.Get(0).(func(uint, string) []dashboards.User); ok {
		r0 = rf(userID, nameParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dashboards.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userID, nameParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboard provides a mock function with given fields: userID
func (_m *DshService) GetDashboard(userID uint) (dashboards.Dashboard, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetDashboard")
	}

	var r0 dashboards.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (dashboards.Dashboard, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) dashboards.Dashboard); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(dashboards.Dashboard)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDepositStat provides a mock function with given fields: userID, trashParam, locParam, startDate, endDate
func (_m *DshService) GetDepositStat(userID uint, trashParam string, locParam string, startDate string, endDate string) ([]dashboards.StatData, error) {
	ret := _m.Called(userID, trashParam, locParam, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for GetDepositStat")
	}

	var r0 []dashboards.StatData
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, string, string, string) ([]dashboards.StatData, error)); ok {
		return rf(userID, trashParam, locParam, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(uint, string, string, string, string) []dashboards.StatData); ok {
		r0 = rf(userID, trashParam, locParam, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dashboards.StatData)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, string, string, string) error); ok {
		r1 = rf(userID, trashParam, locParam, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRewardStatData provides a mock function with given fields: userID, startDate, endDate
func (_m *DshService) GetRewardStatData(userID uint, startDate string, endDate string) ([]dashboards.RewardStatData, error) {
	ret := _m.Called(userID, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for GetRewardStatData")
	}

	var r0 []dashboards.RewardStatData
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, string) ([]dashboards.RewardStatData, error)); ok {
		return rf(userID, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(uint, string, string) []dashboards.RewardStatData); ok {
		r0 = rf(userID, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dashboards.RewardStatData)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, string) error); ok {
		r1 = rf(userID, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userID, targetID
func (_m *DshService) GetUser(userID uint, targetID uint) (dashboards.User, error) {
	ret := _m.Called(userID, targetID)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 dashboards.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (dashboards.User, error)); ok {
		return rf(userID, targetID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) dashboards.User); ok {
		r0 = rf(userID, targetID)
	} else {
		r0 = ret.Get(0).(dashboards.User)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, targetID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserStatus provides a mock function with given fields: userID, targetID, status
func (_m *DshService) UpdateUserStatus(userID uint, targetID uint, status string) error {
	ret := _m.Called(userID, targetID, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(userID, targetID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDshService creates a new instance of DshService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDshService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DshService {
	mock := &DshService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
