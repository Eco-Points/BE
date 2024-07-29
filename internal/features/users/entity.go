package users

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	Fullname string
	Email    string
	Password string
	Phone    string
	Address  string
	IsAdmin  bool
	Point    uint
	Status   string
	ImgURL   string
}

type UHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type UQuery interface {
	Register(newUsers User) error
	Login(email string) (User, error)
	GetUser(ID uint) (User, error)
	UpdateUser(ID uint, updateUser User) error
	DeleteUser(ID uint) error
}

type UService interface {
	Register(newUser User) error
	Login(email string, password string) (User, string, error)
	GetUser(ID uint) (User, error)
	UpdateUser(ID uint, updateUser User, file *multipart.FileHeader) error
	DeleteUser(ID uint) error
}
