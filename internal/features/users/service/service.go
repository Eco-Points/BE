package service

import (
	"eco_points/internal/features/users"
	"eco_points/internal/utils"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	qry users.Query
	pwd utils.PassUtilInterface
	jwt utils.JwtUtilityInterface
}

func NewUserService(q users.Query, p utils.PassUtilInterface, j utils.JwtUtilityInterface) users.Service {
	return &UserServices{
		qry: q,
		pwd: p,
		jwt: j,
	}
}

func (us *UserServices) Register(newUser users.User) error {

	// hashing password
	hashPw, err := us.pwd.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("register generete password error", err.Error())
		return err
	}

	newUser.Password = string(hashPw)

	// register
	err = us.qry.Register(newUser)
	if err != nil {
		log.Println("register sql error:", err.Error())
		return errors.New("error in server")
	}

	return nil
}

func (us *UserServices) Login(email string, password string) (users.User, string, error) {

	result, err := us.qry.Login(email)
	if err != nil {
		log.Println("login sql error: ", err.Error())
		return users.User{}, "", errors.New("error in server")
	}

	// cek password
	err = us.pwd.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Println("Error On Password", err)
		return users.User{}, "", errors.New(bcrypt.ErrMismatchedHashAndPassword.Error())
	}

	// generete token
	token, err := us.jwt.GenereteJwt(result.ID)
	if err != nil {
		log.Println("Error On Jwt ", err)
		return users.User{}, "", errors.New("error on JWT")
	}

	return result, token, nil
}

func (us *UserServices) GetUser(ID uint) (users.User, error) {
	result, err := us.qry.GetUser(ID)

	if err != nil {
		log.Print("get user query error", err.Error())
		return users.User{}, errors.New("internal server error")
	}

	return result, nil
}

func (us *UserServices) UpdateUser(ID uint, updateUser users.User) error {

	if updateUser.Password != "" {
		hashPw, err := us.pwd.GeneratePassword(updateUser.Password)
		if err != nil {
			log.Print("update generete password error", err.Error())
			return errors.New("update generete password error")
		}
		updateUser.Password = string(hashPw)
	}

	// update user
	err := us.qry.UpdateUser(ID, updateUser)

	if err != nil {
		log.Print("update user query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}

func (us *UserServices) DeleteUser(ID uint) error {
	err := us.qry.DeleteUser(ID)

	if err != nil {
		log.Print("delte user query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}
