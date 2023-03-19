package user

import (
	"context"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/repository/user"
)

type UserService struct {
	userRepo user.UserRepo
}

func NewService(userRepo user.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateOrUpdateUser(ctx context.Context, user entity.User) error {
	return s.userRepo.CreateOrUpdateUser(ctx, &user)
}

func (s *UserService) GetUserById(ctx context.Context, userId int64) (*entity.User, error) {
	userEntity, err := s.userRepo.GetUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}
