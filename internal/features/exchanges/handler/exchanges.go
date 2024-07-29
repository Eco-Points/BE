package handler

import (
	"eco_points/internal/features/exchanges"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"

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
