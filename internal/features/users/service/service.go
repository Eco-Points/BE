package service

import (
	"eco_points/internal/features/users"
	"eco_points/internal/utils"
	"errors"
	"log"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
	qry   users.UQuery
	pwd   utils.PassUtilInterface
	jwt   utils.JwtUtilityInterface
	cloud utils.CloudinaryUtilityInterface
}

func NewUserService(q users.UQuery, p utils.PassUtilInterface, j utils.JwtUtilityInterface, c utils.CloudinaryUtilityInterface) users.UService {
	return &UserServices{
		qry:   q,
		pwd:   p,
		jwt:   j,
		cloud: c,
	}
}

func (us *UserServices) Register(newUser users.User) error {

	// hashing password
	hashPw, err := us.pwd.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("register generate password error", err.Error())
		return err
	}
	newUser.IsAdmin = false
	newUser.Password = string(hashPw)
	newUser.Status = "active"

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
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return users.User{}, err
		}
		return users.User{}, errors.New("error in query")
	}

	return result, nil
}

func (us *UserServices) UpdateUser(ID uint, updateUser users.User, file *multipart.FileHeader) error {

	if updateUser.Password != "" {
		hashPw, err := us.pwd.GeneratePassword(updateUser.Password)
		if err != nil {
			log.Print("update generete password error", err.Error())
			return errors.New("update generete password error")
		}
		updateUser.Password = string(hashPw)
	}
	if file != nil {
		src, err := us.cloud.FileCheck(file)
		if err != nil {
			return errors.New("file Error")
		}

		urlImage, err := us.cloud.UploadToCloudinary(src, file.Filename)
		if err != nil {
			return errors.New("upload error")
		}
		updateUser.ImgURL = urlImage
	}
	// update user
	err := us.qry.UpdateUser(ID, updateUser)

	if err != nil {
		log.Print("update user query error", err.Error())
		return errors.New("query error")
	}

	return nil
}

func (us *UserServices) DeleteUser(ID uint) error {
	err := us.qry.DeleteUser(ID)

	if err != nil {
		log.Print("delte user query error", err.Error())
		return errors.New("query error")
	}

	return nil
}
