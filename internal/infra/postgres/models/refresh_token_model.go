package models

import (
	"time"

	"github.com/google/uuid"
)






type RefreshTokenModel struct{
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	TokenHash string `gorm:"not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	RevokedAt *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
}


func (RefreshTokenModel)TableName() string {return  "refresh_tokens"}