package repository

import (
	"eco_points/internal/features/dashboards"
	exQuery "eco_points/internal/features/exchanges/repository"
	lochQuery "eco_points/internal/features/locations/repository"
	trashQuery "eco_points/internal/features/trashes/repository"
	depoQuery "eco_points/internal/features/waste_deposits/repository"
	"time"

	"gorm.io/gorm"
)

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
	Status   string
	Trash    []trashQuery.Trash       `gorm:"foreignKey:UserID"`
	Location []lochQuery.Location     `gorm:"foreignKey:UserID"`
	Deposit  []depoQuery.WasteDeposit `gorm:"foreignKey:UserID"`
}

type Reward struct {
	gorm.Model
	Name          string
	Description   string
	PointRequired uint32
	Stock         uint32
	Image         string
	DeletedAt     gorm.DeletedAt     `gorm:"index"`
	Exchange      []exQuery.Exchange `gorm:"foreignKey:RewardID"`
}

type StatData struct {
	Date  time.Time `json:"date"`
	Total int       `json:"total"`
}

type RewardStatData struct {
	RewardID uint `json:"date"`
	Total    int  `json:"total"`
}

// dari database di pindah ke entity
func (s *StatData) ToStatDataEntity() dashboards.StatData {
	return dashboards.StatData{

		Total: s.Total,
		Date:  s.Date.Format("02-01-2006"),
	}
}

// dari database di pindah ke entity
func (u *User) ToUserEntity() dashboards.User {
	return dashboards.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		CreateAt: u.CreatedAt.Format("02-01-2006"),
		Email:    u.Email,
		Phone:    u.Phone,
		Address:  u.Address,
		Status:   u.Status,
	}
}

// dari entity pindah ke database
func ToUserQuery(input dashboards.User) User {
	return User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
		Address:  input.Address,
		Status:   input.Status,
	}
}

func (r *RewardStatData) ToRewardStatDataEntity() dashboards.RewardStatData {
	return dashboards.RewardStatData{
		RewardID: r.RewardID,
		Total:    r.Total,
	}
}
