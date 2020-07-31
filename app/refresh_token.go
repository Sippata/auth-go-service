package app

// RefreshTokenDetail scheme
type RefreshTokenDetail struct {
	UUID       string
	AccessUUID string
	UserID     string
	Token      string
}

// CreateRefreshTokenDetail from access and refresh token pair.
// Bind refresh_token with access_token
func CreateRefreshTokenDetail(m map[string]string) *RefreshTokenDetail {
	var rtDetail *RefreshTokenDetail
	rtDetail.Token = m["refresh_token"]
	rtDetail.UUID = m["refresh_uuid"]
	rtDetail.AccessUUID = m["access_uuid"]
	rtDetail.UserID = m["user_id"]
	return rtDetail
}

// TokenService provide functionality to serve Token
type TokenService interface {
	Get(uuid string) (*RefreshTokenDetail, error)
	GetByAccessUUID(uuid string) (*RefreshTokenDetail, error)

	Add(*RefreshTokenDetail) error
	Remove(refreshUUID string) error
	RemoveByUserID(userID string) error
}
