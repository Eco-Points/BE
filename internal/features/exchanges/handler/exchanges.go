package handler

import (
	"eco_points/internal/features/exchanges"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ExchangeHandler struct {
	srv exchanges.ExServices
	tu  utils.JwtUtilityInterface
}

func NewExchangeHandler(s exchanges.ExServices, t utils.JwtUtilityInterface) exchanges.ExHandler {
	return &ExchangeHandler{
		srv: s,
		tu:  t,
	}
}

func (eh *ExchangeHandler) AddExchange() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := eh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		var req CreateExchangeRequest
		if err := c.Bind(&req); err != nil {
			log.Println("invalid request body")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid request body", nil))
		}

		newExchange := exchanges.Exchange{
			RewardID: req.RewardID,
			UserID:   uint(userID),
		}

		err := eh.srv.AddExchange(newExchange)
		if err != nil {
			log.Println("failed to create exchange:", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Exchange was successfully created", nil))
	}
}

func (eh *ExchangeHandler) GetExchangeHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID uint
		id := eh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if id < 1 {
			log.Println("error from excel")
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautorized", "bad request/invalid jwt", nil)
		}
		limitQ, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			limitQ = 10
		}
		isAdminQ, err := strconv.ParseBool(c.QueryParam("is_admin"))
		if err != nil {
			isAdminQ = false
		}
		if isAdminQ {
			userID = 0
		} else {
			userID = uint(id)
		}

		result, err := eh.srv.GetExchangeHistory(userID, isAdminQ, uint(limitQ))

		if err != nil {
			log.Println("error get data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}
		return helpers.EasyHelper(c, http.StatusOK, "success", "success get data exchanges", toListExchange(result))
	}
}
