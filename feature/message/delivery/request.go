package delivery

import (
	"time"
	"wa/domain"
)

type MessageInsertRequest struct {
	Message string `json:"message" form:"message"`
}

func (ci *MessageInsertRequest) ToDomain() domain.Message {
	return domain.Message{
		Message:   ci.Message,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
