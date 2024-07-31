// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	dashboards "eco_points/internal/features/dashboards"

	mock "github.com/stretchr/testify/mock"
)

// DshQuery is an autogenerated mock type for the DshQuery type
type DshQuery struct {
	mock.Mock
}

// CheckIsAdmin provides a mock function with given fields: userID
func (_m *DshQuery) CheckIsAdmin(userID uint) (bool, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for CheckIsAdmin")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (bool, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) bool); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields: nameParams
func (_m *DshQuery) GetAllUsers(nameParams string) ([]dashboards.User, error) {
	ret := _m.Called(nameParams)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []dashboards.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]dashboards.User, error)); ok {
		return rf(nameParams)
	}
	if rf, ok := ret.Get(0).(func(string) []dashboards.User); ok {
		r0 = rf(nameParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dashboards.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nameParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDepositCount provides a mock function with given fields:
func (_m *DshQuery) GetDepositCount() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDepositCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExchangeCount provides a mock function with given fields:
func (_m *DshQuery) GetExchangeCount() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExchangeCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userID
func (_m *DshQuery) GetUser(userID uint) (dashboards.User, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 dashboards.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (dashboards.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) dashboards.User); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(dashboards.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserCount provides a mock function with given fields:
func (_m *DshQuery) GetUserCount() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUserCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserStatus provides a mock function with given fields: target_id, status
func (_m *DshQuery) UpdateUserStatus(target_id uint, status string) error {
	ret := _m.Called(target_id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(target_id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDshQuery creates a new instance of DshQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDshQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *DshQuery {
	mock := &DshQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
