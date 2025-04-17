package controllers

import (

	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"encoding/json"
	
	"github.com/gosimple/slug"
)


// CreateSurvey godoc
// @Summary Create a new survey
// @Description Admins can create a new survey with embedded questions.
// @Tags Surveys
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param survey body models.Surveys true "Survey data"
// @Success 201 {object} models.Surveys "Survey created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys [post]
func CreateSurvey(c *fiber.Ctx) error {
	var surveyInput struct {
		SurveyName        string            `json:"survey_name"`
		SurveySlug        string            `json:"survey_slug,omitempty"` 
		SurveyDescription string            `json:"survey_description"`
		Questions         []models.Question `json:"questions"`
	}

	if err := c.BodyParser(&surveyInput); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	surveySlug := surveyInput.SurveySlug
	if surveySlug == "" {
		surveySlug = slug.Make(surveyInput.SurveyName)
	}

	var existingSurvey models.Surveys
	if err := database.Database.Db.Where("survey_slug = ?", surveySlug).First(&existingSurvey).Error; err == nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Survey with this slug already exists")
	}

	questionsJSON, err := json.Marshal(surveyInput.Questions)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to process questions")
	}

	newSurvey := models.Surveys{
		SurveyName:        surveyInput.SurveyName,
		SurveySlug:        surveySlug,
		SurveyDescription: surveyInput.SurveyDescription,
		Questions:         string(questionsJSON),
	}

	if err := database.Database.Db.Create(&newSurvey).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to create survey")
	}

	return c.Status(fiber.StatusCreated).JSON(newSurvey)
}

// UpdateSurvey godoc
// @Summary Update an existing survey
// @Description Admins can update a survey's details including questions.
// @Tags Surveys
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Survey ID"
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param survey body models.Surveys true "Updated survey data"
// @Success 200 {object} models.Surveys "Survey updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "Survey not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/{id} [put]
func UpdateSurvey(c *fiber.Ctx) error {
	id, err := parseSurveyID(c.Params("id"))
	if err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid survey ID")
	}

	var surveyInput struct {
		SurveyName        string            `json:"survey_name"`
		SurveySlug        string            `json:"survey_slug,omitempty"`
		SurveyDescription string            `json:"survey_description"`
		Questions         []models.Question `json:"questions"`
	}

	if err := c.BodyParser(&surveyInput); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	var existingSurvey models.Surveys
	if err := database.Database.Db.First(&existingSurvey, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, fiber.StatusNotFound, "Survey not found")
		}
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch survey")
	}

	existingSurvey.SurveyName = surveyInput.SurveyName
	existingSurvey.SurveySlug = slug.Make(surveyInput.SurveyName)
	if surveyInput.SurveySlug != "" {
		existingSurvey.SurveySlug = surveyInput.SurveySlug
	}
	existingSurvey.SurveyDescription = surveyInput.SurveyDescription

	questionsJSON, err := json.Marshal(surveyInput.Questions)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to process questions")
	}
	existingSurvey.Questions = string(questionsJSON)

	if err := database.Database.Db.Save(&existingSurvey).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to update survey")
	}

	return c.JSON(existingSurvey)
}

// DeleteSurvey godoc
// @Summary Delete a survey
// @Description Admins can delete a survey by ID.
// @Tags Surveys
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param id path int true "Survey ID"
// @Success 200 {object} map[string]string "Survey deleted successfully"
// @Failure 404 {object} ErrorResponse "Survey not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/{id} [delete]
func DeleteSurvey(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.Database.Db.Delete(&models.Surveys{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, fiber.StatusNotFound, "Survey not found")
		}
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete survey")
	}

	return c.JSON(fiber.Map{"message": "Survey deleted successfully"})
}

// GetAllSurveys godoc
// @Summary Fetch all surveys
// @Description Get a list of all surveys.
// @Tags Surveys
// @Accept json
// @Produce json
// @Success 200 {array} models.Surveys "List of surveys"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys [get]
func GetAllSurveys(c *fiber.Ctx) error {
	var surveys []models.Surveys
	if err := database.Database.Db.Find(&surveys).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch surveys")
	}

	return c.JSON(surveys)
}

// GetSurveyDetails godoc
// @Summary Fetch a survey with its embedded questions
// @Description Get a specific survey along with its embedded questions.
// @Tags Surveys
// @Accept json
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} models.Surveys "Survey details with questions"
// @Failure 404 {object} ErrorResponse "Survey not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/{id} [get]
func GetSurveyDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	var survey models.Surveys

	// Fetch survey from the database
	if err := database.Database.Db.First(&survey, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, fiber.StatusNotFound, "Survey not found")
		}
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch survey details")
	}

	var questions []models.Question
	if err := json.Unmarshal([]byte(survey.Questions), &questions); err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to process questions")
	}

	return c.JSON(fiber.Map{
		"id":                survey.ID,
		"survey_name":       survey.SurveyName,
		"survey_description": survey.SurveyDescription,
		"questions":         questions,
		"created_at":        survey.CreatedAt,
		"updated_at":        survey.UpdatedAt,
	})
}

// SubmitSurveyResponse godoc
// @Summary Submit a user's survey response
// @Description Users can submit their responses to a survey.
// @Tags Surveys
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param id path int true "Survey ID"
// @Param response body models.UserAnswers true "User's survey response"
// @Success 201 {object} models.UserAnswers "Response submitted successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "Survey not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/{id}/responses [post]
func SubmitSurveyResponse(c *fiber.Ctx) error {
	id, err := parseSurveyID(c.Params("id"))
	if err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid survey ID")
	}

	var survey models.Surveys
	if err := database.Database.Db.First(&survey, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, fiber.StatusNotFound, "Survey not found")
		}
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch survey")
	}

	userID, ok := c.Locals("userID").(uint)
	if !ok || userID == 0 {
		return sendErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized or invalid user")
	}

	var input struct {
		Answers map[string]interface{} `json:"answers"`
	}
	if err := c.BodyParser(&input); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	response := models.UserAnswers{
		UserID:    userID,
		SurveyID:  id,
		SurveySlug: survey.SurveySlug,
		Answers:   input.Answers,
	}

	if err := database.Database.Db.Create(&response).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to submit response")
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetAllSurveyResponses godoc
// @Summary Fetch all responses for a survey
// @Description Admins can fetch all responses for a specific survey.
// @Tags Surveys
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param id path int true "Survey ID"
// @Success 200 {array} models.UserAnswers "List of survey responses"
// @Failure 404 {object} ErrorResponse "Survey not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/{id}/responses [get]
func GetAllSurveyResponses(c *fiber.Ctx) error {
	id := c.Params("id")
	var responses []models.UserAnswers

	if err := database.Database.Db.Where("survey_id = ?", id).Find(&responses).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch responses")
	}

	return c.JSON(responses)
}

// GetUserSurveyResponse godoc
// @Summary Fetch a user's response for a survey
// @Description Get a specific user's response to a survey.
// @Tags Surveys
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param id path int true "Survey ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} models.UserAnswers "User's survey response"
// @Failure 404 {object} ErrorResponse "Response not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/{id}/responses/{user_id} [get]
func GetUserSurveyResponse(c *fiber.Ctx) error {
	surveyID := c.Params("id")
	userID := c.Params("user_id")
	var response models.UserAnswers

	if err := database.Database.Db.Where("survey_id = ? AND user_id = ?", surveyID, userID).First(&response).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, fiber.StatusNotFound, "Response not found")
		}
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch response")
	}

	return c.JSON(response)
}
// GetSurveyBySlug godoc
// @Summary Fetch a survey by slug
// @Description Get a specific survey by its slug.
// @Tags Surveys
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param slug path string true "Survey slug"
// @Success 200 {object} models.Surveys "Survey details with questions"
// @Failure 404 {object} ErrorResponse "Survey not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/surveys/slug/{slug} [get]
func GetSurveyBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Survey slug is required")
	}

	var survey models.Surveys
	if err := database.Database.Db.Where("survey_slug = ?", slug).First(&survey).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, fiber.StatusNotFound, "Survey not found")
		}
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch survey")
	}

	var questions []models.Question
	if err := json.Unmarshal([]byte(survey.Questions), &questions); err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to process survey questions")
	}

	return c.JSON(fiber.Map{
		"id":                survey.ID,
		"survey_name":       survey.SurveyName,
		"survey_slug":       survey.SurveySlug,
		"survey_description": survey.SurveyDescription,
		"questions":         questions,
		"created_at":        survey.CreatedAt,
		"updated_at":        survey.UpdatedAt,
	})
}
