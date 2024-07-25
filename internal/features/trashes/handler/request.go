package handler

import (
	"eco_points/internal/features/trashes"
	"fmt"
)

type Trash struct {
	TrashType string `json:"type" form:"type"`
	Name      string `json:"name" form:"name"`
	Point     uint   `json:"point" form:"point"`
	Descrip   string `json:"description" form:"description"`
	Image     string `json:"images" form:"images"`
}

func toTrashEntity(tData Trash, id uint) trashes.TrashEntity {
	return trashes.TrashEntity{
		TrashType: tData.TrashType,
		Name:      tData.Name,
		Point:     fmt.Sprintf("%d", tData.Point),
		Descrip:   tData.Descrip,
		Image:     tData.Image,
		UserID:    id,
	}

}
