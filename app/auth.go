package app

import (
	"medods-test/models"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// CreateTokens creates a new pair of Access and Refresh Tokens
func CreateTokens(userID string) (*models.TokenDetails, error) {
	var err error

	atLifeTime, err := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME"))
	if err != nil {
		return nil, err
	}
	rtLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME"))
	if err != nil {
		return nil, err
	}
	td := &models.TokenDetails{}
	td.AccessUUID = uuid.NewV4().String()
	td.RefreshUUID = uuid.NewV4().String()

	// Create Access Token
	atClaims := jwt.MapClaims{}
	atClaims["id"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(atLifeTime)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// Create Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(rtLifeTime)).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}
