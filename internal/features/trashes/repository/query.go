package repository

import (
	"eco_points/internal/features/trashes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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

func ConvertDateTime(data string) string {
	find := strings.SplitAfter(data, ".")[1]
	depotime := strings.Replace(data, find, "", 1)
	find = strings.SplitAfter(depotime, ".")[1]
	depotime = strings.Replace(depotime, find, "", 1)
	return strings.Replace(depotime, ".", "", 1)
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

func (t *trashQuery) DeleteTrash(id uint) error {
	err := t.db.Debug().Delete(&Trash{}, id).Error
	if err != nil {
		log.Println("Error get data ", err)
		return err
	}
	return nil
}

func (t *trashQuery) UpdateTrash(id uint, data trashes.TrashEntity) error {
	query := `UPDATE "eco_points_dev".trashes set`

	if data.TrashType != "" {
		query = query + fmt.Sprintf(` trash_type = '%s',`, data.TrashType)
	}

	if data.Name != "" {
		query = query + fmt.Sprintf(` trash_name = '%s',`, data.Name)
	}

	if data.Point != "" {
		trash_id, err := strconv.Atoi(data.Point)
		if err != nil {
			log.Println("error convert", err)
		}
		query = query + fmt.Sprintf(` point = '%d',`, trash_id)
	}

	if data.Descrip != "" {
		query = query + fmt.Sprintf(` description = '%s',`, data.Descrip)
	}

	query = query + fmt.Sprintf(` updated_at = '%s'`, ConvertDateTime(fmt.Sprintf("%v", time.Now())))
	query = query + fmt.Sprintf(` WHERE id = '%d' and "deleted_at" IS null;`, id)

	err := t.db.Debug().Exec(query).Error
	if err != nil {
		log.Println("error udapte to table", err)
		return err
	}
	//`trash_type = ?, trash_name = ?, point = ?, image =?,description = ?, updated_at='2013-11-17 21:34:10' WHERE id = ? and "deleted_at" IS null;

	return nil
}
