package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello word")
	// Inisialisasi app Fiber
	app := fiber.New()

	// Route sederhana
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "Hello from Go Fiber 🚀",
		})
	})

	// Jalankan server
	log.Fatal(app.Listen(":3000"))
}
