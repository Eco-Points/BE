package handler

import (
	deposits "eco_points/internal/features/waste_deposits"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type depositHandler struct {
	srv deposits.ServiceInterface
	mdl utils.JwtUtilityInterface
}

func NewDepositService(h deposits.ServiceInterface, m utils.JwtUtilityInterface) deposits.HandlerInterface {
	return &depositHandler{
		srv: h,
		mdl: m,
	}
}

func (h *depositHandler) DepositTrash() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input WasteDeposit
		userID := h.mdl.DecodToken(c.Get("user").(*jwt.Token))
		err := c.Bind(&input)

		if err != nil {
			log.Println("error from Bind", err)
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautorized", "bad request/invalid jwt", nil)
		}

		err = h.srv.DepositTrash(toWasteDepositInterface(input, uint(userID)))
		if err != nil {
			log.Println("error insert data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}

		return helpers.EasyHelper(c, http.StatusCreated, "succes", "waste deposit was successfully created", nil)
	}
}

func (h *depositHandler) UpdateWasteDepositStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input struct {
			WasteID uint   `json:"waste_id"`
			Status  string `json:"status"`
		}
		if err := c.Bind(&input); err != nil {
			log.Println("error from Bind", err)
			return helpers.EasyHelper(c, http.StatusBadRequest, "bad request", "invalid input", nil)
		}

		err := h.srv.UpdateWasteDepositStatus(input.WasteID, input.Status)
		if err != nil {
			log.Println("error updating waste deposit status", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something went wrong with the server", nil)
		}

		return helpers.EasyHelper(c, http.StatusOK, "success", "waste deposit status updated successfully", nil)
	}
}
