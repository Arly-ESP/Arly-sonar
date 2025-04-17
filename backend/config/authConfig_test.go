package config

import (
	"testing"

	"github.com/arly/arlyApi/models"
)

func TestInitializeJWT(t *testing.T) {
	validSecret := "abcdefghijklmnopqrstuvwxyz123456" // 32 characters
	t.Setenv("JWT_SECRET", validSecret)
	if err := InitializeJWT(); err != nil {
		t.Errorf("Expected no error for valid secret, got: %v", err)
	}

	t.Setenv("JWT_SECRET", "")
	if err := InitializeJWT(); err == nil {
		t.Error("Expected error for empty secret, got nil")
	}

	t.Setenv("JWT_SECRET", "shortsecret")
	if err := InitializeJWT(); err == nil {
		t.Error("Expected error for short secret, got nil")
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	validSecret := "abcdefghijklmnopqrstuvwxyz123456" // 32 characters
	t.Setenv("JWT_SECRET", validSecret)
	if err := InitializeJWT(); err != nil {
		t.Fatalf("InitializeJWT failed: %v", err)
	}

	user := models.User{
		ID:    1,
		Email: "test@example.com",
	}

	tokenStr, err := GenerateToken(user)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	parsedUserID, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if parsedUserID != user.ID {
		t.Errorf("Expected userID %d, got %d", user.ID, parsedUserID)
	}
}

func TestIsValidEmail(t *testing.T) {
	validEmail := "test@example.com"
	if err := IsValidEmail(validEmail); err != nil {
		t.Errorf("Expected valid email, got error: %v", err)
	}

	invalidEmail := "not-an-email"
	if err := IsValidEmail(invalidEmail); err == nil {
		t.Error("Expected error for invalid email, got nil")
	}
}
