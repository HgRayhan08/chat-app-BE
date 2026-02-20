package repository

import (
	"chat-app/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type ChatRepository struct {
	db *goqu.Database
}

// GetByRoomID implements [domain.ChatRepository].

func NewChatRepository(db *sql.DB) domain.ChatRepository {
	return &ChatRepository{
		db: goqu.New("postgres", db),
	}
}

// CheckRoomChat implements [domain.ChatRepository].
func (c *ChatRepository) CheckRoomChat(ctx context.Context, user1 string, user2 string) (result domain.RoomChat, err error) {
	dataset := c.db.From("rooms").
		Join(goqu.T("room_members").As("rm1"),
			goqu.On(goqu.I("rm1.room_id").Eq(goqu.I("rooms.id")))).
		Join(goqu.T("room_members").As("rm2"),
			goqu.On(goqu.I("rm2.room_id").Eq(goqu.I("rooms.id")))).
		Where(goqu.I("rm1.user_id").Eq(user1)).
		Where(goqu.I("rm2.user_id").Eq(user2)).
		Select(goqu.I("rooms.id"))
	_, err = dataset.ScanValContext(ctx, &result)
	return
}

func (c *ChatRepository) GetByRoomID(ctx context.Context, roomID string) (result domain.RoomChat, err error) {
	dataset := c.db.From("room_chats").Where(goqu.C("room_id").Eq(roomID))
	_, err = dataset.ScanStructContext(ctx, &result)
	return

}

// SaveRoomMember implements [domain.ChatRepository].
func (c *ChatRepository) SaveRoomMember(ctx context.Context, member []*domain.RoomMember) error {
	dataset := c.db.Insert("room_members").Rows(member).Executor()
	_, err := dataset.ExecContext(ctx)

	return err
}

// SaveChat implements [domain.ChatRepository].
func (c *ChatRepository) SaveChat(ctx context.Context, message *domain.Message) error {
	dataset := c.db.Insert("message").Rows(message).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// SaveRoomChat implements [domain.ChatRepository].
func (c *ChatRepository) SaveRoomChat(ctx context.Context, room *domain.RoomChat) error {
	dataset := c.db.Insert("room_chats").Rows(room).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// GetHistoryChat implements [domain.ChatRepository].
func (c *ChatRepository) GetHistoryChat(ctx context.Context, roomID string, userId string) (result []domain.Message, err error) {
	dataset := c.db.From("message").Where(goqu.C("room_id").Eq(roomID)).Where(goqu.C("user_id").Eq(userId)).Order(goqu.C("created_at").Asc()).Limit(50)
	err = dataset.ScanStructsContext(ctx, &result)
	return
}
