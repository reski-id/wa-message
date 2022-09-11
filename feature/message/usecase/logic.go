package usecase

import (
	"errors"
	"fmt"
	"wa/domain"

	validator "github.com/go-playground/validator/v10"
)

type msgUseCase struct {
	msgData  domain.MessageData
	validate *validator.Validate
}

func New(ud domain.MessageData, v *validator.Validate) domain.MessageUseCase {
	return &msgUseCase{
		msgData:  ud,
		validate: v,
	}
}

func (cc *msgUseCase) AddMessage(IDUser int, newMessage domain.Message) (domain.Message, error) {
	if IDUser == -1 {
		return domain.Message{}, errors.New("invalid user")
	}

	newMessage.IDUser = IDUser
	fmt.Println("logic", IDUser)

	res := cc.msgData.Insert(newMessage)
	if res.ID == 0 {
		return domain.Message{}, errors.New("error insert data")
	}
	return res, nil
}

func (cc *msgUseCase) DelMessage(IDMessage int) (bool, error) {
	res := cc.msgData.Delete(IDMessage)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}
