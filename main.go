package main

import (
	"chat-app/internal/api"
	"chat-app/internal/config"
	"chat-app/internal/connection"
	"chat-app/internal/middleware"
	"chat-app/internal/repository"
	"chat-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Get()

	database := connection.GetDatabaseConnection(config.Database)

	//memanggil agar bisa di run
	// _ = database

	// Inisialisasi app Fiber
	app := fiber.New()

	jwtMiddle := middleware.JWTProtected(config)

	// Inisialisasi repository
	userRepository :=
		repository.NewUserRepository(database)

		// Inisialisasi service
	AuthService := service.NewAuthService(config, userRepository)
	chatService := service.ChatService{}

	// Inisialisasi API
	api.NewWebsocketAPI(app, chatService)
	api.NewAuthApi(app, AuthService, jwtMiddle)

	// // Route sederhana
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{})
	// })
	app.Listen(":9000")
	// Jalankan server
	// log.Fatal(app.Listen(":3000"))
}
