package entity

import "time"

type User struct {
	UserId    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	IsAdmin   bool      `json:"is_admin"`
}
