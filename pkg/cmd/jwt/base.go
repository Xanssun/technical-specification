package jwt

import (
	"time"
)

type JWT struct {
	JWTManager interface {
		GenerateAccessToken(userID string) (string, error)
		GenerateRefreshToken(userID string) (string, error)
	}
}

func NewJWTManager(secretKey string, accessTokenTTL, refreshTokenTTL time.Duration) JWT {
	return JWT{
		JWTManager: &JWTManagerImpl{
			secretKey:       secretKey,
			accessTokenTTL:  accessTokenTTL,
			refreshTokenTTL: refreshTokenTTL,
		},
	}
}
