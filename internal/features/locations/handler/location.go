package handler

import (
	"eco_points/internal/features/locations"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type locHandler struct {
	mdl utils.JwtUtilityInterface
	srv locations.ServiceInterface
}

func NewLocHandler(s locations.ServiceInterface, m utils.JwtUtilityInterface) locations.HandlerInterface {
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
