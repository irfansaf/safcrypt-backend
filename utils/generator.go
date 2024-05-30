package utils

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// EncryptPassword Encrypt password using bcrypt
func EncryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ComparePassword Compare hashed password with plain password
func ComparePassword(hashedPassword []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}

// GenerateJWTToken Generate JWT token using user ID and secret key
func GenerateJWTToken(userID uuid.UUID, secretKey string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}
