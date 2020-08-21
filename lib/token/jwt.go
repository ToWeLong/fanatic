package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

const (
	ACCESS  = "access"
	REFRESH = "refresh"
)

type FanaticClaim struct {
	Identity  int32  `json:"identity"`
	Scope     string `json:"scope"`
	TokenType string `json:"type"`
	jwt.StandardClaims
}

func CreateToken(tokenType string, id int32) string {
	switch tokenType {
	case ACCESS:
		accessTime := time.Duration(viper.GetInt("token.access_expire_time"))
		expireTime := time.Now().Add(time.Hour * accessTime)
		return GenerateToken(ACCESS, id, expireTime)
	case REFRESH:
		refreshTime := time.Duration(viper.GetInt("token.refresh_expire_time"))
		expireTime := time.Now().Add(time.Hour * refreshTime)
		return GenerateToken(REFRESH, id, expireTime)
	}
	return ""
}

func VerifyAccessToken(token string) (jwt.MapClaims, error) {
	return VerifyToken(token, ACCESS)
}

func VerifyRefreshToken(token string) (jwt.MapClaims, error) {
	return VerifyToken(token, REFRESH)
}

func GenerateToken(tokenType string, id int32, exp time.Time) string {
	myClaims := FanaticClaim{
		id,
		tokenType,
		"welong",
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	secretKey := viper.GetString("token.secretKey")
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func VerifyToken(tokenStr, tokenType string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		claims := t.Claims.(jwt.MapClaims)
		jwtType := claims["scope"].(string)
		if jwtType != tokenType {
			return nil, errors.New("令牌类型错误")
		}
		secretKey := viper.GetString("token.secretKey")
		return []byte(secretKey), nil
	})
	if token != nil && token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("令牌不合法")
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet|jwt.ValidationErrorExpired) != 0 {
			return nil, errors.New("令牌过期")
		} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			return nil, errors.New("令牌签名不合法")
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
