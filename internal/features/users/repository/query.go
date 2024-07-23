package repository

import (
	"eco_points/internal/features/users"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(connect *gorm.DB) users.Query {
	return &UserQuery{
		db: connect,
	}
}

// Register
func (um *UserQuery) Register(newUser users.User) error {
	cnvData := ToUserQuery(newUser)
	err := um.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

// Login
func (um *UserQuery) Login(email string) (users.User, error) {
	var result User
	err := um.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.User{}, err
	}

	return result.ToUserEntity(), nil
}

// Get one data user
func (um *UserQuery) GetUser(ID uint) (users.User, error) {
	var result User
	err := um.db.Where("id = ?", ID).First(&result).Error

	if err != nil {
		return users.User{}, err
	}

	return result.ToUserEntity(), nil
}

// update data user
func (um *UserQuery) UpdateUser(ID uint, updateUser users.User) error {
	cnvQuery := ToUserQuery(updateUser)
	qry := um.db.Where("id = ?", ID).Updates(&cnvQuery)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (um *UserQuery) DeleteUser(ID uint) error {
	qry := um.db.Where("id = ?", ID).Delete(&User{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
