package repository

import (
	"eco_points/internal/features/trashes"
	"log"

	"gorm.io/gorm"
)

type trashQuery struct {
	db *gorm.DB
}

func NewTrashQuery(dbQuery *gorm.DB) trashes.QueryTrashInterface {
	return &trashQuery{
		db: dbQuery,
	}
}

func (t *trashQuery) AddTrash(tData trashes.TrashEntity) error {
	input := totrashQuery(tData)
	err := t.db.Create(&input).Error
	if err != nil {
		log.Println("Error Inser", err)
		return err
	}
	return nil
}

func (t *trashQuery) GetTrashbyType(ttype string) (trashes.ListTrashEntity, error) {
	input := []Trash{}
	err := t.db.Debug().Find(&input, "trash_type = ?", ttype).Error
	if err != nil {
		log.Println("Error get data ", err)
		return trashes.ListTrashEntity{}, err
	}
	return toListTrashQuery(input), nil
}

func (t *trashQuery) GetTrashLimit() (trashes.ListTrashEntity, error) {
	input := []Trash{}
	err := t.db.Debug().Limit(10).Find(&input).Error
	if err != nil {
		log.Println("Error get data ", err)
		return trashes.ListTrashEntity{}, err
	}
	return toListTrashQuery(input), nil
}
