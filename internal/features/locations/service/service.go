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

func (s *locService) GetLocation(limit uint) ([]locations.LocationInterface, error) {
	var result []locations.LocationInterface
	var err error
	result, err = s.qry.GetLocation(limit)
	if err != nil {
		log.Println("error get data", err)
		return []locations.LocationInterface{}, err
	}

	return result, err
}
