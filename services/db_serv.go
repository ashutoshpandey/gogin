package services

import (
	"gorm.io/gorm"

	"github.com/ashutoshpandey/gogin/config"
)

// Initialization logic
// --------------------------------------------------------

// DbService provides user-related operations
type DbService struct {
	DB *gorm.DB
}

// NewDbService creates a new UserService
func NewDBService() *DbService {
	dbConfig := config.LoadDbConfig()
	print(dbConfig)
	/*
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbConfig.HOST, dbConfig.PORT, dbConfig.USER, dbConfig.DATABASE, dbConfig.PASSWORD)

		// Open the database connection
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Migrate the schema
		db.AutoMigrate(&models.User{})

		return &DbService{DB: db}
	*/
	return nil
}
