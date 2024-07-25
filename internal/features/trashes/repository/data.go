package repository

import (
	"eco_points/internal/features/trashes"
	repoDeposit "eco_points/internal/features/waste_deposits/repository"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type Trash struct {
	gorm.Model
	TrashType string `gorm:"column:trash_type"`
	Name      string `gorm:"column:trash_name"`
	Point     uint
	Descrip   string `gorm:"column:description"`
	Image     string
	UserID    uint
	TrashID   []repoDeposit.WasteDeposit `gorm:"foreignKey:TrashID"`
}

func totrashQuery(data trashes.TrashEntity) Trash {
	dataPoint, err := strconv.Atoi(data.Point)
	if err != nil {
		log.Println("error strconv", err)
	}
	return Trash{
		TrashType: data.TrashType,
		Name:      data.Name,
		Point:     uint(dataPoint),
		Descrip:   data.Descrip,
		Image:     data.Image,
		UserID:    data.UserID,
	}
}
