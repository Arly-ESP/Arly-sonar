package utilities

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"math/big"
	"strings"
	"errors"


)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CreateVerificationCode() (string, error) {
	code, err := GenerateRandomString(6, "numeric")
	if err != nil {
		return "", err
	}
	return code, nil
}

func GenerateRandomString(length int, genType string) (string, error) {
	var letters string
	switch genType {
	case "alphanumeric":
		letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case "numeric":
		letters = "0123456789"
	case "alphabetic":
		letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	default:
		return "", errors.New("invalid generation type, must be 'alphanumeric', 'numeric', or 'alphabetic'")
	}

	var result strings.Builder
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result.WriteByte(letters[randomIndex.Int64()])
	}
	// Pour affi cher le code de vérification
	// fmt.Println("Votre code de vérification est: ", result.String())
	return result.String(), nil
}