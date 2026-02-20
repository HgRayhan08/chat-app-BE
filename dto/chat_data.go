package dto

import "database/sql"

type MessageRequest struct {
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
	RoomId      string `json:"room_id"`
}

type CreateNewMessageRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
type CreateNewMemberRequest struct {
	PhoneNumber string `json:"phone_number"`
	IdRoom      string `json:"room_id"`
}

type GetRoomIdRequest struct {
	RoomId string `json:"room_id"`
}

type RoomChatResponse struct {
	ID        string       `json:"id"`
	RoomId    string       `json:"room_id"`
	Name      string       `json:"name"`
	Phone     string       `json:"phone_number"`
	Active    bool         `json:"is_active"`
	CreatedAt sql.NullTime `json:"created_at"`
}
