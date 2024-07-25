package handler

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type trashHandler struct {
	srv trashes.ServiceInterface
	mdl utils.JwtUtilityInterface
}

func NewTrashHandler(s trashes.ServiceInterface, m utils.JwtUtilityInterface) trashes.HandlerInterface {
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
		log.Println("error from bind", err)
		if err != nil {
			log.Println("error insert data", err)
			// return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error database insert data", nil))
			return helpers.EasyHelper(c, http.StatusInternalServerError, "server error", "something wrong with server", nil)
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "trash was successfully created", nil))

	}
}
