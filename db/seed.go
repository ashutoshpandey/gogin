package db

import (
	"fmt"
	"log"
	"time"

	"github.com/ashutoshpandey/gogin/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUsers seeds the database with initial user data
func SeedUsers(db *gorm.DB) {
	var userCount int64
	// Count the number of existing users
	if err := db.Model(&models.User{}).Count(&userCount).Error; err != nil {
		log.Fatalf("Failed to count users: %v", err)
	}

	// If users already exist, skip seeding
	if userCount > 0 {
		fmt.Println("Users already exist, skipping seeding.")
		return
	}

	fmt.Println("Running seed data")

	// Example seed data
	users := []models.User{
		{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Password:    hashPassword("password123"), // Ensure you hash passwords before storing them
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:        "Jane Smith",
			Email:       "janesmith@example.com",
			Password:    hashPassword("password123"), // Ensure you hash passwords before storing them
			DateOfBirth: time.Date(1985, time.May, 12, 0, 0, 0, 0, time.UTC),
		},
	}

	// Insert seed data
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed users: %v", err)
		}
	}

	log.Println("Successfully seeded users.")
}

// hashPassword hashes a plaintext password using bcrypt
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashedPassword)
}
