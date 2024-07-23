package handler

import (
	"eco_points/internal/features/users"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.Service
	tu  utils.JwtUtilityInterface
}

func NewUserHandler(s users.Service, t utils.JwtUtilityInterface) users.Handler {
	return &UserHandler{
		srv: s,
		tu:  t,
	}
}

func (uc *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}

		err = uc.srv.Register(RegisterToUser(input))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "register success", nil))
	}
}

func (uc *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}

		_, token, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "user logged in", ToLoginResponse(token)))
	}
}

func (uc *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		log.Print("================================> id = ", ID)
		result, err := uc.srv.GetUser(uint(ID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "User successfully retrieved", ToGetUserResponse(result)))
	}
}

func (uc *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		var updateUser UpdateRequest
		err := c.Bind(&updateUser)
		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "Invalid request parameters", nil))
		}

		// Bagian updaload Image
		file, err := c.FormFile("image_picture")

		if err == nil {
			// open file
			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "profile image error", nil))
			}
			defer src.Close()

			urlImage, err := utils.UploadToCloudinary(src, file.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "profile image error", nil))
			}
			updateUser.ImgURL = urlImage

		}

		err = uc.srv.UpdateUser(uint(ID), UpdateToUser(updateUser))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "User profile updated", nil))
	}
}

func (uc *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		err := uc.srv.DeleteUser(uint(ID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "User account deleted", nil))
	}
}
