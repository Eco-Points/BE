// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	users "eco_points/internal/features/users"

	mock "github.com/stretchr/testify/mock"
)

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ID
func (_m *Query) DeleteUser(ID uint) error {
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
func (_m *Query) GetUser(ID uint) (users.User, error) {
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

// Login provides a mock function with given fields: email
func (_m *Query) Login(email string) (users.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (users.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) users.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUsers
func (_m *Query) Register(newUsers users.User) error {
	ret := _m.Called(newUsers)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(users.User) error); ok {
		r0 = rf(newUsers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ID, updateUser
func (_m *Query) UpdateUser(ID uint, updateUser users.User) error {
	ret := _m.Called(ID, updateUser)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.User) error); ok {
		r0 = rf(ID, updateUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuery creates a new instance of Query. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Query {
	mock := &Query{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
