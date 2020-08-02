package app

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// CreateTokens creates a new TokenPair
func CreateTokens(userID string) (map[string]string, error) {
	atLifeTime, err := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME"))
	if err != nil {
		return nil, err
	}
	rtLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME"))
	if err != nil {
		return nil, err
	}

	accessID := uuid.New().String()
	exp := time.Now().Add(time.Minute * time.Duration(atLifeTime)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Id:        accessID,
		Subject:   userID,
		ExpiresAt: exp,
	})
	accessToken, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	exp = time.Now().Add(time.Minute * time.Duration(rtLifeTime)).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Id:        uuid.New().String(),
		Subject:   userID,
		ExpiresAt: exp,
		Audience:  accessID,
	})
	refreshToken, err := rt.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// ParseToken from string
func ParseToken(tokenString string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
