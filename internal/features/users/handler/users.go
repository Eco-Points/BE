package handler

import (
	"eco_points/internal/features/users"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UService
	tu  utils.JwtUtilityInterface
}

func NewUserHandler(s users.UService, t utils.JwtUtilityInterface) users.UHandler {
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
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "user register failed", nil))
		}

		err = uc.srv.Register(RegisterToUser(input))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "user register successful", nil))
	}
}

func (uc *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "user login failed", nil))
		}

		result, token, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user login successful", ToLoginResponse(result, token)))
	}
}

func (uc *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		result, err := uc.srv.GetUser(uint(ID))
		if err != nil {
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "user data not found", nil))
			}
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "All user data fetched successfully", ToGetUserResponse(result)))
	}
}

func (uc *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		var updateUser UpdateRequest
		err := c.Bind(&updateUser)
		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Failed to update account", nil))
		}

		// Bagian updaload Image
		file, err := c.FormFile("image_url")
		if file != nil {
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
			}
		}
		err = uc.srv.UpdateUser(uint(ID), UpdateToUser(updateUser), file)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "Successfully updated account", nil))
	}
}

func (uc *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		err := uc.srv.DeleteUser(uint(ID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Successfully deleted account", nil))
	}
}
