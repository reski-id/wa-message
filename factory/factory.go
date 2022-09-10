package factory

import (
	msg "wa/feature/message/data"
	messageDelivery "wa/feature/message/delivery"
	mu "wa/feature/message/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "wa/feature/user/data"
	userDelivery "wa/feature/user/delivery"
	us "wa/feature/user/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	//factory
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	msgData := msg.New(db)
	msgCase := mu.New(msgData, validator)
	mercHandler := messageDelivery.New(msgCase)
	messageDelivery.RouteMessage(e, mercHandler)
}
