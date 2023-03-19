package user

import (
	"context"
	"github.com/rudderstack_source_app/entity"
)

type UserRepo interface {
	GetUserByUserId(ctx context.Context, userId int64) (*entity.User, error)
	CreateOrUpdateUser(ctx context.Context, user *entity.User) error
}
