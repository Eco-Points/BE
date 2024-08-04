// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	exchanges "eco_points/internal/features/exchanges"

	mock "github.com/stretchr/testify/mock"
)

// ExServices is an autogenerated mock type for the ExServices type
type ExServices struct {
	mock.Mock
}

// AddExchange provides a mock function with given fields: newExchange
func (_m *ExServices) AddExchange(newExchange exchanges.Exchange) error {
	ret := _m.Called(newExchange)

	if len(ret) == 0 {
		panic("no return value specified for AddExchange")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(exchanges.Exchange) error); ok {
		r0 = rf(newExchange)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExchangeHistory provides a mock function with given fields: userid, isAdmin, limit
func (_m *ExServices) GetExchangeHistory(userid uint, isAdmin bool, limit uint) ([]exchanges.ListExchangeInterface, error) {
	ret := _m.Called(userid, isAdmin, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetExchangeHistory")
	}

	var r0 []exchanges.ListExchangeInterface
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, bool, uint) ([]exchanges.ListExchangeInterface, error)); ok {
		return rf(userid, isAdmin, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, bool, uint) []exchanges.ListExchangeInterface); ok {
		r0 = rf(userid, isAdmin, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]exchanges.ListExchangeInterface)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, bool, uint) error); ok {
		r1 = rf(userid, isAdmin, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExServices creates a new instance of ExServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExServices {
	mock := &ExServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
