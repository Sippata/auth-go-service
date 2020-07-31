package app

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// CreateTokens creates a new pair of Access and Refresh Tokens
func CreateTokens(userID string) (map[string]string, error) {
	atLifeTime, err := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME"))
	if err != nil {
		return nil, err
	}
	rtLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME"))
	if err != nil {
		return nil, err
	}

	// Create Access Token
	accessToken, atUUID, err := createToken(userID, os.Getenv("ACCESS_SECRET"), atLifeTime)
	if err != nil {
		return nil, err
	}

	// Create Refresh Token
	refreshToken, rtUUID, err := createToken(userID, os.Getenv("REFRES_SECRET"), rtLifeTime)
	if err != nil {
		return nil, err
	}

	tokensInfo := make(map[string]string, 5)
	tokensInfo["access_token"] = accessToken
	tokensInfo["access_uuid"] = atUUID
	tokensInfo["refresh_token"] = refreshToken
	tokensInfo["refresh_uuid"] = rtUUID
	tokensInfo["user_id"] = userID

	return tokensInfo, nil
}

func createToken(userID string, secret string, lifetime int) (string, string, error) {
	uuid := uuid.New().String()
	claims := jwt.MapClaims{}
	claims["id"] = uuid
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(lifetime)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := t.SignedString([]byte(secret))
	return token, uuid, err
}
