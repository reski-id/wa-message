package data

import (
	"wa/domain"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Message string `json:"message" form:"message"`
	IDUser  string `json:"user_id" form:"user_id"`
}

func (m *Message) ToDomain() domain.Message {
	return domain.Message{
		ID:        int(m.ID),
		Message:   m.Message,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
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
	res.Message = data.Message

	return res
}
