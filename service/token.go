package service

import (
	"fanatic/lib/token"
)

type TokenService interface {
	RefreshExchangeAccessToken(refreshToken string) (string, error)
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (t *tokenService) RefreshExchangeAccessToken(refreshToken string) (string, error) {
	claims, err := token.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	uid := claims["identity"].(float64)
	return token.CreateToken(token.ACCESS, int32(uid)), nil
}
