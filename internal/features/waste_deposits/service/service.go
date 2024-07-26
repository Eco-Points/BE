package service

import (
	deposits "eco_points/internal/features/waste_deposits"
	"log"
)

type depositService struct {
	qry deposits.QueryInterface
}

func NewDepositsService(q deposits.QueryInterface) deposits.ServiceInterface {
	return &depositService{
		qry: q,
	}
}

func (q *depositService) DepositTrash(data deposits.WasteDepositInterface) error {
	err := q.qry.DepositTrash(data)
	if err != nil {
		log.Println("error insert data", err)
		return err
	}
	return nil

}

func (q *depositService) UpdateWasteDepositStatus(wasteID uint, status string) error {
	err := q.qry.UpdateWasteDepositStatus(wasteID, status)
	if err != nil {
		log.Println("error updating waste deposit status", err)
		return err
	}
	return nil
}
