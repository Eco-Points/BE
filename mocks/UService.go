// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	users "eco_points/internal/features/users"
)

// UService is an autogenerated mock type for the UService type
type UService struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ID
func (_m *UService) DeleteUser(ID uint) error {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ID
func (_m *UService) GetUser(ID uint) (users.User, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (users.User, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uint) users.User); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *UService) Login(email string, password string) (users.User, string, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 users.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (users.User, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) users.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: newUser
func (_m *UService) Register(newUser users.User) error {
	ret := _m.Called(newUser)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(users.User) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ID, updateUser, file
func (_m *UService) UpdateUser(ID uint, updateUser users.User, file *multipart.FileHeader) error {
	ret := _m.Called(ID, updateUser, file)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.User, *multipart.FileHeader) error); ok {
		r0 = rf(ID, updateUser, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUService creates a new instance of UService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UService {
	mock := &UService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}