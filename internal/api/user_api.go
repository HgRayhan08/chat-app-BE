package api

import (
	"chat-app/domain"
	"chat-app/dto"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userApi struct {
	userService domain.UserService
}

func NewUserApi(app *fiber.App, userService domain.UserService, jwtMidd fiber.Handler) {
	api := userApi{userService: userService}
	_ = api
	app.Get("/check-number", jwtMidd, api.GetByPhoneNumber)
}

func (u userApi) GetByPhoneNumber(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.SearchPhoneNumberRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("phone list:", req.PhoneNumber)

	if len(req.PhoneNumber) == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "phone_number tidak boleh kosong",
		})
	}

	result, err := u.userService.GetByPhoneNumber(c, req.PhoneNumber)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseSucsessData(fiber.StatusOK, "Success Get Data", result))
}
