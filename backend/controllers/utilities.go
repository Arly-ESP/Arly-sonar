package controllers


import (
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	
	"github.com/arly/arlyApi/models"

)


func sendErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(ErrorResponse{Error: message})
}

func validateStruct(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}

func createAuthResponse(user models.User, token string) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
		FirstSession: user.FirstSession,
	}
}


type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateResponseUser(user models.User) UserSerializer {
	return UserSerializer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func parseSurveyID(id string) (uint, error) {
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(parsedID), nil
}

type ChatRequest struct {
	Message string `json:"message" validate:"required"` 
}

type ChatResponse struct {
	Response   string    `json:"response"` 
	MessageID  uint      `json:"message_id"` 
	ChatID     uint      `json:"chat_id"`   
	Timestamp  time.Time `json:"timestamp"`
}

type ChatWithMessages struct {
	ChatID   uint             `json:"chat_id"`  
	Messages []models.Message `json:"messages"` 
}

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	FirstSession bool `json:"first_session"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
	Code        string `json:"code" validate:"required"`
}