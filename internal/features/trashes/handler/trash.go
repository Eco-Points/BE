package handler

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (h *trashHandler) DeleteTrash() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := h.mdl.DecodToken(c.Get("user").(*jwt.Token))
		log.Println("user ID = ", userID)
		trash_id := c.Param("id")
		id, err := strconv.Atoi(trash_id)
		if err != nil {
			log.Println("error Param", err)
			return helpers.EasyHelper(c, http.StatusBadRequest, "faileds", "bad request/invalid jwt", nil)
		}

		err = h.srv.DeleteTrash(uint(id))
		if err != nil {
			log.Println("error get data", err)
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}
		return helpers.EasyHelper(c, http.StatusOK, "success", "successfully deleted the trash", nil)
	}
}
