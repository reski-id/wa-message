package data

import (
	"fmt"
	"wa/domain"

	"gorm.io/gorm"
)

type messageData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.MessageData {
	return &messageData{
		db: db,
	}
}

func (cd *messageData) Insert(newmessage domain.Message) domain.Message {
	cnv := ToLocal(newmessage)
	err := cd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Message{}
	}
	return cnv.ToDomain()
}
