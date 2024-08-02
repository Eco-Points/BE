package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ExcelMakeUtilityInterface interface {
	MakeExcel(name string, date string, data interface{}) (string, error)
}
type ExcelMakeUtility struct {
}

func NewExcelMake() ExcelMakeUtilityInterface {
	return &ExcelMakeUtility{}
}

type M map[string]interface {
}

type ExcelDeposit struct {
	ID         uint
	TrashID    uint
	UserID     uint
	LocationID uint
	Quantity   uint
	Point      uint
	Status     string
	Type       string
	DepoTime   string
	Fullname   string
}

type ExcelExchange struct {
	ID           uint
	PointUsed    uint
	Fullname     string
	Reward       string
	ExchangeTime string
}

func (e *ExcelMakeUtility) MakeExcel(name string, date string, data interface{}) (string, error) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if strings.Contains(name, "depo") || strings.Contains(name, "Depo") {
		var saveData []ExcelDeposit

		jsonStr, ok := data.(string)
		if !ok {
			fmt.Println("Error: data is not a string")
			return "", nil
		}

		err := json.Unmarshal([]byte(jsonStr), &saveData)
		if err != nil {
			log.Println("error unmarshal", err)
		}
		f.SetCellValue("Sheet1", "A2", "No.")
		f.SetCellValue("Sheet1", "B2", "Point")
		f.SetCellValue("Sheet1", "C2", "Quantity")
		f.SetCellValue("Sheet1", "D2", "Status")
		f.SetCellValue("Sheet1", "E2", "Type")
		f.SetCellValue("Sheet1", "F2", "Name")
		f.SetCellValue("Sheet1", "G2", "Depo Time")

		for i := 0; i < len(saveData); i++ {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+3), i+1)
			f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+3), saveData[i].Point)
			f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+3), saveData[i].Quantity)
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+3), saveData[i].Status)
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+3), saveData[i].Type)
			f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+3), saveData[i].Fullname)
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+3), saveData[i].DepoTime)

		}
	} else if strings.Contains(name, "rew") || strings.Contains(name, "Rew") {
		var saveData []ExcelExchange

		jsonStr, ok := data.(string)
		if !ok {
			fmt.Println("Error: data is not a string")
			return "", nil
		}

		err := json.Unmarshal([]byte(jsonStr), &saveData)
		if err != nil {
			log.Println("error unmarshal", err)
		}
		f.SetCellValue("Sheet1", "A2", "No.")
		f.SetCellValue("Sheet1", "B2", "Name")
		f.SetCellValue("Sheet1", "C2", "Reward")
		f.SetCellValue("Sheet1", "D2", "Point Used")
		f.SetCellValue("Sheet1", "E2", "Exchange Time")

		for i := 0; i < len(saveData); i++ {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+3), i+1)
			f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+3), saveData[i].Fullname)
			f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+3), saveData[i].Reward)
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+3), saveData[i].PointUsed)
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+3), saveData[i].ExchangeTime)
		}
	} else {
		return "", errors.New("cannot make table")
	}
	excelFile := fmt.Sprintf("%s_%s.xlsx", date, name)
	if err := f.SaveAs(excelFile); err != nil {
		fmt.Println(err)
	}
	return excelFile, nil
}
