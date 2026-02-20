package service

import (
	"chat-app/domain"
	"context"
	"fmt"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// GetByPhoneNumber implements [domain.UserService].
func (u *userService) GetByPhoneNumber(ctx context.Context, req []string) ([]string, error) {
	fmt.Println("ini nomornya", req)
	dataPhone, err := u.userRepository.GetListByPhoneNumber(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("ini data phone", dataPhone)
	return dataPhone, nil
}

// GetUserDetail implements [domain.UserService].
func (u *userService) GetUserDetail(ctx context.Context, userId string) (domain.User, error) {
	panic("unimplemented")
}
