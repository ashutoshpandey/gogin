package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model in the database
type User struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100)"`
	Password    string    `gorm:"type:varchar(512)"`
	Email       string    `gorm:"type:varchar(100);unique_index"`
	DateOfBirth time.Time `gorm:"type:date"`
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}
