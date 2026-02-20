package domain

import (
	"context"
	"database/sql"
)

type User struct {
	Id          string       `db:"id"`
	Username    string       `db:"username"`
	PhoneNumber string       `db:"phone_number"`
	Email       string       `db:"email"`
	Password    string       `db:"password"`
	CreatedAt   sql.NullTime `db:"created_at"`
}

type RefreshToken struct {
	Id        string       `db:"id"`
	UserId    string       `db:"user_id"`
	Token     string       `db:"token"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	SaveRefreshToken(ctx context.Context, token *RefreshToken) error
	FindByUserId(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByuserIdRefreshToken(ctx context.Context, id string) error
	FindByResfreshToken(ctx context.Context, refreshToken string) (RefreshToken, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
	GetListByPhoneNumber(ctx context.Context, data []string) ([]string, error)
}

type UserService interface {
	GetByPhoneNumber(ctx context.Context, req []string) ([]string, error)
	GetUserDetail(ctx context.Context, userId string) (User, error)
}
