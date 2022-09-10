package delivery

import "wa/domain"

type MessageResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message" form:"message"`
}

func FromDomain(data domain.Message) MessageResponse {
	var res MessageResponse
	res.Message = data.Message
	return res
}
