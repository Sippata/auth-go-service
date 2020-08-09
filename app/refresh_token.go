package app

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// RefreshToken scheme
type RefreshToken struct {
	ID        string
	UserID    string
	TokenHash string
}

// RefreshTokenService provide functionality to serve Token
type RefreshTokenService interface {
	Get(uuid string) (string, error)

	Add(*jwt.Token) error
	Remove(uuid string) error
	RemoveByUserID(uuid string) error
}

// Compare hash and token
func Compare(hash string, token string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
}

// GenerateHash for token
func GenerateHash(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), 10)
	return string(hash), err
}
