package service

import (
	"chat-app/domain"
	"chat-app/dto"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type chatService struct {
	chatRepository domain.ChatRepository
	userRepository domain.UserRepository
}

func NewChatService(chatRepository domain.ChatRepository) domain.ChatService {
	return &chatService{
		chatRepository: chatRepository,
	}
}

// CreateRoomMember implements [domain.ChatService].
func (c *chatService) CreateRoomMember(ctx context.Context, f *fiber.Ctx, roomChat dto.CreateNewMemberRequest) (result []domain.RoomMember, err error) {
	userId := f.Locals("user_id").(string)

	send, err := c.userRepository.FindByUserId(ctx, userId)
	if err != nil {
		return []domain.RoomMember{}, err
	}
	recipient, err := c.userRepository.FindByPhoneNumber(ctx, roomChat.PhoneNumber)
	if err != nil {
		return []domain.RoomMember{}, err
	}
	member := []*domain.RoomMember{
		// pengirim
		{
			ID:     uuid.NewString(),
			RoomID: roomChat.IdRoom,
			UserID: send.Id,
		},
		// penerima
		{
			ID:     uuid.NewString(),
			RoomID: roomChat.IdRoom,
			UserID: recipient.Id,
		},
	}
	err = c.chatRepository.SaveRoomMember(ctx, member)
	if err != nil {
		return []domain.RoomMember{}, err
	}
	return []domain.RoomMember{*member[0], *member[1]}, nil

}

// CreateRoomChat implements [domain.ChatService].
func (c *chatService) CreateRoomChat(ctx context.Context, f *fiber.Ctx, roomChat dto.CreateNewMessageRequest) (domain.RoomChat, error) {
	userId := f.Locals("user_id").(string)

	user, err := c.userRepository.FindByUserId(ctx, userId)
	if err != nil {
		return domain.RoomChat{}, err
	}

	checkRoom, err := c.chatRepository.CheckRoomChat(ctx, user.PhoneNumber, roomChat.PhoneNumber)
	if err != nil {
		return domain.RoomChat{}, err
	}
	if checkRoom.ID == "" {
		roomChatId := uuid.NewString()
		err := c.chatRepository.SaveRoomChat(ctx, &domain.RoomChat{
			ID:          roomChatId,
			PhoneNumber: roomChat.PhoneNumber,
			Name:        roomChat.Name,
			IsActive:    true,
		})
		if err != nil {
			return domain.RoomChat{}, err
		}
	}

	room, err := c.chatRepository.GetByRoomID(ctx, checkRoom.ID)
	if err != nil {
		return domain.RoomChat{}, err
	}

	return room, nil
}

// GetChatByRoom implements [domain.ChatService].
func (c *chatService) GetChatByRoom(ctx context.Context, roomID string) ([]domain.Message, error) {
	panic("unimplemented")
}

// LoadAllRoomChats implements [domain.ChatService].
func (c *chatService) LoadAllRoomChats(ctx context.Context, roomID string) ([]domain.Message, error) {
	panic("unimplemented")
}

// SendMessage implements [domain.ChatService].
func (c *chatService) SendMessage(ctx context.Context, chat *domain.Message) error {
	if chat == nil {
		return nil
	}
	return c.chatRepository.SaveChat(ctx, chat)
}
