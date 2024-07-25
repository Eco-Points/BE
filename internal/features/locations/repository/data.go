package repository

import (
	"eco_points/internal/features/locations"
	depoQuery "eco_points/internal/features/waste_deposits/repository"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Address    string
	Long       string
	Lat        string
	Status     string
	Start      string
	End        string
	Phone      string
	UserID     uint
	LocationID []depoQuery.WasteDeposit `gorm:"foreignKey:LocationID"`
}

func toLocQuery(data locations.LocationInterface) Location {
	return Location{
		Address: data.Address,
		Long:    data.Long,
		Lat:     data.Lat,
		Status:  data.Status,
		Start:   data.Start_time,
		End:     data.End_time,
		Phone:   data.Phone,
		UserID:  data.UserID,
	}
}
