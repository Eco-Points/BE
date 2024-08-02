package repository

import (
	"eco_points/internal/features/exchanges"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ExchangeModel struct {
	db *gorm.DB
}

func NewExchangeModel(connection *gorm.DB) exchanges.ExQuery {
	return &ExchangeModel{
		db: connection,
	}
}

var DbSchema string

func (q *ExchangeModel) SetDbSchema(schema string) {
	DbSchema = schema
}

func (em *ExchangeModel) AddExchange(newExchange exchanges.Exchange) error {
	var user User
	var reward Reward

	// Check if reward exists and has stock
	if err := em.db.First(&reward, newExchange.RewardID).Error; err != nil {
		return errors.New("reward not found")
	}
	if reward.Stock < 1 {
		return errors.New("reward out of stock")
	}

	// Check if user has enough points
	if err := em.db.First(&user, newExchange.UserID).Error; err != nil {
		return errors.New("user not found")
	}
	if uint32(user.Point) < reward.PointRequired {
		return errors.New("not enough points")
	}

	// Deduct points from user and reduce stock from reward
	user.Point -= uint(reward.PointRequired)
	reward.Stock -= 1

	// Set the points used for the exchange
	newExchange.PointUsed = reward.PointRequired

	cnvExchange := ToExchangeData(newExchange)

	// Create exchange
	if err := em.db.Create(&cnvExchange).Error; err != nil {
		return err
	}

	// Update user and reward
	if err := em.db.Save(&user).Error; err != nil {
		return err
	}
	if err := em.db.Save(&reward).Error; err != nil {
		return err
	}

	return nil
}

func (em *ExchangeModel) GetExchangeHistory(userid uint, isAdmin bool, limit uint) ([]exchanges.ListExchangeInterface, error) {
	var exchange []ListExchange
	if isAdmin {
		query := fmt.Sprintf(`select e.id, e.point_used, r."name" as rewards, u.fullname as name, e.created_at from "%s".exchanges e 
							join "%s".rewards r on r.id = e.reward_id	
							join "%s".users u on u.id = e.user_id
							where e."deleted_at" IS NULL `, DbSchema, DbSchema, DbSchema)
		if limit != 0 {
			query = query + fmt.Sprintf("limit %d", limit)
		}
		err := em.db.Debug().Raw(query).Scan(&exchange).Error
		if err != nil {
			log.Println("error get data", err)
			return []exchanges.ListExchangeInterface{}, err
		}
		return toListExchangeInterface(exchange), nil
	} else {
		query := fmt.Sprintf(`select e.id, e.point_used, r."name" as rewards, u.fullname as name, e.created_at from "%s".exchanges e 
							join "%s".rewards r on r.id = e.reward_id	
							join "%s".users u on u.id = e.user_id
							where user_id = %d and e."deleted_at" IS NULL `, DbSchema, DbSchema, DbSchema, userid)
		if limit != 0 {
			query = query + fmt.Sprintf("limit %d;", limit)
		}
		err := em.db.Debug().Raw(query).Scan(&exchange).Error
		if err != nil {
			log.Println("error get data", err)
			return []exchanges.ListExchangeInterface{}, err
		}
		return toListExchangeInterface(exchange), nil
	}
}
