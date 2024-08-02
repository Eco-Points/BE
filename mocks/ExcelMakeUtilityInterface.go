// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ExcelMakeUtilityInterface is an autogenerated mock type for the ExcelMakeUtilityInterface type
type ExcelMakeUtilityInterface struct {
	mock.Mock
}

// MakeExcel provides a mock function with given fields: name, date, data
func (_m *ExcelMakeUtilityInterface) MakeExcel(name string, date string, data interface{}) (string, error) {
	ret := _m.Called(name, date, data)

	if len(ret) == 0 {
		panic("no return value specified for MakeExcel")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, interface{}) (string, error)); ok {
		return rf(name, date, data)
	}
	if rf, ok := ret.Get(0).(func(string, string, interface{}) string); ok {
		r0 = rf(name, date, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, interface{}) error); ok {
		r1 = rf(name, date, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExcelMakeUtilityInterface creates a new instance of ExcelMakeUtilityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExcelMakeUtilityInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExcelMakeUtilityInterface {
	mock := &ExcelMakeUtilityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
