package handler

import "eco_points/internal/features/exchanges"

type ListExchange struct {
	ID           uint   `json:"id"`
	PointUsed    uint   `json:"point_used"`
	Fullname     string `json:"fullname"`
	Reward       string `json:"reward"`
	ExchangeTime string `json:"exchange_time"`
}

func toListExchange(data []exchanges.ListExchangeInterface) []ListExchange {
	returnValue := []ListExchange{}
	for _, v := range data {
		exchange := ListExchange{
			ID:           v.ID,
			PointUsed:    v.PointUsed,
			ExchangeTime: v.ExchangeTime,
			Fullname:     v.Fullname,
			Reward:       v.Reward,
		}
		returnValue = append(returnValue, exchange)
	}
	return returnValue
}
