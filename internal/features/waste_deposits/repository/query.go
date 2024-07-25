package repository

import (
	deposits "eco_points/internal/features/waste_deposits"
	"log"

	"gorm.io/gorm"
)

type depositQuery struct {
	db *gorm.DB
}

func NewTrashQuery(d *gorm.DB) deposits.QueryInterface {
	return &depositQuery{
		db: d,
	}
}

func (d *depositQuery) DepositTrash(data deposits.WasteDepositInterface) error {
	err := d.db.Debug().Model(&Trash{}).Select("point").Where("id = ?", data.TrashID).First(&data.Point).Error
	if err != nil {
		log.Println("error trash type not available", err)
		return err
	}

	input := toWasteDeposit(data)
	err = d.db.Debug().Create(&input).Error
	if err != nil {
		log.Println("error insert to tabel", err)
		return err
	}
	return nil
}
