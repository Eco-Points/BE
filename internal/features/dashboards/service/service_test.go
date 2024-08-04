package service_test

import (
	"eco_points/internal/features/dashboards"
	"eco_points/internal/features/dashboards/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDashboard(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	t.Run("success get dashboard", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUserCount").Return(10, nil).Once()
		mockQuery.On("GetDepositCount").Return(20, nil).Once()
		mockQuery.On("GetExchangeCount").Return(30, nil).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.Nil(t, err)
		assert.Equal(t, dashboards.Dashboard{
			UserCount:     10,
			DepositCount:  20,
			ExchangeCount: 30,
		}, result)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
		assert.Equal(t, dashboards.Dashboard{}, result)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUserCount").Return(0, errors.New("query error")).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.NotNil(t, err)
		assert.Equal(t, "an unexpected error occurred", err.Error())
		assert.Equal(t, dashboards.Dashboard{}, result)
	})

	t.Run("failure - CheckIsAdmin query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, dashboards.Dashboard{}, result)
	})

	t.Run("failure - GetUserCount query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUserCount").Return(0, errors.New("query error")).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.NotNil(t, err)
		assert.Equal(t, "an unexpected error occurred", err.Error())
		assert.Equal(t, dashboards.Dashboard{}, result)
	})

	t.Run("failure - GetDepositCount query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUserCount").Return(10, nil).Once()
		mockQuery.On("GetDepositCount").Return(0, errors.New("query error")).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.NotNil(t, err)
		assert.Equal(t, "an unexpected error occurred", err.Error())
		assert.Equal(t, dashboards.Dashboard{}, result)
	})

	t.Run("failure - GetExchangeCount query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUserCount").Return(10, nil).Once()
		mockQuery.On("GetDepositCount").Return(20, nil).Once()
		mockQuery.On("GetExchangeCount").Return(0, errors.New("query error")).Once()

		result, err := dashboardService.GetDashboard(1)
		assert.NotNil(t, err)
		assert.Equal(t, "an unexpected error occurred", err.Error())
		assert.Equal(t, dashboards.Dashboard{}, result)
	})

}

func TestGetAllUsers(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	users := []dashboards.User{
		{ID: 1, Fullname: "John Doe"},
		{ID: 2, Fullname: "Jane Doe"},
	}

	t.Run("success get all users", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetAllUsers", "%John%").Return(users, nil).Once()

		result, err := dashboardService.GetAllUsers(1, "John")
		assert.Nil(t, err)
		assert.Equal(t, users, result)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		result, err := dashboardService.GetAllUsers(1, "John")
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
		assert.Equal(t, []dashboards.User{}, result)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetAllUsers", "%John%").Return(nil, errors.New("query error")).Once()

		result, err := dashboardService.GetAllUsers(1, "John")
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, []dashboards.User{}, result)
	})

	t.Run("failure - CheckIsAdmin query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()

		result, err := dashboardService.GetAllUsers(1, "John")
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, []dashboards.User{}, result)
	})

	t.Run("failure - GetAllUsers query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetAllUsers", "%John%").Return(nil, errors.New("query error")).Once()

		result, err := dashboardService.GetAllUsers(1, "John")
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, []dashboards.User{}, result)
	})

}

func TestGetUser(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	user := dashboards.User{ID: 1, Fullname: "John Doe"}

	t.Run("success get user", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUser", uint(2)).Return(user, nil).Once()

		result, err := dashboardService.GetUser(1, 2)
		assert.Nil(t, err)
		assert.Equal(t, user, result)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		result, err := dashboardService.GetUser(1, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
		assert.Equal(t, dashboards.User{}, result)
	})

	t.Run("failure - CheckIsAdmin query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, errors.New("check admin query error")).Once()

		result, err := dashboardService.GetUser(1, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "check admin query error", err.Error())
		assert.Equal(t, dashboards.User{}, result)
	})

	t.Run("failure - GetUser query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUser", uint(2)).Return(dashboards.User{}, errors.New("query error")).Once()

		result, err := dashboardService.GetUser(1, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, dashboards.User{}, result)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetUser", uint(2)).Return(dashboards.User{}, errors.New("record not found")).Once()

		result, err := dashboardService.GetUser(1, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "not found", err.Error())
		assert.Equal(t, dashboards.User{}, result)
	})
}

func TestUpdateUserStatus(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	t.Run("success update user status", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("UpdateUserStatus", uint(2), "active").Return(nil).Once()

		err := dashboardService.UpdateUserStatus(1, 2, "active")
		assert.Nil(t, err)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		err := dashboardService.UpdateUserStatus(1, 2, "active")
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
	})

	t.Run("failure - CheckIsAdmin query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, errors.New("check admin query error")).Once()

		err := dashboardService.UpdateUserStatus(1, 2, "active")
		assert.NotNil(t, err)
		assert.Equal(t, "check admin query error", err.Error())
	})

	t.Run("failure - UpdateUserStatus query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("UpdateUserStatus", uint(2), "active").Return(errors.New("query error")).Once()

		err := dashboardService.UpdateUserStatus(1, 2, "active")
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
	})

	t.Run("failure - user not found for update", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("UpdateUserStatus", uint(2), "active").Return(errors.New("record not found")).Once()

		err := dashboardService.UpdateUserStatus(1, 2, "active")
		assert.NotNil(t, err)
		assert.Equal(t, "not found", err.Error())
	})
}

func TestGetDepositStat(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	stats := []dashboards.StatData{
		{Date: "01-08-2024", Total: 100},
	}

	t.Run("success get deposit stats", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetDepositStat", mock.Anything).Return(stats, nil).Once()

		result, err := dashboardService.GetDepositStat(1, "", "", "01-08-2024", "01-08-2024")
		assert.Nil(t, err)
		assert.Equal(t, stats, result)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		result, err := dashboardService.GetDepositStat(1, "", "", "01-08-2024", "01-08-2024")
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
		assert.Equal(t, []dashboards.StatData{}, result)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetDepositStat", mock.Anything).Return(nil, errors.New("query error")).Once()

		result, err := dashboardService.GetDepositStat(1, "", "", "01-08-2024", "01-08-2024")
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, []dashboards.StatData{}, result)
	})

}

func TestGetRewardStatData(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	stats := []dashboards.RewardStatData{
		{RewardID: 1, Total: 50},
	}

	t.Run("success get reward stats", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetRewardStatData", mock.Anything).Return(stats, nil).Once()
		mockQuery.On("GetRewardNameByID", uint(1)).Return("Reward Name", nil).Once()

		result, err := dashboardService.GetRewardStatData(1, "01-08-2024", "01-08-2024")
		assert.Nil(t, err)
		assert.Equal(t, stats, result)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		result, err := dashboardService.GetRewardStatData(1, "01-08-2024", "01-08-2024")
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
		assert.Equal(t, []dashboards.RewardStatData{}, result)
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("GetRewardStatData", mock.Anything).Return(nil, errors.New("query error")).Once()

		result, err := dashboardService.GetRewardStatData(1, "01-08-2024", "01-08-2024")
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Equal(t, []dashboards.RewardStatData{}, result)
	})
}

func TestDeleteUserByAdmin(t *testing.T) {
	mockQuery := new(mocks.DshQuery)
	dashboardService := service.NewDashboardService(mockQuery)

	t.Run("success delete user", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("DeleteUserByAdmin", uint(2)).Return(nil).Once()

		err := dashboardService.DeleteUserByAdmin(1, 2)
		assert.Nil(t, err)
	})

	t.Run("failure - not admin", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()

		err := dashboardService.DeleteUserByAdmin(1, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
	})

	t.Run("failure - self deletion", func(t *testing.T) {
		err := dashboardService.DeleteUserByAdmin(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "not allowed", err.Error())
	})

	t.Run("failure - query error", func(t *testing.T) {
		mockQuery.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		mockQuery.On("DeleteUserByAdmin", uint(2)).Return(errors.New("query error")).Once()

		err := dashboardService.DeleteUserByAdmin(1, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "query error", err.Error())
	})
}
