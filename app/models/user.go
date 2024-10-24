package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Provider string

const (
	Github Provider = "github"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" `
	Email      string    `json:"email" gorm:"unique"`
	Username   string    `json:"username" gorm:"unique"`
	Provider   Provider  `json:"provider" gorm:"not null"`
	ProviderID string    `json:"provider_id" gorm:"unique;not null"`
	AvatarURL  string    `json:"avatar_url" gorm:"not null"`
	CreatedAt  string    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  string    `json:"updated_at" gorm:"autoUpdateTime"`
}
