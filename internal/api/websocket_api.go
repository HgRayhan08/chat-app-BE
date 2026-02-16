package api

import (
	"chat-app/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type websocketAPI struct {
	chatService service.ChatService
}

func NewWebsocketAPI(app *fiber.App, chatService service.ChatService) {
	api := websocketAPI{chatService: chatService}

	app.Get("/ws", websocket.New(api.handleWebSocket))
}

// Handler WebSocket
func (ca websocketAPI) handleWebSocket(ws *websocket.Conn) {
	// c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer ws.Close()
	for {
		// Baca pesan dari client
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		log.Println("Received:", string(msg))

		// Kirim balik ke client (echo)
		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
