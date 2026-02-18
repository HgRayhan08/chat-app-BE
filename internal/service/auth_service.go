package service

import (
	"chat-app/domain"
	"chat-app/dto"
	"chat-app/internal/config"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuthService(conf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &AuthService{
		conf:           conf,
		userRepository: userRepository,
	}
}

// Login implements [domain.AuthService].
func (a *AuthService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	fmt.Println("emailnya", req.Email)
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	fmt.Println("ini user", user)
	if err != nil && user.Id == "" {
		return dto.AuthResponse{}, errors.New("Email Tidak terdaftar, silakan Registrasi")
	}
	fmt.Println("user", user)
	refreshToken := uuid.NewString()

	// chek refresh token
	oldToken := a.userRepository.FindByuserIdRefreshToken(ctx, user.Id)
	if oldToken != nil {
		fmt.Println("ada refresh token")
		err = a.userRepository.DeleteRefreshToken(ctx, user.Id)
		fmt.Println("berhasil di hapus")
		if err != nil {
			return dto.AuthResponse{}, errors.New("gagal menghapus refresh token")
		}
	}

	// create refresh token
	err = a.userRepository.SaveRefreshToken(ctx, &domain.RefreshToken{
		Id:        uuid.NewString(),
		Token:     refreshToken,
		UserId:    user.Id,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	})
	if err != nil {
		return dto.AuthResponse{}, errors.New("gagal membuat refresh token")
	}

	// 5. Generate JWT
	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.JWT.Exp) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(a.conf.JWT.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("gagal generate jwt")
	}

	// 6. Response
	return dto.AuthResponse{
		Code:    200,
		Message: "Success Login",
		User: dto.User{
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Username:    user.Username,
		},
		TokenRefresh: refreshToken,
		TokenJwt:     tokenString,
	}, nil
}

// Register implements [domain.AuthService].
func (a *AuthService) Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	cekEmail, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err == nil && cekEmail.Id != "" {
		return dto.AuthResponse{}, errors.New("Email sudah terdaftar, silakan login")
	}

	generatePass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return dto.AuthResponse{}, errors.New("Failed Generate Password")
	}
	user := domain.User{
		Id:          uuid.NewString(),
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    string(generatePass),
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}
	err = a.userRepository.Save(ctx, &user)
	fmt.Println(err)
	if err != nil {
		fmt.Println("terjadi error", err)
		return dto.AuthResponse{}, errors.New("Failed Register")
	}
	return dto.AuthResponse{
		Code:    200,
		Message: "Success Register",
		User:    dto.User{Email: user.Email, PhoneNumber: user.PhoneNumber, Username: user.Username},
	}, nil

}

// RefreshToken implements [domain.AuthService].
func (a *AuthService) RefreshToken(ctx context.Context, req dto.TokenRequest) (result dto.RefreshTokenResponse, err error) {
	refresh, err := a.userRepository.FindByResfreshToken(ctx, req.Token)

	if err != nil {
		return dto.RefreshTokenResponse{}, errors.New("Failed Refresh Token")
	}

	if refresh.Id == "" {
		return dto.RefreshTokenResponse{}, errors.New("Failed Refresh Token")
	}
	fmt.Println(refresh.UserId)
	claim := jwt.MapClaims{
		"id":  refresh.UserId,
		"exp": time.Now().Add(time.Duration(a.conf.JWT.Exp) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(a.conf.JWT.Key))
	if err != nil {
		return dto.RefreshTokenResponse{}, errors.New("Failed Generate Token")
	}

	return dto.RefreshTokenResponse{
		Code:    200,
		Message: "Success Refresh Token",
		Token:   tokenString,
	}, nil
}

// Logout implements [domain.AuthService].
func (a *AuthService) Logout(ctx context.Context, f *fiber.Ctx) error {
	userId := f.Locals("user_id").(string)
	fmt.Println("ini user id", userId)
	err := a.userRepository.DeleteRefreshToken(ctx, userId)
	fmt.Println(err)
	if err != nil {
		return errors.New("Failed Logout")
	}
	return nil
}
