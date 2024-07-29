package service

import (
	deposits "eco_points/internal/features/waste_deposits"
	"fmt"
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

func (s *depositService) DepositTrash(data deposits.WasteDepositInterface) error {
	err := s.qry.DepositTrash(data)
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

func (s *depositService) GetUserDeposit(id uint, limit uint, offset uint) (deposits.ListWasteDepositInterface, error) {
	result, err := s.qry.GetUserDeposit(id, limit, offset)
	if err != nil {
		log.Println("error insert data", err)
		return deposits.ListWasteDepositInterface{}, err
	}
	return result, nil
}

func (s *depositService) GetDepositbyId(deposit_id uint) (deposits.WasteDepositInterface, error) {
	fmt.Println("run service")
	result, err := s.qry.GetDepositbyId(deposit_id)
	if err != nil {
		log.Println("error insert data", err)
		return deposits.WasteDepositInterface{}, err
	}
	return result, nil
}
