package api

import (
	"chat-app/domain"
	"chat-app/dto"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuthApi(app *fiber.App, authService domain.AuthService, jwtMiddle fiber.Handler) {
	auth := authApi{authService: authService}
	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	app.Post("/refresh-token", auth.RefreshToken)
	app.Post("/logout", jwtMiddle, auth.Logout)
}

func (a *authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	err := ctx.BodyParser(&req)

	if err != nil {
		return err
	}

	result, err := a.authService.Login(c, req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)

}

func (a *authApi) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	if req.Email == "" || req.Password == "" || req.PhoneNumber == "" || req.Username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "All fields are required: email, password, phone number, and username",
		})
	}
	if len(req.PhoneNumber) <= 13 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone number must be at least 13 characters",
		})
	}

	result, err := a.authService.Register(c, req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (a *authApi) RefreshToken(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TokenRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	refreshToken, err := a.authService.RefreshToken(c, req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.RefreshTokenResponse{
		Code:    fiber.StatusOK,
		Message: "Success Refresh Token",
		Token:   refreshToken.Token,
	})

}

func (a *authApi) Logout(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	err := a.authService.Logout(c, ctx)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.AuthResponse{
		Code:    fiber.StatusOK,
		Message: "Success Logout",
	})
}
