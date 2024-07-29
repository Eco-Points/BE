package service_test

import (
	"eco_points/internal/features/users"
	"eco_points/internal/features/users/service"
	"eco_points/mocks"
	"errors"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	qry := mocks.NewUQuery(t)
	psw := mocks.NewPassUtilInterface(t)
	jw := mocks.NewJwtUtilityInterface(t)
	cld := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewUserService(qry, psw, jw, cld)

	RegisterRequest := users.User{
		Fullname: "Hafiz",
		Email:    "hafiz@gmail.com",
		Password: "hashedPassword",
		IsAdmin:  false,
		Status:   "active",
	}
	t.Run("success register", func(t *testing.T) {
		psw.On("GeneratePassword", "hashedPassword").Return([]byte("hashedPassword"), nil).Once()
		qry.On("Register", RegisterRequest).Return(nil).Once()
		err := srv.Register(RegisterRequest)
		assert.Nil(t, err)
	})

	t.Run("fail register on password hashing", func(t *testing.T) {

		expectedError := errors.New("register generate password error")
		psw.On("GeneratePassword", "hashedPassword").Return([]byte("another password"), expectedError).Once()

		err := srv.Register(RegisterRequest)
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
	})

	t.Run("fail register on query error", func(t *testing.T) {
		expectedError := errors.New("error in server")
		psw.On("GeneratePassword", "hashedPassword").Return([]byte("hashedPassword"), nil).Once()
		qry.On("Register", RegisterRequest).Return(expectedError).Once()

		err := srv.Register(RegisterRequest)
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
	})
}

func TestLogin(t *testing.T) {
	qry := mocks.NewUQuery(t)
	psw := mocks.NewPassUtilInterface(t)
	jw := mocks.NewJwtUtilityInterface(t)
	cld := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewUserService(qry, psw, jw, cld)

	LoginRequest := users.User{
		Email:    "hafiz@gmail.com",
		Password: "password",
	}
	Result := users.User{
		Email:    "hafiz@gmail.com",
		Password: "hashedPassword",
	}
	token := "token_string"
	t.Run("success login", func(t *testing.T) {
		qry.On("Login", LoginRequest.Email).Return(Result, nil).Once()
		psw.On("ComparePassword", []byte(Result.Password), []byte(LoginRequest.Password)).Return(nil).Once()
		jw.On("GenereteJwt", Result.ID).Return(token, nil).Once()

		Result, token, err := srv.Login(LoginRequest.Email, LoginRequest.Password)
		assert.Nil(t, err)
		assert.NotNil(t, Result, token)
	})

	t.Run("fail login user not found", func(t *testing.T) {
		expectedError := errors.New("error in server")
		qry.On("Login", LoginRequest.Email).Return(users.User{}, expectedError).Once()

		Result, token, err := srv.Login(LoginRequest.Email, LoginRequest.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, users.User{}, Result)
		assert.Equal(t, "", token)
	})

	t.Run("fail login password mismatch", func(t *testing.T) {
		expectedError := errors.New(bcrypt.ErrMismatchedHashAndPassword.Error())
		qry.On("Login", LoginRequest.Email).Return(Result, nil).Once()
		psw.On("ComparePassword", []byte(Result.Password), []byte(LoginRequest.Password)).Return(expectedError).Once()

		Result, token, err := srv.Login(LoginRequest.Email, LoginRequest.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, users.User{}, Result)
		assert.Equal(t, "", token)
	})

	t.Run("fail login token error", func(t *testing.T) {
		expectedError := errors.New("error on JWT")
		qry.On("Login", LoginRequest.Email).Return(Result, nil).Once()
		psw.On("ComparePassword", []byte(Result.Password), []byte(LoginRequest.Password)).Return(nil).Once()
		jw.On("GenereteJwt", Result.ID).Return("", expectedError).Once()

		Result, token, err := srv.Login(LoginRequest.Email, LoginRequest.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, users.User{}, Result)
		assert.Equal(t, "", token)
	})
}

func TestGetUser(t *testing.T) {
	qry := mocks.NewUQuery(t)
	psw := mocks.NewPassUtilInterface(t)
	jw := mocks.NewJwtUtilityInterface(t)
	cld := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewUserService(qry, psw, jw, cld)

	result := users.User{
		Email: "hafiz@gmail.com",
	}

	t.Run("success get user data", func(t *testing.T) {
		qry.On("GetUser", uint(1)).Return(result, nil).Once()

		result, err := srv.GetUser(uint(1))
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("fail get user data on query", func(t *testing.T) {
		expectedError := errors.New("error in query")
		qry.On("GetUser", uint(1)).Return(users.User{}, expectedError).Once()

		result, err := srv.GetUser(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, expectedError, err)
		assert.Equal(t, users.User{}, result)
	})

	t.Run("fail get user data not found", func(t *testing.T) {
		expectedError := errors.New(gorm.ErrRecordNotFound.Error())
		qry.On("GetUser", uint(1)).Return(users.User{}, expectedError).Once()

		result, err := srv.GetUser(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, expectedError, err.Error())
		assert.Equal(t, expectedError, err)
		assert.Equal(t, users.User{}, result)
	})
}

func TestUpdateUser(t *testing.T) {
	qry := mocks.NewUQuery(t)
	psw := mocks.NewPassUtilInterface(t)
	jw := mocks.NewJwtUtilityInterface(t)
	cld := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewUserService(qry, psw, jw, cld)

	input := users.User{
		Email:    "hafiz@gmail.com",
		Password: "password",
	}
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="test.png"`)
	h.Set("Content-Type", "image/jpg")

	file := &multipart.FileHeader{
		Filename: "test.jpg",
		Header:   h,
		Size:     123,
	}

	src, _ := file.Open()
	defer src.Close()

	var userID uint = 1
	t.Run("fail update on password", func(t *testing.T) {
		expectedError := errors.New("update generete password error")
		psw.On("GeneratePassword", input.Password).Return([]byte("anotherPassword"), expectedError).Once()

		err := srv.UpdateUser(userID, input, file)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("fail update on file error", func(t *testing.T) {
		expectedError := errors.New("file Error")
		psw.On("GeneratePassword", input.Password).Return([]byte("password"), nil).Once()
		cld.On("FileCheck", file).Return(src, expectedError).Once()

		err := srv.UpdateUser(userID, input, file)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("fail update on file upload", func(t *testing.T) {
		expectedError := errors.New("upload error")
		psw.On("GeneratePassword", input.Password).Return([]byte("password"), nil).Once()
		cld.On("FileCheck", file).Return(src, nil).Once()
		cld.On("UploadToCloudinary", src, file.Filename).Return("", expectedError).Once()

		err := srv.UpdateUser(userID, input, file)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("fail update on update query", func(t *testing.T) {
		expectedError := errors.New("query error")
		psw.On("GeneratePassword", input.Password).Return([]byte("password"), nil).Once()
		cld.On("FileCheck", file).Return(src, nil).Once()
		cld.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
		qry.On("UpdateUser", userID, input).Return(expectedError).Once()

		err := srv.UpdateUser(userID, input, file)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("success update user", func(t *testing.T) {

		psw.On("GeneratePassword", input.Password).Return([]byte("password"), nil).Once()
		cld.On("FileCheck", file).Return(src, nil).Once()
		cld.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
		qry.On("UpdateUser", userID, input).Return(nil).Once()

		err := srv.UpdateUser(userID, input, file)
		assert.Nil(t, err)

	})
}

func TestDeleteUser(t *testing.T) {
	qry := mocks.NewUQuery(t)
	psw := mocks.NewPassUtilInterface(t)
	jw := mocks.NewJwtUtilityInterface(t)
	cld := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewUserService(qry, psw, jw, cld)

	var userID uint = 1
	t.Run("fail delete on delete query", func(t *testing.T) {
		expectedError := errors.New("query error")
		qry.On("DeleteUser", userID).Return(expectedError).Once()

		err := srv.DeleteUser(userID)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)

	})
	t.Run("success delete user", func(t *testing.T) {
		qry.On("DeleteUser", userID).Return(nil).Once()

		err := srv.DeleteUser(userID)
		assert.Nil(t, err)

	})
}
