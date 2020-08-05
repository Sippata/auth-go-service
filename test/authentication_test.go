package test

import (
	"testing"
)



func TestUserCanSuccessfullyLogin(t *testing.T) {
	
}
func TestUserGets_403_OnInvalidCredentials(t *testing.T) {}
func TestUserResives_401_OnExpiredToken(t *testing.T) {}
func TestUserCanRefreshAccessToken(t *testing.T) {}
func TestUserCanUseRefreshTokenOnlyOnes(t *testing.T) {}
func TestRefreshTokenBecomeInvalidOnLogout(t *testing.T) {}
func TestMultipleRefreshTokensAreValid(t *testing.T) {}
