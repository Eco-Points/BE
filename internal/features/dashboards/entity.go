package dashboards

import "github.com/labstack/echo/v4"

type Dashboard struct {
	UserCount     int `json:"user_count"`
	DepositCount  int `json:"deposit_count"`
	ExchangeCount int `json:"exchange_count"`
}

type User struct {
	ID       uint
	Fullname string
	Email    string
	Password string
	Phone    string
	Address  string
	IsAdmin  bool
	Status   string
	Point    uint
	ImgURL   string
}

type DshHandler interface {
	GetDashboard() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUserStatus() echo.HandlerFunc
}

type DshService interface {
	GetDashboard(userID uint) (Dashboard, error)
	GetAllUsers(userID uint, nameParams string) ([]User, error)
	GetUser(userID uint, targetID uint) (User, error)
	UpdateUserStatus(userID uint, targetID uint, status string) error
}

type DshQuery interface {
	GetUserCount() (int, error)
	GetExchangeCount() (int, error)
	GetDepositCount() (int, error)
	CheckIsAdmin(userID uint) (bool, error)
	GetAllUsers(nameParams string) ([]User, error)
	GetUser(userID uint) (User, error)
	UpdateUserStatus(target_id uint, status string) error
}
