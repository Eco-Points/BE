package service

import (
	deposits "eco_points/internal/features/waste_deposits"
	"eco_points/internal/utils"
	"log"
	"strconv"
	"strings"
)

type depositService struct {
	qry   deposits.QueryDepoInterface
	email utils.GomailUtilityInterface
}

func NewDepositsService(q deposits.QueryDepoInterface, e utils.GomailUtilityInterface) deposits.ServiceDepoInterface {
	return &depositService{
		qry:   q,
		email: e,
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
	userID, points, err := q.qry.UpdateWasteDepositStatus(wasteID, status)
	if err != nil {
		log.Println("error updating waste deposit status", err)
		return err
	}
	userEmail, userName, err := q.qry.GetUserEmailData(userID)
	if err != nil {
		log.Println("error send email", err)
		return err
	}
	message := "Halo, <b>" + userName + "</b>"
	if strings.Contains(strings.ToLower(status), "verif") {
		message = message + "<br/> Pengajuan deposit sampahmu sudah diverifikasi oleh admin keren dari tim Eco Points. <br/>Selamat kamu mendapatkan <b>" + strconv.Itoa(points) + " poin</b> "
	} else {
		message = message + "<br/> Pengajuan deposit sampahmu belum memenuhi syarat untuk mendapatkan poin"
	}
	err = q.email.SendEmail(points, message, userEmail, userName)
	if err != nil {
		log.Println("error send email", err)
		return err
	}
	return nil
}

func (s *depositService) GetUserDeposit(id uint, limit uint, offset uint, is_admin bool) (deposits.ListWasteDepositInterface, error) {
	result, err := s.qry.GetUserDeposit(id, limit, offset, is_admin)
	if err != nil {
		log.Println("error insert data", err)
		return deposits.ListWasteDepositInterface{}, err
	}
	return result, nil
}

func (s *depositService) GetDepositbyId(deposit_id uint) (deposits.WasteDepositInterface, error) {
	result, err := s.qry.GetDepositbyId(deposit_id)
	if err != nil {
		log.Println("error insert data", err)
		return deposits.WasteDepositInterface{}, err
	}
	return result, nil
}
