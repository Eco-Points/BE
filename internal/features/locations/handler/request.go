package handler

import "eco_points/internal/features/locations"

type Location struct {
	Address    string `json:"address"`
	Long       string `json:"long"`
	Lat        string `json:"lat"`
	Status     string `json:"status"`
	Start_time string `json:"start"`
	End_time   string `json:"end"`
	Phone      string `json:"phone"`
}

func toLocInterface(data Location, id uint) locations.LocationInterface {
	return locations.LocationInterface{
		Address:    data.Address,
		Long:       data.Long,
		Lat:        data.Lat,
		Status:     data.Status,
		Start_time: data.Start_time,
		End_time:   data.End_time,
		Phone:      data.Phone,
		UserID:     id,
	}

}
