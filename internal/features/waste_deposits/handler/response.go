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
	Status   string `json:"status"`
	Fullname string `json:"fullname"`
}

type RespWasteDeposit struct {
	Type     string `json:"type"`
	Point    uint   `json:"point"`
	DepoTime string `json:"depotime"`
	Quantity uint   `json:"quantity"`
	ID       uint   `json:"id"`
	Status   string `json:"status"`
	Fullname string `json:"fullname"`
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
			Status   string
			Fullname string
		}{
			Type:     v.Type,
			Point:    v.Point,
			DepoTime: ConvertDateTime(v.DepoTime),
			Quantity: v.Quantity,
			ID:       v.ID,
			Status:   v.Status,
			Fullname: v.Fullname,
		}
		result = append(result, struct {
			Type     string "json:\"type\""
			Point    uint   "json:\"point\""
			DepoTime string "json:\"depotime\""
			Quantity uint   "json:\"quantity\""
			ID       uint   "json:\"id\""
			Status   string "json:\"status\""
			Fullname string "json:\"fullname\""
		}(dataList))
	}
	return result
}

func toWasteDepositResponse(data deposits.WasteDepositInterface) RespWasteDeposit {
	return RespWasteDeposit{
		Type:     data.Type,
		Point:    data.Point,
		DepoTime: ConvertDateTime(data.DepoTime),
		Quantity: data.Quantity,
		ID:       data.ID,
		Status:   data.Status,
		Fullname: data.Fullname,
	}

}
