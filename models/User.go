package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string    `gorm:"size:255;not null;unique" json:"Username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;" json:"password"`
	Photos  []Photo   `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photos"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error{
	uuid := uuid.New()
	user.ID = uuid
	return nil
}