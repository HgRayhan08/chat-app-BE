package main

import (
	"chat-app/internal/api"
	"chat-app/internal/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello word")
	// Inisialisasi app Fiber
	app := fiber.New()

	chatService := service.ChatService{}

	api.NewWebsocketAPI(app, chatService)

	// // Route sederhana
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{})
	// })

	// Jalankan server
	log.Fatal(app.Listen(":3000"))
}
