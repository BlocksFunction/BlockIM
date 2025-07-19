package model

import (
	"time"
)

type User struct {
	UserID    int64      `gorm:"primaryKey;column:user_id"`
	Username  string     `gorm:"unique;not null;size:64;column:username"`
	Email     string     `gorm:"unique;not null;size:255;column:email"`
	Password  string     `gorm:"not null;column:password"`
	IsBanned  bool       `gorm:"not null;default:false;column:is_banned"`
	IsActive  bool       `gorm:"default:false"`
	IsAdmin   bool       `gorm:"not null;default:false;column:is_admin"`
	CreatedAt time.Time  `gorm:"autoCreateTime;column:created_at"`
	LastLogin *time.Time `gorm:"autoUpdateTime;column:last_login"`
}
