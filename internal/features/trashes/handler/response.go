package handler

import (
	"eco_points/internal/features/trashes"
)

type ListTrash []struct {
	TrashType string `json:"type"`
	Name      string `json:"name"`
	Point     uint   `json:"point"`
	Descrip   string `json:"description"`
	Image     string `json:"images"`
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
}

func toListTrashResponse(data trashes.ListTrashEntity) ListTrash {
	var result ListTrash
	for _, v := range data {
		dataList := struct {
			TrashType string `json:"type"`
			Name      string `json:"name"`
			Point     uint   `json:"point"`
			Descrip   string `json:"description"`
			Image     string `json:"images"`
			ID        uint   `json:"id"`
			UserID    uint   `json:"user_id"`
		}{
			TrashType: v.TrashType,
			Name:      v.Name,
			Point:     v.Point,
			Descrip:   v.Descrip,
			Image:     v.Image,
			UserID:    v.UserID,
			ID:        v.ID,
		}

		result = append(result, dataList)
	}
	return result
}
