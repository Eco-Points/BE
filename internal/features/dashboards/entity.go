package dashboards

import (
	"github.com/labstack/echo/v4"
)

type Dashboard struct {
	UserCount     int `json:"user_count"`
	DepositCount  int `json:"deposit_count"`
	ExchangeCount int `json:"exchange_count"`
}

type User struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	CreateAt string `json:"registered_date"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}

type StatData struct {
	Date  string `json:"date"`
	Total int    `json:"total"`
}

type RewardStatData struct {
	RewardID   uint   `json:"reward_id"`
	RewardName string `json:"reward_name"`
	Total      int    `json:"total"`
}
type DshHandler interface {
	GetDashboard() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUserStatus() echo.HandlerFunc
	GetDepositStat() echo.HandlerFunc
	GetRewardStatData() echo.HandlerFunc
}

type DshService interface {
	GetDashboard(userID uint) (Dashboard, error)
	GetAllUsers(userID uint, nameParams string) ([]User, error)
	GetUser(userID uint, targetID uint) (User, error)
	UpdateUserStatus(userID uint, targetID uint, status string) error
	GetDepositStat(userID uint, trashParam, locParam, startDate, endDate string) ([]StatData, error)
	GetRewardStatData(userID uint, startDate, endDate string) ([]RewardStatData, error)
}

type DshQuery interface {
	GetUserCount() (int, error)
	GetExchangeCount() (int, error)
	GetDepositCount() (int, error)
	CheckIsAdmin(userID uint) (bool, error)
	GetAllUsers(nameParams string) ([]User, error)
	GetUser(userID uint) (User, error)
	UpdateUserStatus(target_id uint, status string) error
	GetDepositStat(whereParam string) ([]StatData, error)
	GetRewardStatData(whereParam string) ([]RewardStatData, error)
	GetRewardNameByID(rewardID uint) (string, error)
}
