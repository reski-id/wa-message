package data

import (
	"wa/domain"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Message         string `json:"message" form:"message"`
	MessageParentID int    `json:"message_parent_id" form:"message_parent_id"`
	IDUser          int    `json:"user_id" form:"user_id"`
}

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname"`
}

func (m *Message) ToDomain() domain.Message {
	return domain.Message{
		ID:              int(m.ID),
		Message:         m.Message,
		IDUser:          m.IDUser,
		MessageParentID: m.MessageParentID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func ParseToArr(arr []Message) []domain.Message {
	var res []domain.Message

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Message) Message {
	var res Message
	res.ID = uint(data.ID)
	res.MessageParentID = data.MessageParentID
	res.IDUser = data.IDUser
	res.Message = data.Message
	return res
}
