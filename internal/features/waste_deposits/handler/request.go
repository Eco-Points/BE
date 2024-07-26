package handler

import deposits "eco_points/internal/features/waste_deposits"

type WasteDeposit struct {
	TrashID    uint `json:"trash_id"`
	LocationID uint `json:"location_id"`
	Quantity   uint `json:"quantity"`
}

type UpdateWasteDepositStatusRequest struct {
	WasteID uint   `json:"waste_id"`
	Status  string `json:"status"`
}

func toWasteDepositInterface(data WasteDeposit, id uint) deposits.WasteDepositInterface {
	return deposits.WasteDepositInterface{
		TrashID:    data.TrashID,
		LocationID: data.LocationID,
		Quantity:   data.Quantity,
		UserID:     id,
	}

}
