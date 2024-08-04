// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	io "io"
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// CloudinaryUtilityInterface is an autogenerated mock type for the CloudinaryUtilityInterface type
type CloudinaryUtilityInterface struct {
	mock.Mock
}

// FileCheck provides a mock function with given fields: file
func (_m *CloudinaryUtilityInterface) FileCheck(file *multipart.FileHeader) (multipart.File, error) {
	ret := _m.Called(file)

	if len(ret) == 0 {
		panic("no return value specified for FileCheck")
	}

	var r0 multipart.File
	var r1 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) (multipart.File, error)); ok {
		return rf(file)
	}
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) multipart.File); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multipart.File)
		}
	}

	if rf, ok := ret.Get(1).(func(*multipart.FileHeader) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileOpener provides a mock function with given fields: file
func (_m *CloudinaryUtilityInterface) FileOpener(file *multipart.FileHeader) (multipart.File, error) {
	ret := _m.Called(file)

	if len(ret) == 0 {
		panic("no return value specified for FileOpener")
	}

	var r0 multipart.File
	var r1 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) (multipart.File, error)); ok {
		return rf(file)
	}
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) multipart.File); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multipart.File)
		}
	}

	if rf, ok := ret.Get(1).(func(*multipart.FileHeader) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadToCloudinary provides a mock function with given fields: file, filename
func (_m *CloudinaryUtilityInterface) UploadToCloudinary(file io.Reader, filename string) (string, error) {
	ret := _m.Called(file, filename)

	if len(ret) == 0 {
		panic("no return value specified for UploadToCloudinary")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(io.Reader, string) (string, error)); ok {
		return rf(file, filename)
	}
	if rf, ok := ret.Get(0).(func(io.Reader, string) string); ok {
		r0 = rf(file, filename)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(io.Reader, string) error); ok {
		r1 = rf(file, filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCloudinaryUtilityInterface creates a new instance of CloudinaryUtilityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloudinaryUtilityInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloudinaryUtilityInterface {
	mock := &CloudinaryUtilityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
