package domain

import "database/sql"

// Chat menyimpan pesan user di room tertentu
type Chat struct {
	ID        string       `db:"id"`
	RoomID    string       `db:"room_id"`
	UserID    string       `db:"user_id"`
	Message   string       `db:"message"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// RoomChat menyimpan info room
type RoomChat struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	IsActive  bool         `db:"is_active"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// Repository untuk chat (bisa fleksibel)
type ChatRepository interface {
	Save(chat *Chat) error
	// optional, hanya jika mau load history
	GetByRoomID(roomID string) ([]Chat, error)
}

// Service untuk logic chat
type ChatService interface {
	SendMessage(chat *Chat) error
	// optional, untuk ambil history
	LoadRoomChats(roomID string) ([]Chat, error)
}
