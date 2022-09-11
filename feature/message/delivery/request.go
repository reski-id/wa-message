package delivery

import (
	"time"
	"wa/domain"
)

type MessageInsertRequest struct {
	Message         string `json:"message" form:"message"`
	MessageParentID int    `json:"message_parent_id" form:"message_parent_id"`
	IDUser          int    `json:"user_id" form:"user_id"`
}

func (ci *MessageInsertRequest) ToDomain() domain.Message {
	return domain.Message{
		Message:         ci.Message,
		MessageParentID: ci.MessageParentID,
		IDUser:          ci.IDUser,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
}
