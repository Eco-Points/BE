package service

import (
	"eco_points/internal/features/locations"
	"log"
)

type locService struct {
	qry locations.QueryLocInterface
}

func NewLocService(q locations.QueryLocInterface) locations.ServiceLocInterface {
	return &locService{
		qry: q,
	}
}

func (s *locService) AddLocation(data locations.LocationInterface) error {
	err := s.qry.AddLocation(data)
	if err != nil {
		log.Println("error insert data", err)
		return err
	}
	return nil
}

func (s *locService) GetLocation(id uint) ([]locations.LocationInterface, error) {
	var result []locations.LocationInterface
	var err error
	if id != 0 {
		result, err = s.qry.GetLocation(id)
		if err != nil {
			log.Println("error get data", err)
			return []locations.LocationInterface{}, err
		}
	} else {
		result, err = s.qry.GetAllLocation()
		if err != nil {
			log.Println("error get data", err)
			return []locations.LocationInterface{}, err
		}
	}
	return result, err
}
