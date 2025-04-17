package controllers

import (
	"errors"
	"github.com/arly/arlyApi/models"
	"gorm.io/gorm"

	"github.com/arly/arlyApi/database"

	"github.com/gofiber/fiber/v2"
	"time"
	"github.com/arly/arlyApi/enums"
)



// LogUserMood godoc
// @Summary Log a user's mood
// @Description Logs the mood of a user for a given day. If a mood already exists for the day, it updates the mood.
// @Security BearerAuth
// @Tags Mood
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param activity body models.UserActivity true "User activity information"
// @Success 200 {object} models.UserActivity "Mood updated successfully"
// @Success 201 {object} models.UserActivity "Mood logged successfully"
// @Failure 400 {object} ErrorResponse "Invalid input or missing required fields"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/mood [post]
func LogUserMood(c *fiber.Ctx) error {
	var activity models.UserActivity

	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	if activity.Mood == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Mood is required"})
	}

	validMoods := map[enums.MoodEnum]bool{
		enums.Happy:   true,
		enums.Nice:     true,
		enums.Poker: true,
		enums.Sad:   true,
		enums.Bad: true,
	}
	if _, isValid := validMoods[enums.MoodEnum(activity.Mood)]; !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid mood value"})
	}

	userID := c.Locals("userID").(uint)
	if activity.Date.IsZero() {
		activity.Date = time.Now().Truncate(24 * time.Hour)
	}

	var existingActivity models.UserActivity
	if err := database.Database.Db.Where("user_id = ? AND date = ?", userID, activity.Date).First(&existingActivity).Error; err == nil {
		existingActivity.Mood = activity.Mood
		if saveErr := database.Database.Db.Save(&existingActivity).Error; saveErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update activity", "error": saveErr.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(existingActivity)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch activity", "error": err.Error()})
	}

	newActivity := models.UserActivity{
		UserID:      userID,
		Mood:        activity.Mood,
		Date:        activity.Date,
		MessageCount: 0,
	}
	if createErr := database.Database.Db.Create(&newActivity).Error; createErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to log activity", "error": createErr.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newActivity)
}

// DeleteUser godoc
// @Summary Delete the authenticated user's account
// @Description Deletes the currently authenticated user's account.
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Success 200 {object} fiber.Map "User deleted successfully"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user [delete]
func DeleteCurrentUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var user models.User
	if err := database.Database.Db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user details", "error": err.Error()})
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete user", "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

// GetUserActivity godoc
// @Summary Get the activity of the authenticated user for a specific date
// @Description Fetches the activity of the authenticated user for the specified date. Defaults to today's activity if no date is provided.
// @Tags Activity
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param date query string false "Date in YYYY-MM-DD format (defaults to today)"
// @Success 200 {object} models.UserActivity "User's activity for the day"
// @Failure 400 {object} ErrorResponse "Invalid date format"
// @Failure 404 {object} ErrorResponse "No activity found for the user on the specified date"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user/activity [get]
func GetUserActivity(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	date := c.Query("date")
	var queryDate time.Time
	var err error

	if date == "" {
		queryDate = time.Now().Truncate(24 * time.Hour)
	} else {
		queryDate, err = time.Parse("2024-01-02", date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format. Use YYYY-MM-DD"})
		}
	}

	var userActivity models.UserActivity
	if err := database.Database.Db.Where("user_id = ? AND date = ?", userID, queryDate).First(&userActivity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No activity found for the user on the specified date"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user activity", "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(userActivity)
}

// GetUserActivities godoc
// @Summary Get the activity history of the authenticated user
// @Description Fetches all activity logs of the authenticated user, optionally filtered by a date range.
// @Tags Activity
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param start_date query string false "Start date in YYYY-MM-DD format"
// @Param end_date query string false "End date in YYYY-MM-DD format"
// @Success 200 {array} models.UserActivity "List of user activities"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "No activity found for the user"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user/activities [get]
func GetUserActivities(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateParam != "" {
		startDate, err = time.Parse("2024-01-02", startDateParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid start_date format. Use YYYY-MM-DD"})
		}
	}
	if endDateParam != "" {
		endDate, err = time.Parse("2024-01-02", endDateParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid end_date format. Use YYYY-MM-DD"})
		}
	}

	query := database.Database.Db.Where("user_id = ?", userID)
	if !startDate.IsZero() {
		query = query.Where("date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("date <= ?", endDate)
	}

	var userActivities []models.UserActivity
	if err := query.Find(&userActivities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No activity found for the user"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user activities", "error": err.Error()})
	}

	if len(userActivities) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No activity found for the user"})
	}

	return c.Status(fiber.StatusOK).JSON(userActivities)
}


// GetUser godoc
// @Summary Get the authenticated user's details
// @Description Retrieves the details of the currently authenticated user.
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Success 200 {object} UserSerializer "User found"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user [get]
func GetUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var user models.User
	if err := database.Database.Db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user details", "error": err.Error()})
	}

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

// UpdateUser godoc
// @Summary Update the authenticated user's details
// @Description Updates the details of the currently authenticated user.
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param user body models.User true "Updated user data"
// @Success 200 {object} UserSerializer "User updated"
// @Failure 400 {object} ErrorResponse "Invalid input or missing required fields"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user [put]
func UpdateUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var user models.User
	if err := database.Database.Db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user details", "error": err.Error()})
	}

	var updatedUserData models.User
	if err := c.BodyParser(&updatedUserData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	user.FirstName = updatedUserData.FirstName
	user.LastName = updatedUserData.LastName
	user.Email = updatedUserData.Email

	if err := database.Database.Db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update user details", "error": err.Error()})
	}

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}
