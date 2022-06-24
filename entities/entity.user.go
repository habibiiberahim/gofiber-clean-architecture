package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"gorm.io/gorm"
)

//represent the table structure in a table
type EntityUser struct {
	ID        string `gorm:"primaryKey;"`
	Fullname  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityUser) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = pkg.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUser) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	entity.Password = pkg.HashPassword(entity.Password)
	return nil
}