package models

// TokenDetails represent scheme of token
type TokenDetails struct {
	// TODO: probably access_token should not be stored in a DB
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
}
