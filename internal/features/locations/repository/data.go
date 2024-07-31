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
	Name       string
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
		Name:    data.Name,
	}
}

func toLocEntity(dataQry []Location) []locations.LocationInterface {
	result := []locations.LocationInterface{}
	for _, v := range dataQry {
		result = append(result, locations.LocationInterface{
			Address:    v.Address,
			Long:       v.Long,
			Lat:        v.Lat,
			Status:     v.Status,
			Start_time: v.Start,
			End_time:   v.End,
			Phone:      v.Phone,
			UserID:     v.UserID,
			Name:       v.Name,
			ID:         v.ID,
		})
	}
	return result
}

// func toLocEntity(dataQry []Location) []locations.LocationInterface {
// 	returnVal := []locations.LocationInterface{}
// 	for _, v := range dataQry {
// 		dataList := struct {
// 			Address    string
// 			Long       string
// 			Lat        string
// 			Status     string
// 			Start_time string
// 			End_time   string
// 			Phone      string
// 			UserID     uint
// 		}{
// 			Address:    v.Address,
// 			Long:       v.Long,
// 			Lat:        v.Lat,
// 			Status:     v.Status,
// 			Start_time: v.Start,
// 			End_time:   v.End,
// 			Phone:      v.Phone,
// 			UserID:     v.UserID,
// 		}
// 		returnVal = append(returnVal, dataList)
// 	}
// 	return returnVal
// }
