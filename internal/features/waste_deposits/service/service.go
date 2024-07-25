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
