package delivery

import (
	"log"
	"net/http"
	"strconv"
	"wa/domain"
	"wa/feature/common"

	"github.com/labstack/echo/v4"
)

type msgHandler struct {
	msgUsecase domain.MessageUseCase
}

func New(cu domain.MessageUseCase) domain.MessageHandler {
	return &msgHandler{
		msgUsecase: cu,
	}
}

func (ch *msgHandler) InsertMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp MessageInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		var IDUser = common.ExtractData(c)
		data, err := ch.msgUsecase.AddMessage(IDUser, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (ch *msgHandler) DeleteMessage() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := ch.msgUsecase.DelMessage(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data",
		})
	}
}
