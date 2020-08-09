package app

import "golang.org/x/crypto/bcrypt"

// RefreshToken scheme
type RefreshToken struct {
	UserID    string
	TokenHash string
}

// RefreshTokenService provide functionality to serve Token
type RefreshTokenService interface {
	Get(string, string) (string, error)

	Add(string, string) error
	Remove(string, string) error
	RemoveByUserID(string) error
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
