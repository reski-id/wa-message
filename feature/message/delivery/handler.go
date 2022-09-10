package delivery

import (
	"log"
	"net/http"
	"wa/domain"

	"github.com/labstack/echo/v4"
)

type msgHandler struct {
	mercUsecase domain.MessageUseCase
}

func New(cu domain.MessageUseCase) domain.MessageHandler {
	return &msgHandler{
		mercUsecase: cu,
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

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
		})

	}
}
