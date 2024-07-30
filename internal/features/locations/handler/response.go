package handler

import "eco_points/internal/features/locations"

func toResposne(data []locations.LocationInterface) []Location {
	result := []Location{}

	for _, v := range data {
		result = append(result, Location{
			Address:    v.Address,
			Long:       v.Long,
			Lat:        v.Lat,
			Status:     v.Status,
			Start_time: v.Start_time,
			End_time:   v.End_time,
			Phone:      v.Phone,
			UserID:     v.UserID,
		})
	}
	return result
}
