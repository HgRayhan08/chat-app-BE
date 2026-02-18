package domain

import (
	"chat-app/dto"
	"context"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
	Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
	RefreshToken(ctx context.Context, req dto.TokenRequest) (dto.RefreshTokenResponse, error)
	Logout(ctx context.Context, f *fiber.Ctx) error
}
