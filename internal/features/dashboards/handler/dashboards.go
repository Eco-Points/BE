package handler

import (
	"eco_points/internal/features/dashboards"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	srv dashboards.DshService
	tu  utils.JwtUtilityInterface
}

func NewDashboardHandler(s dashboards.DshService, t utils.JwtUtilityInterface) dashboards.DshHandler {
	return &DashboardHandler{
		srv: s,
		tu:  t,
	}
}

func (uc *DashboardHandler) GetDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		result, err := uc.srv.GetDashboard(uint(userID))
		if err != nil {
			if strings.ContainsAny(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "not allowed", nil))
			}
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "successfully get the deposit history", result))
	}
}

func (uc *DashboardHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		nameParams := c.QueryParam("name")

		result, err := uc.srv.GetAllUsers(uint(userID), nameParams)
		if err != nil {
			if strings.ContainsAny(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "not found", nil))
			}
			if strings.ContainsAny(err.Error(), "not allowed") {
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "not allowed", nil))
			}
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "successfully get all users datas", result))
	}
}
