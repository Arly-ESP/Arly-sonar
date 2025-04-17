package controllers

import (
	"fmt"
	"time"

	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/models"
	"github.com/arly/arlyApi/services"
	"github.com/arly/arlyApi/templates"
	tests_services "github.com/arly/arlyApi/tests/services"
	"github.com/arly/arlyApi/utilities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a new user, sends a verification code to their email, and returns a token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User true "User registration data"
// @Success 201 {object} AuthResponse "User registered successfully"
// @Failure 400 {object} ErrorResponse "Invalid input or user already exists"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/register [post]
func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validateStruct(user); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var existingUser models.User
	if result := database.Database.Db.Where("email = ?", user.Email).First(&existingUser); result.Error == nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "User already exists")
	}

	hashedPassword, err := utilities.HashPassword(user.Password)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}
	user.Password = hashedPassword

	code, err := utilities.CreateVerificationCode()
	utilities.LogInfo(code)
	hashedCode, err := utilities.HashPassword(code)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate verification code")
	}

	user.VerificationCode = hashedCode
	expiryTime := time.Now().Add(30 * time.Minute)
	user.VerificationCodeExpiry = &expiryTime

	if err := database.Database.Db.Create(&user).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	token, err := config.GenerateToken(user)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	emailBody, err := templates.GenerateVerificationCodeEmail(code)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate email template")
	}

	emailService, ok := c.Locals("emailService").(*services.EmailService)
	if !ok {
		emailServiceTest := c.Locals("emailService").(*tests_services.EmailService)

		err = emailServiceTest.SendEmail([]string{user.Email}, "Verify Your Account", emailBody)
		if err != nil {
			utilities.LogError("Failed to send verification email", err)
		}

		return c.Status(fiber.StatusCreated).JSON(createAuthResponse(user, token))
	}

	err = emailService.SendEmail([]string{user.Email}, "Verify Your Account", emailBody)
	if err != nil {
		utilities.LogError("Failed to send verification email", err)
	}

	return c.Status(fiber.StatusCreated).JSON(createAuthResponse(user, token))
}

// LoginUser godoc
// @Summary Log in a user
// @Description Authenticates a user by checking their email, password, and verification status.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param auth body AuthRequest true "User login data"
// @Success 200 {object} AuthResponse "User authenticated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input or user not found"
// @Failure 401 {object} ErrorResponse "Invalid credentials or user not verified"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/login [post]
func LoginUser(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validateStruct(req); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	if result := database.Database.Db.Where("email = ?", req.Email).First(&user); result.Error != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "User not found")
	}

	if !user.Verified {
		return sendErrorResponse(c, fiber.StatusUnauthorized, "User is not verified. Please verify your email.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return sendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := config.GenerateToken(user)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(createAuthResponse(user, token))
}

// VerifyUser godoc
// @Summary Verify a user's email
// @Description Verifies a user's email using the provided code and returns a token upon successful verification.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param verification body VerifyRequest true "Verification data"
// @Success 200 {object} AuthResponse "User verified successfully"
// @Failure 400 {object} ErrorResponse "Invalid or missing data"
// @Failure 401 {object} ErrorResponse "Invalid or expired verification code"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/verify [post]
func VerifyUser(c *fiber.Ctx) error {
	var req VerifyRequest
	if err := c.BodyParser(&req); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validateStruct(req); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	if result := database.Database.Db.Where("email = ?", req.Email).First(&user); result.Error != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "User not found")
	}

	if user.VerificationCodeExpiry != nil && time.Now().After(*user.VerificationCodeExpiry) {
		code, err := utilities.CreateVerificationCode()
		utilities.LogInfo(code)

		hashedCode, err := utilities.HashPassword(code)
		if err != nil {
			return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate new verification code")
		}

		user.VerificationCode = hashedCode
		expiryTime := time.Now().Add(30 * time.Minute)
		user.VerificationCodeExpiry = &expiryTime

		if err := database.Database.Db.Save(&user).Error; err != nil {
			return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to update verification code")
		}

		emailService := c.Locals("emailService").(*services.EmailService)
		body := fmt.Sprintf("<p>Your new verification code is: %s</p>", code)
		if err := emailService.SendEmail([]string{user.Email}, "New Verification Code", body); err != nil {
			utilities.LogError("Failed to send new verification email", err)
			return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to send new verification code")
		}

		return sendErrorResponse(c, fiber.StatusUnauthorized, "Verification code has expired. A new code has been sent to your email. Please try again.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.VerificationCode), []byte(req.Code)); err != nil {
		return sendErrorResponse(c, fiber.StatusUnauthorized, "Invalid verification code")
	}

	user.Verified = true
	user.VerificationCode = ""
	user.VerificationCodeExpiry = nil

	if err := database.Database.Db.Save(&user).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to update user verification status")
	}

	token, err := config.GenerateToken(user)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(createAuthResponse(user, token))
}

// RequestPasswordReset godoc
// @Summary Request a password reset
// @Description Sends a password reset code to the user's email.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param email query string true "User email"
// @Success 200 {object} map[string]string "Password reset code sent successfully"
// @Failure 400 {object} ErrorResponse "Email is required or user not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/password-reset [get]
func RequestPasswordReset(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Email is required")
	}

	var user models.User
	if result := database.Database.Db.Where("email = ?", email).First(&user); result.Error != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "User not found")
	}

	code, err := utilities.CreateVerificationCode()
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate reset code")
	}

	hashedCode, err := utilities.HashPassword(code)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash reset code")
	}

	user.VerificationCode = hashedCode
	expiryTime := time.Now().Add(30 * time.Minute)
	user.VerificationCodeExpiry = &expiryTime

	if err := database.Database.Db.Save(&user).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to save reset code")
	}

	emailService, ok := c.Locals("emailService").(*services.EmailService)
	if !ok {
		emailServiceTest := c.Locals("emailService").(*tests_services.EmailService)

		body := fmt.Sprintf("<p>Your password reset code is: <strong>%s</strong></p>", code)
		if err := emailServiceTest.SendEmail([]string{user.Email}, "Reset Your Password", body); err != nil {
			utilities.LogError("Failed to send reset email", err)
			return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to send reset email")
		}

		return c.JSON(fiber.Map{
			"message": "Password reset code sent to your email.",
		})
	}

	body := fmt.Sprintf("<p>Your password reset code is: <strong>%s</strong></p>", code)
	if err := emailService.SendEmail([]string{user.Email}, "Reset Your Password", body); err != nil {
		utilities.LogError("Failed to send reset email", err)
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to send reset email")
	}

	return c.JSON(fiber.Map{
		"message": "Password reset code sent to your email.",
	})
}

// ResetPassword godoc
// @Summary Reset a user's password
// @Description Resets a user's password using the provided reset code and new password.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param reset body ResetPasswordRequest true "Password reset data"
// @Success 200 {object} map[string]string "Password reset successfully"
// @Failure 400 {object} ErrorResponse "Invalid input or user not found"
// @Failure 401 {object} ErrorResponse "Invalid or expired reset code"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/password-reset [post]
func ResetPassword(c *fiber.Ctx) error {
	var req ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := validateStruct(req); err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	if result := database.Database.Db.Where("email = ?", req.Email).First(&user); result.Error != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "User not found")
	}

	if user.VerificationCodeExpiry != nil && time.Now().After(*user.VerificationCodeExpiry) {
		return sendErrorResponse(c, fiber.StatusUnauthorized, "Reset code has expired. Please request a new one.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.VerificationCode), []byte(req.Code)); err != nil {
		return sendErrorResponse(c, fiber.StatusUnauthorized, "Invalid reset code")
	}

	hashedPassword, err := utilities.HashPassword(req.NewPassword)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash new password")
	}

	user.Password = hashedPassword
	user.VerificationCode = ""
	user.VerificationCodeExpiry = nil

	if err := database.Database.Db.Save(&user).Error; err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed to update password")
	}

	return c.JSON(fiber.Map{
		"message": "Password reset successfully",
	})
}
