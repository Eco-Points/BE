package repository

import (
	"eco_points/internal/features/exchanges"
	"errors"

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
