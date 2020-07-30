package app

// RefreshTokenDetail scheme
type RefreshTokenDetail struct {
	UUID       string
	AccessUUID string
	Token      string
}

// CreateRefreshTokenDetail from access and refresh token pair.
// Bind refresh_token with access_token
func CreateRefreshTokenDetail(m map[string]string) *RefreshTokenDetail {
	var rtDetail *RefreshTokenDetail
	rtDetail.Token = m["refresh_token"]
	rtDetail.UUID = m["refresh_uuid"]
	rtDetail.AccessUUID = m["access_uuid"]
	return rtDetail
}

// TokenService provide functionality to serve Token
type TokenService interface {
	Get(uuid string) (*RefreshTokenDetail, error)
	GetByAccessUUID(uuid string) (*RefreshTokenDetail, error)

	Create(*RefreshTokenDetail) error
	Delete(refreshUUID string) error
}
