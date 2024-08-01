package handler

import (
	"eco_points/internal/features/excelize"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type excelHandler struct {
	srv excelize.ServiceMakeExcelInterface
	mdl utils.JwtUtilityInterface
}

func NewExcelHandler(s excelize.ServiceMakeExcelInterface, j utils.JwtUtilityInterface) excelize.HandlerMakeExcelInterface {
	return &excelHandler{
		srv: s,
		mdl: j,
	}
}

func (h *excelHandler) MakeExcel() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID uint
		id := h.mdl.DecodToken(c.Get("user").(*jwt.Token))
		if id < 1 {
			log.Println("error from excel")
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautorized", "bad request/invalid jwt", nil)
		}
		limitQ, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			limitQ = 10
		}
		isVerifQ, err := strconv.ParseBool(c.QueryParam("is_verif"))
		if err != nil {
			isVerifQ = false
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
		tableQ := c.QueryParam("table")

		// if tableQ != "deposit" {
		// 	log.Println("error from param", err)
		// 	return helpers.EasyHelper(c, http.StatusBadRequest, "unautorized", "bad request/invalid jwt", nil)
		// }

		var now = time.Now()

		date := fmt.Sprintf("%d%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		link, err := h.srv.MakeExcel(date, tableQ, userID, isVerifQ, uint(limitQ))

		if err != nil {
			log.Println("error get data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}
		return helpers.EasyHelper(c, http.StatusOK, "success", "success make excel", toUrlResponse(date, link))
	}
}
