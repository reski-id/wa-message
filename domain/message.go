package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Message struct {
	//struct
	ID        int
	Message   string
	IDUser    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageUseCase interface {
	//usecase
	AddMessage(IDUser int, useMessage Message) (Message, error)
}

type MessageHandler interface {
	//handler
	InsertMessage() echo.HandlerFunc
}

type MessageData interface {
	//data
	Insert(insertMessage Message) Message
}
