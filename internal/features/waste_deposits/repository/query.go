package repository

import (
	deposits "eco_points/internal/features/waste_deposits"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type depositQuery struct {
	db *gorm.DB
}

func NewDepoQuery(d *gorm.DB) deposits.QueryInterface {
	return &depositQuery{
		db: d,
	}
}

var DbSchema string

func (q *depositQuery) SetDbSchema(schema string) {
	DbSchema = schema
}

func (q *depositQuery) DepositTrash(data deposits.WasteDepositInterface) error {
	err := q.db.Debug().Model(&Trash{}).Select("point").Where("id = ?", data.TrashID).First(&data.Point).Error
	if err != nil {
		log.Println("error trash type not available", err)
		return err
	}

	input := toWasteDeposit(data)
	err = q.db.Debug().Create(&input).Error
	if err != nil {
		log.Println("error insert to table", err)
		return err
	}
	return nil
}

func (d *depositQuery) UpdateWasteDepositStatus(waste_id uint, status string) error {
	err := d.db.Debug().Model(&WasteDeposit{}).Where("id = ?", waste_id).Update("status", status).Error
	if err != nil {
		log.Println("error insert to tabel", err)
		return err
	}
	return nil
}

func (q *depositQuery) GetUserDeposit(id uint, limit uint, offset uint) (deposits.ListWasteDepositInterface, error) {
	result := []ListWasteDeposit{}
	query := fmt.Sprintf(`select wd.id, t.trash_type, wd.point, wd.status, wd.quantity, u.fullname, wd.created_at from "%s".waste_deposits wd 
	join "%s".trashes t on t.id = wd.trash_id 
	join "%s".users u on u.id = wd.user_id 
	where wd.user_id = %d and wd."deleted_at" IS NULL limit %d offset %d;`, DbSchema, DbSchema, DbSchema, id, limit, offset)

	err := q.db.Debug().Raw(query).Scan(&result).Error
	if err != nil {
		log.Println("error select to table", err)
		return deposits.ListWasteDepositInterface{}, err
	}
	rV := toWasteDepositListInterface(result)
	return rV, nil
}

func (d *depositQuery) GetDepositbyId(deposit_id uint) (deposits.WasteDepositInterface, error) {
	result := WasteDeposit{}
	err := d.db.First(&result, deposit_id)

	if err.Error != nil {
		log.Println("error select to table", err)
		return deposits.WasteDepositInterface{}, err.Error
	}

	trashData := Trash{}
	err = d.db.First(&trashData, result.TrashID)

	if err.Error != nil {
		log.Println("error select to table", err)
		return deposits.WasteDepositInterface{}, err.Error
	}

	userData := User{}
	err = d.db.First(&userData, result.UserID)

	if err.Error != nil {
		log.Println("error select to table", err)
		return deposits.WasteDepositInterface{}, err.Error
	}

	return toWasteDepositInterface(result, trashData, userData), err.Error
}
