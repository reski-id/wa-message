package delivery

import "wa/domain"

type MessageResponse struct {
	// Message string `json:"message" form:"message"`

	IDUser          int    `json:"user_id" form:"user_id"`
	MessageParentID int    `json:"message_parent_id" form:"message_parent_id"`
	Message         string `json:"message" form:"message"`
}

func FromDomain(data domain.Message) MessageResponse {
	var res MessageResponse
	res.IDUser = data.IDUser
	res.MessageParentID = data.MessageParentID
	res.Message = data.Message
	return res
}
