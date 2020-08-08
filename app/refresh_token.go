package app

// RefreshToken scheme
type RefreshToken struct {
	UserID string
	Token  string
}

// RefreshTokenService provide functionality to serve Token
type RefreshTokenService interface {
	Get(string, string) (string, error)

	Add(string, string) error
	Remove(string, string) error
	RemoveByUserID(string) error
}
