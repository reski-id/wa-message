package data

import (
	"fmt"
	"log"
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

func (cd *messageData) Delete(msgID int) bool {
	err := cd.db.Where("ID = ?", msgID).Delete(&Message{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}
