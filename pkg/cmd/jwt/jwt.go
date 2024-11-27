package jwt

import (
	"time"

	_ "github.com/golang-jwt/jwt/v5"
)

type JWTManagerImpl struct {
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func (m *JWTManagerImpl) GenerateAccessToken(userID string) (string, error) {
	// Реализация генерации Access Token
	return "accessToken", nil
}

func (m *JWTManagerImpl) GenerateRefreshToken(userID string) (string, error) {
	// Реализация генерации Refresh Token
	return "refreshToken", nil
}
