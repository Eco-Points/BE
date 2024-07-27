// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	trashes "eco_points/internal/features/trashes"
)

// ServiceTrashInterface is an autogenerated mock type for the ServiceTrashInterface type
type ServiceTrashInterface struct {
	mock.Mock
}

// AddTrash provides a mock function with given fields: tData, file
func (_m *ServiceTrashInterface) AddTrash(tData trashes.TrashEntity, file *multipart.FileHeader) error {
	ret := _m.Called(tData, file)

	if len(ret) == 0 {
		panic("no return value specified for AddTrash")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(trashes.TrashEntity, *multipart.FileHeader) error); ok {
		r0 = rf(tData, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTrash provides a mock function with given fields: ttype
func (_m *ServiceTrashInterface) GetTrash(ttype string) (trashes.ListTrashEntity, error) {
	ret := _m.Called(ttype)

	if len(ret) == 0 {
		panic("no return value specified for GetTrash")
	}

	var r0 trashes.ListTrashEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (trashes.ListTrashEntity, error)); ok {
		return rf(ttype)
	}
	if rf, ok := ret.Get(0).(func(string) trashes.ListTrashEntity); ok {
		r0 = rf(ttype)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(trashes.ListTrashEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ttype)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceTrashInterface creates a new instance of ServiceTrashInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceTrashInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceTrashInterface {
	mock := &ServiceTrashInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
