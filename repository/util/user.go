package util

import (
	"github.com/rudderstack_source_app/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    int64              `bson:"user_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	IsAdmin   bool               `bson:"is_admin"`
}

func UserDaoToEntity(u User) *entity.User {
	return &entity.User{
		UserId:    u.UserId,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		IsAdmin:   u.IsAdmin,
	}
}

func UserEntityToDao(u *entity.User) *User {
	return &User{
		UserId:    u.UserId,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		IsAdmin:   u.IsAdmin,
	}
}
