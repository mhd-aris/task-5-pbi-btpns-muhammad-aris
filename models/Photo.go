package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title  string    `gorm:"size:255;not null;unique" json:"title"`
	Caption     string    `gorm:"not null;unique" json:"caption"`
	PhotoUrl  string    `gorm:"not null;" json:"photo_url"`
	UserID  uuid.UUID `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) error{
	uuid := uuid.New()
	photo.ID = uuid
	return nil
}