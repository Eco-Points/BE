package repository

import (
	"eco_points/internal/features/locations"
	"log"

	"gorm.io/gorm"
)

type locQuery struct {
	db *gorm.DB
}

func NewLocQuery(d *gorm.DB) locations.QueryInterface {
	return &locQuery{
		db: d,
	}
}

func (q *locQuery) AddLocation(data locations.LocationInterface) error {
	input := toLocQuery(data)
	err := q.db.Create(&input)
	if err != nil {
		log.Println("error insert to table ", err)
		return err.Error
	}

	return nil

}
