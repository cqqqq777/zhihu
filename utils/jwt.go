package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
	g "zhihu/global"
)

type MyClaims struct {
	UserID int
	jwt.RegisteredClaims
}

func GenToken(uid int) (string, error) {
	claim := MyClaims{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Auth.Jwt.ExpiresTime) * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(g.Config.Auth.Jwt.SecretKey))
}
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.Config.Auth.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
