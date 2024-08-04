package service

import (
	"eco_points/internal/features/exchanges"
	"errors"
	"log"
)

type ExchangeServices struct {
	qry exchanges.ExQuery
}

func NewExchangeService(q exchanges.ExQuery) exchanges.ExServices {
	return &ExchangeServices{
		qry: q,
	}
}

func (es *ExchangeServices) AddExchange(newExchange exchanges.Exchange) error {
	err := es.qry.AddExchange(newExchange)
	if err != nil {
		return errors.New("failed to create exchange")
	}
	return nil
}

func (es *ExchangeServices) GetExchangeHistory(userid uint, isAdmin bool, limit uint) ([]exchanges.ListExchangeInterface, error) {
	result, err := es.qry.GetExchangeHistory(userid, isAdmin, limit)
	if err != nil {
		log.Println("error on query", err)
		return []exchanges.ListExchangeInterface{}, err

	}
	return result, nil
}
