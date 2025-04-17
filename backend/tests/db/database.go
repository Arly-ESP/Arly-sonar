package tests_db

import (
	"os"

	"github.com/arly/arlyApi/models"
	"github.com/arly/arlyApi/utilities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := DbConfig()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		utilities.LogFatal("Error connecting to the database", err)
	}

	utilities.LogInfo("Database connected successfully.")

	utilities.LogInfo("Starting database migrations...")
	err = db.AutoMigrate(
		&models.UserActivity{},
		&models.User{},
		&models.Chat{},
		&models.Context{},
		&models.Message{},
		&models.Subscription{},
		&models.SubscriptionType{},
		&models.Preferences{},
		&models.Gamification{},
		&models.UserPersonalityMBTI{},
		&models.UserAnswers{},
		&models.Surveys{},
		// &models.PersonalityMBTI{}, // TODO Implement this model correctly
		// &models.Chatbot{},         // TODO fix implementation
		// &models.TrainingData{},    // TODO fix implementation
	)
	if err != nil {
		utilities.LogFatal("Error during migrations", err)
	}

	utilities.LogInfo("Migrations completed successfully.")

	Database = DbInstance{Db: db}

	seedAdminUser(db)
}

func seedAdminUser(db *gorm.DB) {
	var admin models.User

	if result := db.First(&admin, 1); result.RowsAffected > 0 {
		utilities.LogInfo("Admin user already exists.")
		return
	}

	adminFirstName := os.Getenv("ADMIN_FIRSTNAME")
	adminLastName := os.Getenv("ADMIN_LASTNAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminFirstName == "" || adminLastName == "" || adminEmail == "" || adminPassword == "" {
		utilities.LogFatal("Missing admin credentials in environment variables.", nil)
	}

	hashedPassword, err := utilities.HashPassword(adminPassword)
	if err != nil {
		utilities.LogFatal("Failed to hash admin password", err)
	}

	admin = models.User{
		ID:           1,
		FirstName:    adminFirstName,
		LastName:     adminLastName,
		Email:        adminEmail,
		Password:     hashedPassword,
		Verified:     true,
		FirstSession: false,
		IsDeleted:    false,
	}

	if err := db.Create(&admin).Error; err != nil {
		utilities.LogFatal("Failed to create admin user", err)
	}

	utilities.LogInfo("Admin user created successfully.")
}
