package utils

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	rooms map[string]map[*websocket.Conn]string
	mu    sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]map[*websocket.Conn]string),
	}
}

// user join ke room tertentu sesuai dengan id room yang sudah dibuat
func (h *Hub) Join(roomID string, conn *websocket.Conn, userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.rooms[roomID] == nil {
		h.rooms[roomID] = make(map[*websocket.Conn]string)
	}
	h.rooms[roomID][conn] = userID
}

func (h *Hub) Leave(roomID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.rooms[roomID], conn)
	if len(h.rooms[roomID]) == 0 {
		delete(h.rooms, roomID)
	}
}

func (h *Hub) Broadcast(roomID string, message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for conn := range h.rooms[roomID] {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
