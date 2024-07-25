package handler

import (
	deposits "eco_points/internal/features/waste_deposits"
	"strings"
)

type ListWasteDeposit []struct {
	Type     string `json:"type"`
	Point    uint   `json:"point"`
	DepoTime string `json:"depotime"`
	Quantity uint   `json:"quantity"`
	ID       uint   `json:"id"`
}

func ConvertDateTime(data string) string {
	find := strings.SplitAfter(data, ".")[1]
	depotime := strings.Replace(data, find, "", 1)
	return strings.Replace(depotime, ".", "", 1)
}

func toListWasteDepositResponse(data deposits.ListWasteDepositInterface) ListWasteDeposit {
	var result ListWasteDeposit

	for _, v := range data {

		dataList := struct {
			Type     string
			Point    uint
			DepoTime string
			Quantity uint
			ID       uint
		}{
			Type:     v.Type,
			Point:    v.Point,
			DepoTime: ConvertDateTime(v.DepoTime),
			Quantity: v.Quantity,
			ID:       v.ID,
		}
		result = append(result, struct {
			Type     string "json:\"type\""
			Point    uint   "json:\"point\""
			DepoTime string "json:\"depotime\""
			Quantity uint   "json:\"quantity\""
			ID       uint   "json:\"id\""
		}(dataList))
	}

	return result
}
