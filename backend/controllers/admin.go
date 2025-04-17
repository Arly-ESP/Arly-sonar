package controllers

import (
	"errors"
	"github.com/arly/arlyApi/models"
	"gorm.io/gorm"

	"github.com/arly/arlyApi/database"
	"github.com/gofiber/fiber/v2"
	"github.com/arly/arlyApi/utilities"
)

// GetUsers godoc
// @Summary Get all users (Admin only)
// @Description Retrieves a list of all users. Admin access required.
// @Tags Admin
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Produce json
// @Success 200 {array} UserSerializer "List of all users"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/admin/users [get]
func GetUsers(c *fiber.Ctx) error {
	utilities.LogInfo("Admin access grantedssss")
	var users []models.User
	if err := database.Database.Db.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch users", "error": err.Error()})
	}
	utilities.LogInfo("Admin access grantedxxxxx")

	responseUsers := make([]UserSerializer, 0)
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	utilities.LogInfo("Admin access finished")


	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

// GetUser godoc
// @Summary Get user details by ID (Admin only)
// @Description Retrieves a single user's details by their ID. Admin access required.
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param id path int true "User ID"
// @Success 200 {object} UserSerializer "User details"
// @Failure 400 {object} ErrorResponse "Invalid User ID"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/admin/users/{id} [get]
func GetUserDetails(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid User ID"})
	}

	var user models.User
	if err := database.Database.Db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user details", "error": err.Error()})
	}

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

// DeleteUser godoc
// @Summary Delete a user by ID (Admin only)
// @Description Deletes a user from the database by their ID. Admin access required.
// @Tags Admin
// @Security BearerAuth
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param id path int true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 400 {object} ErrorResponse "Invalid User ID"
// @Failure 401 {object} ErrorResponse "Unauthorized access"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/admin/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid User ID"})
	}

	var user models.User
	if err := database.Database.Db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch user details", "error": err.Error()})
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete user", "error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
