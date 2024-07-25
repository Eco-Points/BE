package config

import (
	t_rep "eco_points/internal/features/trashes/repository"
	u_rep "eco_points/internal/features/users/repository"
	d_rep "eco_points/internal/features/waste_deposits/repository"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	User        string
	Host        string
	Password    string
	Port        string
	DBName      string
	JWTSecret   string
	CldKey      string
	MidTransKey string
}

func ImportSetting() setting {
	var result setting

	err := godotenv.Load(".env")

	if err != nil {
		return setting{}
	}

	result.User = os.Getenv("DB_USER")
	result.Host = os.Getenv("DB_HOST")
	result.Port = os.Getenv("DB_PORT")
	result.DBName = os.Getenv("DB_NAME")
	result.Password = os.Getenv("DB_PASSWORD")
	result.JWTSecret = os.Getenv("JWT_SECRET")
	result.CldKey = os.Getenv("CLOUDINARY_KEY")
	result.MidTransKey = os.Getenv("MIDTRANS_KEY")
	return result
}

func ConnectDB() (*gorm.DB, error) {
	s := ImportSetting()
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "eco_points_dev.",
		},
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&u_rep.User{}, &t_rep.Trash{}, &d_rep.WasteDeposit{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
