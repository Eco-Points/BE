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

func (q *depositQuery) GetUserDeposit(id uint, limit uint, offset uint) (deposits.ListWasteDepositInterface, error) {
	result := ListWasteDeposit{}
	query := fmt.Sprintf(`select wd.id, t.trash_type, wd.point, wd.quantity, wd.updated_at from "%s".waste_deposits wd 
	join "%s".trashes t on t.id = wd.trash_id 
	where wd.user_id = %d and wd."deleted_at" IS NULL limit %d offset %d;`, DbSchema, DbSchema, id, limit, offset)

	err := q.db.Debug().Raw(query).Scan(&result).Error
	if err != nil {
		log.Println("error select to table", err)
		return deposits.ListWasteDepositInterface{}, err
	}
	rV := toWasteDepositListInterface(result)
	return rV, nil
}
