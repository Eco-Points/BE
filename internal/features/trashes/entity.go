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

type QueryInterface interface {
	AddTrash(tData TrashEntity) error
}

type ServiceInterface interface {
	AddTrash(tData TrashEntity, file *multipart.FileHeader) error
}

type HandlerInterface interface {
	AddTrash() echo.HandlerFunc
}
