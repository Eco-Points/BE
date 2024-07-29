package service

import (
	"eco_points/internal/features/exchanges"
	"errors"
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
