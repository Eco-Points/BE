package repository

import (
	"eco_points/internal/features/exchanges"

	"gorm.io/gorm"
)

type Exchange struct {
	gorm.Model
	RewardID  uint
	UserID    uint
	PointUsed uint32
	User      User   `gorm:"foreignKey:UserID"`
	Reward    Reward `gorm:"foreignKey:RewardID"`
}

type User struct {
	gorm.Model
	Fullname string
	Email    string
	Password string
	Phone    string
	Address  string
	IsAdmin  bool
	Point    uint
	ImgURL   string
	Exchange []Exchange `gorm:"foreignKey:UserID"`
}

type Reward struct {
	gorm.Model
	Name          string
	Description   string
	PointRequired uint32
	Stock         uint32
	Image         string
	Exchange      []Exchange `gorm:"foreignKey:RewardID"`
}

func (ex *Exchange) ToExchangeEntity() exchanges.Exchange {
	return exchanges.Exchange{
		ID:        ex.ID,
		RewardID:  ex.RewardID,
		UserID:    ex.UserID,
		PointUsed: ex.PointUsed,
		CreatedAt: ex.CreatedAt,
	}
}

func ToExchangeData(input exchanges.Exchange) Exchange {
	return Exchange{
		RewardID:  input.RewardID,
		UserID:    input.UserID,
		PointUsed: input.PointUsed,
	}
}
