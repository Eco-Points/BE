package repository

import (
	"eco_points/internal/features/locations"
	"log"

	"gorm.io/gorm"
)

type locQuery struct {
	db *gorm.DB
}

func NewLocQuery(d *gorm.DB) locations.QueryLocInterface {
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

func (q *locQuery) GetLocation(id uint) ([]locations.LocationInterface, error) {
	var returnVal []Location
	err := q.db.Debug().Where("user_id = ?", id).Find(&returnVal)
	if err.Error != nil {
		log.Println("error get data by id ", err.Error)
		return []locations.LocationInterface{}, err.Error
	}

	return toLocEntity(returnVal), nil

}

func (q *locQuery) GetAllLocation() ([]locations.LocationInterface, error) {
	returnVal := []Location{}
	err := q.db.Debug().Find(&returnVal)
	if err.Error != nil {
		log.Println("error get data ", err.Error)
		return []locations.LocationInterface{}, err.Error
	}

	return toLocEntity(returnVal), nil

}
