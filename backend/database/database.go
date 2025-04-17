package database

import (
	"os"
	"encoding/json"

	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/models"
	"github.com/arly/arlyApi/utilities"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := config.DbConfig()
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
	seedSurvey(db)
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


func seedSurvey(db *gorm.DB) {
	var survey models.Surveys
	if result := db.Where("survey_slug = ?", "onboarding-personnalite-preferences").First(&survey); result.RowsAffected > 0 {
		utilities.LogInfo("Survey already exists.")
		return
	}

	questions := []models.Question{
		{
			Question:        "Quel est ton prénom, ton sexe (homme, femme, autre) et ton âge ?",
			QuestionType:    "text",
			QuestionOptions: map[string]interface{}{},
			Order:           1,
		},
		{
			Question:     "Parmi les traits suivants, lequel te correspond le plus ?",
			QuestionType: "radio",
			QuestionOptions: map[string]interface{}{
				"options": []string{"Rebelle", "Organisé(e)", "Calme", "Aventureux(se)", "Créatif(ve)"},
			},
			Order: 2,
		},
		{
			Question:     "Aimes-tu prendre des décisions rapidement ?",
			QuestionType: "radio",
			QuestionOptions: map[string]interface{}{
				"options": []string{"Oui", "Non"},
			},
			Order: 3,
		},
		{
			Question:     "Es-tu à l’aise avec les imprévus ?",
			QuestionType: "radio",
			QuestionOptions: map[string]interface{}{
				"options": []string{"Oui", "Non"},
			},
			Order: 4,
		},
		{
			Question:     "Préfères-tu travailler seul(e) plutôt qu’en équipe ?",
			QuestionType: "radio",
			QuestionOptions: map[string]interface{}{
				"options": []string{"Oui", "Non"},
			},
			Order: 5,
		},
		{
			Question:     "Es-tu du genre à finir ce que tu as commencé, même si c’est difficile ?",
			QuestionType: "radio",
			QuestionOptions: map[string]interface{}{
				"options": []string{"Oui", "Non"},
			},
			Order: 6,
		},
		{
			Question:     "Prends-tu le temps de te déconnecter de ton téléphone chaque jour ?",
			QuestionType: "radio",
			QuestionOptions: map[string]interface{}{
				"options": []string{"Oui", "Non"},
			},
			Order: 7,
		},
		{
			Question:     "Quels sont tes centres d’intérêt parmi les suivants ?",
			QuestionType: "checkbox",
			QuestionOptions: map[string]interface{}{
				"options": []string{
					"Technologie", "Nature", "Sport", "Lecture", "Cinéma", "Musique",
					"Cuisine", "Voyage", "Gaming", "Mode", "Art", "Sciences",
					"Philosophie", "Activisme", "Méditation",
				},
			},
			Order: 8,
		},
	}

	questionsJSON, err := json.Marshal(questions)
	if err != nil {
		utilities.LogFatal("Failed to marshal survey questions", err)
	}

	survey = models.Surveys{
		SurveyName:        "Onboarding : Personnalité et Préférences",
		SurveySlug:        "onboarding-personnalite-preferences",
		SurveyDescription: "Un questionnaire pour comprendre votre personnalité, vos préférences et vos centres d'intérêt.",
		Questions:         string(questionsJSON),
	}

	if err := db.Create(&survey).Error; err != nil {
		utilities.LogFatal("Failed to create survey", err)
	}

	utilities.LogInfo("Survey created successfully.")
}
