package factory

import (
	msg "wa/feature/message/data"
	messageDelivery "wa/feature/message/delivery"
	mu "wa/feature/message/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	//factory
	validator := validator.New()

	msgData := msg.New(db)
	msgCase := mu.New(msgData, validator)
	mercHandler := messageDelivery.New(msgCase)
	messageDelivery.RouteMessage(e, mercHandler)
}
