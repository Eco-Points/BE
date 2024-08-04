package service_test

import (
	"eco_points/internal/features/dashboards"
	"eco_points/internal/features/dashboards/service"
	"eco_points/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetDashboard(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	expectedResult := dashboards.Dashboard{
		UserCount:     1,
		DepositCount:  2,
		ExchangeCount: 3,
	}
	t.Run("success get dashboard", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(1, nil).Once()
		qry.On("GetDepositCount").Return(2, nil).Once()
		qry.On("GetExchangeCount").Return(3, nil).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("success get dashboard 2", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(0, gorm.ErrRecordNotFound).Once()
		qry.On("GetDepositCount").Return(2, nil).Once()
		qry.On("GetExchangeCount").Return(3, nil).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.UserCount)
		assert.Equal(t, expectedResult.DepositCount, result.DepositCount)
		assert.Equal(t, expectedResult.ExchangeCount, result.ExchangeCount)
	})
	t.Run("success get dashboard 3", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(0, gorm.ErrRecordNotFound).Once()
		qry.On("GetDepositCount").Return(0, gorm.ErrRecordNotFound).Once()
		qry.On("GetExchangeCount").Return(3, nil).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.UserCount)
		assert.Equal(t, 0, result.DepositCount)
		assert.Equal(t, expectedResult.ExchangeCount, result.ExchangeCount)
	})
	t.Run("success get dashboard 4", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(0, gorm.ErrRecordNotFound).Once()
		qry.On("GetDepositCount").Return(0, gorm.ErrRecordNotFound).Once()
		qry.On("GetExchangeCount").Return(0, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.UserCount)
		assert.Equal(t, 0, result.DepositCount)
		assert.Equal(t, 0, result.ExchangeCount)
	})
	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.Dashboard{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.Dashboard{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.Dashboard{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user count query", func(t *testing.T) {
		expectedError := errors.New("an unexpected error occurred")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(0, errors.New("some error")).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.Dashboard{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on Deposit count query", func(t *testing.T) {
		expectedError := errors.New("an unexpected error occurred")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(1, nil).Once()
		qry.On("GetDepositCount").Return(0, errors.New("some error")).Once()
		result, err := srv.GetDashboard(uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.Dashboard{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on Exchange count query", func(t *testing.T) {
		expectedError := errors.New("an unexpected error occurred")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUserCount").Return(1, nil)
		qry.On("GetDepositCount").Return(2, nil)
		qry.On("GetExchangeCount").Return(0, errors.New("some error"))
		result, err := srv.GetDashboard(uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.Dashboard{}, result)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	expectedResult1 := dashboards.User{
		ID:       1,
		Fullname: "a",
	}
	expectedResult2 := dashboards.User{
		ID:       2,
		Fullname: "b",
	}
	var expectedResult []dashboards.User
	expectedResult = append(expectedResult, expectedResult1, expectedResult2)

	t.Run("success get all users", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetAllUsers", "%%").Return(expectedResult, nil).Once()
		result, err := srv.GetAllUsers(uint(1), "")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})
	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetAllUsers(uint(1), "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		result, err := srv.GetAllUsers(uint(1), "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		result, err := srv.GetAllUsers(uint(1), "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user user not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetAllUsers", "%%").Return([]dashboards.User{}, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetAllUsers(uint(1), "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on Deposit get all users query", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetAllUsers", "%%").Return([]dashboards.User{}, errors.New("other error")).Once()
		result, err := srv.GetAllUsers(uint(1), "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetUser(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	expectedResult := dashboards.User{
		ID:       1,
		Fullname: "a",
	}

	t.Run("success get all users", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUser", uint(1)).Return(expectedResult, nil).Once()
		result, err := srv.GetUser(uint(1), uint(1))
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetUser(uint(1), uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("check admin query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		result, err := srv.GetUser(uint(1), uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		result, err := srv.GetUser(uint(1), uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user on other user not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUser", uint(1)).Return(dashboards.User{}, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetUser(uint(1), uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on get user query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetUser", uint(1)).Return(dashboards.User{}, errors.New("some other error")).Once()
		result, err := srv.GetUser(uint(1), uint(1))
		assert.NotNil(t, result, err)
		assert.Equal(t, dashboards.User{}, result)
		assert.Equal(t, expectedError, err)
	})

}

func TestUpdateUserStatus(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	t.Run("success get update user status", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("UpdateUserStatus", uint(2), "suspended").Return(nil).Once()
		err := srv.UpdateUserStatus(uint(1), uint(2), "suspended")
		assert.Nil(t, err)
	})

	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		err := srv.UpdateUserStatus(uint(1), uint(2), "suspended")
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("check admin query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		err := srv.UpdateUserStatus(uint(1), uint(2), "suspended")
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		err := srv.UpdateUserStatus(uint(1), uint(2), "suspended")
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on target user not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("UpdateUserStatus", uint(2), "suspended").Return(gorm.ErrRecordNotFound).Once()
		err := srv.UpdateUserStatus(uint(1), uint(2), "suspended")
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on update user query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("UpdateUserStatus", uint(2), "suspended").Return(errors.New("some query error")).Once()
		err := srv.UpdateUserStatus(uint(1), uint(2), "suspended")
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

}

func TestGetDepositStat(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	expectedResult1 := dashboards.StatData{
		Date:  "01-02-2024",
		Total: 10,
	}
	expectedResult2 := dashboards.StatData{
		Date:  "02-02-2024",
		Total: 9,
	}
	var expectedResult []dashboards.StatData
	expectedResult = append(expectedResult, expectedResult1, expectedResult2)

	t.Run("success get deposit stat", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetDepositStat", "status = 'verified'").Return(expectedResult, nil).Once()
		result, err := srv.GetDepositStat(uint(1), "", "", "", "")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})
	t.Run("success get deposit stat with params", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetDepositStat", "status = 'verified' AND trash_id = 1 AND location_id = 1 AND created_at BETWEEN '2024-02-01 00:00:00' AND date '2024-02-02 00:00:00' + interval '1 day'").Return(expectedResult, nil).Once()
		result, err := srv.GetDepositStat(uint(1), "1", "1", "01-02-2024", "02-02-2024")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})
	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetDepositStat(uint(1), "", "", "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.StatData{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("check admin query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		result, err := srv.GetDepositStat(uint(1), "", "", "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.StatData{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		result, err := srv.GetDepositStat(uint(1), "", "", "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.StatData{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on data stat not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetDepositStat", "status = 'verified'").Return([]dashboards.StatData{}, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetDepositStat(uint(1), "", "", "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.StatData{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on get stat data query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetDepositStat", "status = 'verified'").Return([]dashboards.StatData{}, errors.New("some other error")).Once()
		result, err := srv.GetDepositStat(uint(1), "", "", "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.StatData{}, result)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetRewardStatData(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	expectedResult1 := dashboards.RewardStatData{
		RewardID:   1,
		RewardName: "Pesawat",
		Total:      2,
	}
	expectedResult2 := dashboards.RewardStatData{
		RewardID:   3,
		RewardName: "Becak",
		Total:      4,
	}
	var expectedResult []dashboards.RewardStatData
	expectedResult = append(expectedResult, expectedResult1, expectedResult2)

	t.Run("success get reward stat data", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetRewardStatData", "").Return(expectedResult, nil).Once()
		qry.On("GetRewardNameByID", uint(1)).Return("Pesawat", nil).Once()
		qry.On("GetRewardNameByID", uint(3)).Return("Becak", nil).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})
	t.Run("success get reward stat data with params", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetRewardStatData", "created_at BETWEEN date '2024-02-01 00:00:00' - INTERVAL '1 DAY' AND date '2024-02-02 00:00:00' + INTERVAL '1 DAY'").Return(expectedResult, nil).Once()
		qry.On("GetRewardNameByID", uint(1)).Return("Pesawat", nil).Once()
		qry.On("GetRewardNameByID", uint(3)).Return("Becak", nil).Once()
		result, err := srv.GetRewardStatData(uint(1), "01-02-2024", "02-02-2024")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})
	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("check admin query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on reward data stat not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetRewardStatData", "").Return([]dashboards.RewardStatData{}, gorm.ErrRecordNotFound).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on get reward stat data query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetRewardStatData", "").Return([]dashboards.RewardStatData{}, errors.New("some other error")).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on reward data stat name not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetRewardStatData", "").Return(expectedResult, nil).Once()
		qry.On("GetRewardNameByID", uint(1)).Return("", gorm.ErrRecordNotFound).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on reward data stat name query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("GetRewardStatData", "").Return(expectedResult, nil).Once()
		qry.On("GetRewardNameByID", uint(1)).Return("", errors.New("some other error")).Once()
		result, err := srv.GetRewardStatData(uint(1), "", "")
		assert.NotNil(t, result, err)
		assert.Equal(t, []dashboards.RewardStatData{}, result)
		assert.Equal(t, expectedError, err)
	})
}

func TestDeleteUserByAdmin(t *testing.T) {
	qry := mocks.NewDshQuery(t)
	srv := service.NewDashboardService(qry)

	t.Run("success delete user account", func(t *testing.T) {
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("DeleteUserByAdmin", uint(2)).Return(nil).Once()
		err := srv.DeleteUserByAdmin(uint(1), uint(2))
		assert.Nil(t, err)
	})

	t.Run("error on admin try to delete itself", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		err := srv.DeleteUserByAdmin(uint(1), uint(1))
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on check isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, gorm.ErrRecordNotFound).Once()
		err := srv.DeleteUserByAdmin(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on check isNotAdmin query error", func(t *testing.T) {
		expectedError := errors.New("check admin query error")
		qry.On("CheckIsAdmin", uint(1)).Return(false, errors.New("query error")).Once()
		err := srv.DeleteUserByAdmin(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user isNotAdmin", func(t *testing.T) {
		expectedError := errors.New("not allowed")
		qry.On("CheckIsAdmin", uint(1)).Return(false, nil).Once()
		err := srv.DeleteUserByAdmin(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("error on user account not found", func(t *testing.T) {
		expectedError := errors.New("not found")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("DeleteUserByAdmin", uint(2)).Return(gorm.ErrRecordNotFound).Once()
		err := srv.DeleteUserByAdmin(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("error on delet user account query error", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("CheckIsAdmin", uint(1)).Return(true, nil).Once()
		qry.On("DeleteUserByAdmin", uint(2)).Return(errors.New("other error")).Once()
		err := srv.DeleteUserByAdmin(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}
