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

func (s *depositService) DepositTrash(data deposits.WasteDepositInterface) error {
	err := s.qry.DepositTrash(data)
	if err != nil {
		log.Println("error insert data", err)
		return err
	}
	return nil

}
func (s *depositService) GetUserDeposit(id uint) (deposits.ListWasteDepositInterface, error) {
	result, err := s.qry.GetUserDeposit(id)
	if err != nil {
		log.Println("error insert data", err)
		return deposits.ListWasteDepositInterface{}, err
	}
	return result, nil
}
