package utilities

import (
	"regexp"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}
	if hashed == password {
		t.Errorf("Hashed password should not equal the original password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		t.Errorf("Hash does not match original password: %v", err)
	}
}

func TestCreateVerificationCode(t *testing.T) {
	code, err := CreateVerificationCode()
	if err != nil {
		t.Fatalf("CreateVerificationCode returned error: %v", err)
	}
	if len(code) != 6 {
		t.Errorf("Expected verification code length 6, got %d", len(code))
	}
	matched, err := regexp.MatchString(`^\d{6}$`, code)
	if err != nil {
		t.Fatalf("Error compiling regex: %v", err)
	}
	if !matched {
		t.Errorf("Verification code is not numeric: %s", code)
	}
}

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name    string
		genType string
		allowed string
	}{
		{"alphanumeric", "alphanumeric", "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"},
		{"numeric", "numeric", "0123456789"},
		{"alphabetic", "alphabetic", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"},
	}

	length := 10
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := GenerateRandomString(length, tt.genType)
			if err != nil {
				t.Errorf("GenerateRandomString(%q) returned error: %v", tt.genType, err)
			}
			if len(s) != length {
				t.Errorf("Expected length %d for %q, got %d", length, tt.genType, len(s))
			}
			for _, ch := range s {
				if !strings.Contains(tt.allowed, string(ch)) {
					t.Errorf("Character %q in output for %q is not allowed", ch, tt.genType)
				}
			}
		})
	}

	_, err := GenerateRandomString(length, "invalid")
	if err == nil {
		t.Errorf("Expected error for invalid generation type, got nil")
	}
}

func TestGenerateRandomStringZeroLength(t *testing.T) {
	s, err := GenerateRandomString(0, "numeric")
	if err != nil {
		t.Errorf("GenerateRandomString(0, numeric) returned error: %v", err)
	}
	if s != "" {
		t.Errorf("Expected empty string for zero length, got %q", s)
	}
}
