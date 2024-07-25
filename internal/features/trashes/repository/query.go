package repository

import (
	"eco_points/internal/features/trashes"
	"log"

	"gorm.io/gorm"
)

type trashQuery struct {
	db *gorm.DB
}

func NewTrashQuery(dbQuery *gorm.DB) trashes.QueryInterface {
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
