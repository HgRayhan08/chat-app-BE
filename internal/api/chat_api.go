package api

import (
	"chat-app/domain"
	"chat-app/dto"
	"chat-app/internal/utils"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type chatApi struct {
	chatService domain.ChatService
	hub         *utils.Hub
}

func NewWebsocketAPI(app *fiber.App, chatService domain.ChatService, jwtMiddle fiber.Handler) {
	api := chatApi{chatService: chatService, hub: utils.NewHub()}
	// socket
	// contoh url: ws://localhost:9000/ws?room_id=abc123&user_id=xyz456
	app.Get("/ws", websocket.New(api.handleWebSocket))
	// rest-api
	app.Post("/new-chat", jwtMiddle, api.CreateNewChat)
	app.Post("/new-member", jwtMiddle, api.CreateNewMember)
	app.Get("/all-room", jwtMiddle, api.GetAllRoomChat)
	app.Get("/all-message", jwtMiddle, api.GetAllMessageByRoom)
}

func (ca chatApi) handleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	roomID := ws.Query("room_id")
	userID := ws.Query("user_id")

	if roomID == "" || userID == "" {
		log.Println("room_id / user_id kosong")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// join ke room
	ca.hub.Join(roomID, ws, userID)
	defer ca.hub.Leave(roomID, ws)

	log.Printf("User %s join room %s\n", userID, roomID)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// simpan ke DB
		err = ca.chatService.SendMessage(ctx, &domain.Message{
			ID:        uuid.NewString(),
			RoomID:    roomID,
			UserID:    userID,
			Message:   string(msg),
			CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
			UpdatedAt: sql.NullTime{Valid: false},
		})
		if err != nil {
			log.Println("save message error:", err)
			continue
		}

		// broadcast ke semua di room
		ca.hub.Broadcast(roomID, msg)
	}
}

func (ca *chatApi) CreateNewChat(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateNewMessageRequest

	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	result, err := ca.chatService.CreateRoomChat(c, ctx, req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseSucsessData(fiber.StatusOK, "Sucses", result))
}

func (ca *chatApi) CreateNewMember(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateNewMemberRequest

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			dto.ResponseError(
				fiber.StatusBadRequest,
				err.Error(),
			),
		)
	}
	result, err := ca.chatService.CreateRoomMember(c, ctx, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError(fiber.StatusBadRequest, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseSucsessData(fiber.StatusOK, "Sucses", result))
}

func (ca *chatApi) GetAllRoomChat(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	allRoom, err := ca.chatService.LoadAllRoomChats(c, ctx)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseSucsessData(fiber.StatusOK, "Success Get Data", allRoom))
}

func (ca *chatApi) GetAllMessageByRoom(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.GetRoomIdRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	result, err := ca.chatService.GetMessageByRoom(c, req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseSucsessData(fiber.StatusOK, "Success Get Data", result))
}
