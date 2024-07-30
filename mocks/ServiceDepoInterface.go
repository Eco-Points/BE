// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	deposits "eco_points/internal/features/waste_deposits"

	mock "github.com/stretchr/testify/mock"
)

// ServiceDepoInterface is an autogenerated mock type for the ServiceDepoInterface type
type ServiceDepoInterface struct {
	mock.Mock
}

// DepositTrash provides a mock function with given fields: data
func (_m *ServiceDepoInterface) DepositTrash(data deposits.WasteDepositInterface) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for DepositTrash")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(deposits.WasteDepositInterface) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDepositbyId provides a mock function with given fields: deposit_id
func (_m *ServiceDepoInterface) GetDepositbyId(deposit_id uint) (deposits.WasteDepositInterface, error) {
	ret := _m.Called(deposit_id)

	if len(ret) == 0 {
		panic("no return value specified for GetDepositbyId")
	}

	var r0 deposits.WasteDepositInterface
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (deposits.WasteDepositInterface, error)); ok {
		return rf(deposit_id)
	}
	if rf, ok := ret.Get(0).(func(uint) deposits.WasteDepositInterface); ok {
		r0 = rf(deposit_id)
	} else {
		r0 = ret.Get(0).(deposits.WasteDepositInterface)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(deposit_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDeposit provides a mock function with given fields: id, limit, offset, is_admin
func (_m *ServiceDepoInterface) GetUserDeposit(id uint, limit uint, offset uint, is_admin bool) (deposits.ListWasteDepositInterface, error) {
	ret := _m.Called(id, limit, offset, is_admin)

	if len(ret) == 0 {
		panic("no return value specified for GetUserDeposit")
	}

	var r0 deposits.ListWasteDepositInterface
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint, bool) (deposits.ListWasteDepositInterface, error)); ok {
		return rf(id, limit, offset, is_admin)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, uint, bool) deposits.ListWasteDepositInterface); ok {
		r0 = rf(id, limit, offset, is_admin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(deposits.ListWasteDepositInterface)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, uint, bool) error); ok {
		r1 = rf(id, limit, offset, is_admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWasteDepositStatus provides a mock function with given fields: waste_id, status
func (_m *ServiceDepoInterface) UpdateWasteDepositStatus(waste_id uint, status string) error {
	ret := _m.Called(waste_id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWasteDepositStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(waste_id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServiceDepoInterface creates a new instance of ServiceDepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceDepoInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceDepoInterface {
	mock := &ServiceDepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
