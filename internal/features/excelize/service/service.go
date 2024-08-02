package service

import (
	"eco_points/internal/features/excelize"
	"fmt"
	"log"
	"os"
)

type excelMake struct {
	qry excelize.QueryMakeExcelInterface
}

func NewExcelMake(q excelize.QueryMakeExcelInterface) excelize.ServiceMakeExcelInterface {
	return &excelMake{
		qry: q,
	}
}

func (s *excelMake) MakeExcel(date string, table string, userid uint, isVerif bool, limit uint) (string, error) {
	name, link, err := s.qry.MakeExcelFromDeposit(table, date, userid, isVerif, limit)
	if err != nil {
		log.Println("error from service ", err)
		if name != "" {
			if err := s.DeleteFile(fmt.Sprintf("./%s", name)); err != nil {
				log.Println("Error deleting file: ", err)
			}
		}
		return "", err
	}

	if name != "" {

		if err := s.DeleteFile(fmt.Sprintf("./%s", name)); err != nil {
			log.Println("Error deleting file: ", err)
		}
	}
	return link, nil
}

func (s *excelMake) DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Failed to delete file %s: %v\n", filePath, err)
		return err
	}
	return nil
}
