package services

import (
	"eco_points/internal/features/locations"
	"log"
)

type locService struct {
	qry locations.QueryInterface
}

func NewLocService(q locations.QueryInterface) locations.ServiceInterface {
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
