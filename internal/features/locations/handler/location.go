package handler

import (
	"eco_points/internal/features/locations"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type locHandler struct {
	mdl utils.JwtUtilityInterface
	srv locations.ServiceLocInterface
}

func NewLocHandler(s locations.ServiceLocInterface, m utils.JwtUtilityInterface) locations.HandlerLocInterface {
	return &locHandler{
		mdl: m,
		srv: s,
	}
}

func (h *locHandler) AddLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Location
		userID := h.mdl.DecodToken(c.Get("user").(*jwt.Token))
		err := c.Bind(&input)

		if err != nil {
			log.Println("error from Bind", err)
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautorized", "bad request/invalid jwt", nil)
		}

		err = h.srv.AddLocation(toLocInterface(input, uint(userID)))
		if err != nil {
			log.Println("error insert data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}

		return helpers.EasyHelper(c, http.StatusCreated, "succes", "location was successfully created", nil)
	}
}

func (h *locHandler) GetLocation() echo.HandlerFunc {
	return func(c echo.Context) error {

		jwtUserID, err := h.mdl.DecodTokenV2(c)

		if err != nil {
			log.Println("error from jwt", err)
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautoriezd", "bad request/invalid jwt", nil)
		}

		if jwtUserID < 1 {
			log.Println("user ID = ", jwtUserID)
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautoriezd", "bad request/invalid jwt", nil)
		}

		fmt.Println("param")

		paramUserID := c.Param("id")
		fmt.Println("param")

		userID, err := strconv.Atoi(paramUserID)
		fmt.Println("atoi")

		if err != nil {
			userID = 0
		}
		fmt.Println("srv")

		result, err := h.srv.GetLocation(uint(userID))
		if err != nil {
			log.Println("error insert data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}

		return helpers.EasyHelper(c, http.StatusOK, "succes", "successfully get location", toResposne(result))

	}
}
