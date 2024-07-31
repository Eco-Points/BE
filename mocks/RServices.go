// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	rewards "eco_points/internal/features/rewards"
)

// RServices is an autogenerated mock type for the RServices type
type RServices struct {
	mock.Mock
}

// AddReward provides a mock function with given fields: newReward, src, filename
func (_m *RServices) AddReward(newReward rewards.Reward, src multipart.File, filename string) error {
	ret := _m.Called(newReward, src, filename)

	if len(ret) == 0 {
		panic("no return value specified for AddReward")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(rewards.Reward, multipart.File, string) error); ok {
		r0 = rf(newReward, src, filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReward provides a mock function with given fields: rewardID
func (_m *RServices) DeleteReward(rewardID uint) error {
	ret := _m.Called(rewardID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReward")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(rewardID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllRewards provides a mock function with given fields: limit, offset
func (_m *RServices) GetAllRewards(limit int, offset int) ([]rewards.Reward, int, error) {
	ret := _m.Called(limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllRewards")
	}

	var r0 []rewards.Reward
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]rewards.Reward, int, error)); ok {
		return rf(limit, offset)
	}
	if rf, ok := ret.Get(0).(func(int, int) []rewards.Reward); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rewards.Reward)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRewardByID provides a mock function with given fields: rewardID
func (_m *RServices) GetRewardByID(rewardID uint) (rewards.Reward, error) {
	ret := _m.Called(rewardID)

	if len(ret) == 0 {
		panic("no return value specified for GetRewardByID")
	}

	var r0 rewards.Reward
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (rewards.Reward, error)); ok {
		return rf(rewardID)
	}
	if rf, ok := ret.Get(0).(func(uint) rewards.Reward); ok {
		r0 = rf(rewardID)
	} else {
		r0 = ret.Get(0).(rewards.Reward)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(rewardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReward provides a mock function with given fields: rewardID, updatedReward, src, filename
func (_m *RServices) UpdateReward(rewardID uint, updatedReward rewards.Reward, src multipart.File, filename string) error {
	ret := _m.Called(rewardID, updatedReward, src, filename)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReward")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, rewards.Reward, multipart.File, string) error); ok {
		r0 = rf(rewardID, updatedReward, src, filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRServices creates a new instance of RServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *RServices {
	mock := &RServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}