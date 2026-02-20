package repository

import (
	"chat-app/domain"
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: goqu.New("postgres", db),
	}
}

// GetByEmail implements [domain.UserRepository].
func (u *UserRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

// Save implements [domain.UserRepository].
func (u *UserRepository) Save(ctx context.Context, user *domain.User) error {
	dataset := u.db.Insert("users").Rows(user).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// SaveRefreshToken implements [domain.UserRepository].
func (u *UserRepository) SaveRefreshToken(ctx context.Context, token *domain.RefreshToken) error {
	dataset := u.db.Insert("refresh_tokens").Rows(token).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// FindByUserId implements [domain.UserRepository].
func (u *UserRepository) FindByUserId(ctx context.Context, id string) (result domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("id").Eq(id))
	err = dataset.ScanValsContext(ctx, &result)
	return
}

// FindByResfreshToken implements [domain.UserRepository].
func (u *UserRepository) FindByResfreshToken(ctx context.Context, refreshToken string) (result domain.RefreshToken, err error) {
	dataset := u.db.From("refresh_tokens").Where(goqu.C("token").Eq(refreshToken))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// FindByuserIdRefreshToken implements [domain.UserRepository].
func (u *UserRepository) FindByuserIdRefreshToken(ctx context.Context, id string) error {
	dataset := u.db.Delete("refresh_tokens").Where(goqu.C("user_id").Eq(id)).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// DeleteRefreshToken implements [domain.UserRepository].
func (u *UserRepository) DeleteRefreshToken(ctx context.Context, userId string) error {
	dataset := u.db.Delete("refresh_tokens").Where(goqu.C("user_id").Eq(userId)).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}

// GetByPhoneNumber implements [domain.UserRepository].
func (u *UserRepository) GetListByPhoneNumber(ctx context.Context, phones []string) (result []string, err error) {
	fmt.Println("hallo", phones)
	err = u.db.From("users").Select("phone_number").Where(goqu.C("phone_number").In(phones)).ScanValsContext(ctx, &result)
	return
}

// GetByPhoneNumber implements [domain.UserRepository].
func (u *UserRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (result domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("phone_number").Eq(phoneNumber))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}
