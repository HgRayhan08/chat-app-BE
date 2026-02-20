package domain

import (
	"chat-app/dto"
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// Chat menyimpan pesan user di room tertentu
type Message struct {
	ID        string       `db:"id"`
	RoomID    string       `db:"room_id"`
	UserID    string       `db:"user_id"`
	Message   string       `db:"message"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// RoomChat menyimpan info room
type RoomChat struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	PhoneNumber string       `db:"phone_number"`
	IsActive    bool         `db:"is_active"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type RoomMember struct {
	ID        string       `db:"id"`
	RoomID    string       `db:"room_id"`
	UserID    string       `db:"user_id"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// Repository untuk chat (bisa fleksibel)
type ChatRepository interface {
	// table chat
	SaveChat(ctx context.Context, chat *Message) error
	GetHistoryChat(ctx context.Context, roomID string, userId string) ([]Message, error)
	// table room member
	SaveRoomMember(ctx context.Context, member []*RoomMember) error
	// table room chat
	SaveRoomChat(ctx context.Context, room *RoomChat) error
	CheckRoomChat(ctx context.Context, user1 string, user2 string) (RoomChat, error)
	GetByRoomID(ctx context.Context, roomID string) (RoomChat, error)
}

// Service untuk logic chat
type ChatService interface {
	SendMessage(ctx context.Context, chat *Message) error
	CreateRoomMember(ctx context.Context, f *fiber.Ctx, roomChat dto.CreateNewMemberRequest) ([]RoomMember, error)
	// GetChatByRoom(ctx context.Context, roomID string) ([]Message, error)
	// room chat
	CreateRoomChat(ctx context.Context, f *fiber.Ctx, roomChat dto.CreateNewMessageRequest) (RoomChat, error)
	// optional, untuk ambil history
	LoadAllRoomChats(ctx context.Context, roomID string) ([]Message, error)
}
