package repository

import (
	"eco_points/internal/features/exchanges"
	"strings"

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

type ListExchange struct {
	gorm.Model
	ID        uint   `gorm:"column:id"`
	PointUsed uint   `gorm:"column:point_used"`
	Fullname  string `gorm:"column:name"`
	Reward    string `gorm:"column:rewards"`
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

func ConvertDateTime(data string) string {
	find := strings.SplitAfter(data, ".")[1]
	depotime := strings.Replace(data, find, "", 1)
	find = strings.SplitAfter(depotime, ".")[1]
	depotime = strings.Replace(depotime, find, "", 1)
	return strings.Replace(depotime, ".", "", 1)
}

func toListExchangeInterface(data []ListExchange) []exchanges.ListExchangeInterface {
	var returnValue []exchanges.ListExchangeInterface
	for _, v := range data {
		exchange := exchanges.ListExchangeInterface{
			ID:           v.ID,
			PointUsed:    v.PointUsed,
			ExchangeTime: ConvertDateTime(v.CreatedAt.String()),
			Fullname:     v.Fullname,
			Reward:       v.Reward,
		}
		returnValue = append(returnValue, exchange)
	}
	return returnValue
}
