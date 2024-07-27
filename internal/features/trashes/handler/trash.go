package handler

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type trashHandler struct {
	srv trashes.ServiceTrashInterface
	mdl utils.JwtUtilityInterface
}

func NewTrashHandler(s trashes.ServiceTrashInterface, m utils.JwtUtilityInterface) trashes.HandlerTrashInterface {
	return &trashHandler{
		srv: s,
		mdl: m,
	}
}

func (h *trashHandler) AddTrash() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Trash
		userID := h.mdl.DecodToken(c.Get("user").(*jwt.Token))
		log.Println("user ID = ", userID)
		err := c.Bind(&input)
		if err != nil {
			log.Println("error from bind", err)
			// return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "error on request data", nil))
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautoriezd", "bad request/invalid jwt", nil)
		}

		file, err := c.FormFile("images")
		if err != nil {
			log.Println("error when get image", err)
			return helpers.EasyHelper(c, http.StatusBadRequest, "unautoriezd", "bad request/invalid jwt", nil)
		}

		err = h.srv.AddTrash(toTrashEntity(input, uint(userID)), file)
		if err != nil {
			log.Println("error insert data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "trash was successfully created", nil))

	}
}

func (h *trashHandler) GetTrash() echo.HandlerFunc {
	return func(c echo.Context) error {
		ttype := c.QueryParam("type")
		point := c.QueryParam("point")
		fmt.Println(ttype)
		fmt.Println(point)

		result, err := h.srv.GetTrash(ttype)
		if err != nil {
			log.Println("error get data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}
		return helpers.EasyHelper(c, http.StatusOK, "success", "successfully get the trash", toListTrashResponse(result))
	}
}
