package user

import (
	"context"
	"github.com/rudderstack_source_app/entity"
)

type UserServiceInterface interface {
	CreateOrUpdateUser(ctx context.Context, user entity.User) error
	GetUserById(ctx context.Context, userId int64) (*entity.User, error)
}
