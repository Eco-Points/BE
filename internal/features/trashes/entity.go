package trashes

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type TrashEntity struct {
	TrashType string
	Name      string
	Point     string
	Descrip   string
	Image     string
	UserID    uint
}

type ListTrashEntity []struct {
	TrashType string
	Name      string
	Point     uint
	Descrip   string
	Image     string
	UserID    uint
	ID        uint
}

type QueryTrashInterface interface {
	AddTrash(tData TrashEntity) error
	GetTrashbyType(ttype string) (ListTrashEntity, error)
	GetTrashLimit() (ListTrashEntity, error)
	DeleteTrash(id uint) error
}

type ServiceTrashInterface interface {
	AddTrash(tData TrashEntity, file *multipart.FileHeader) error
	GetTrash(ttype string) (ListTrashEntity, error)
	DeleteTrash(id uint) error
}

type HandlerTrashInterface interface {
	AddTrash() echo.HandlerFunc
	GetTrash() echo.HandlerFunc
	DeleteTrash() echo.HandlerFunc
}
