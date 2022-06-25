package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"gorm.io/gorm"
)

//represent the table structure in a table
type User struct {
	ID        string 
	Fullname  string 
	Email     string 
	Password  string 
	Active    bool  
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = pkg.HashPassword(user.Password)
	user.CreatedAt = time.Now().Local()
	return nil
}

func (user *User) BeforeUpdate(db *gorm.DB) error {
	user.UpdatedAt = time.Now().Local()
	user.Password = pkg.HashPassword(user.Password)
	return nil
}