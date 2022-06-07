package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
}
