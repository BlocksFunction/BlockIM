package model

import (
	"time"
)

type VerificationToken struct {
	Token     string `gorm:"uniqueIndex;not null"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"constraint:OnDelete:CASCADE;"`
	ExpiresAt time.Time
}
