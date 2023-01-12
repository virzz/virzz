package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/virzz/virzz/common"
)

// Claims -
type Claims struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetHeaderToken(authorization string) (token string, err error) {
	tokens := strings.Split(authorization, " ")
	if len(tokens) != 2 {
		return "", fmt.Errorf("authorization format error")
	}
	if tokens[0] != "Bearer" {
		return "", fmt.Errorf("lack of Bearer")
	}
	return tokens[1], nil
}

// GenerateToken generate tokens used for auth
func GenerateToken(token, username string) (string, error) {
	claims := Claims{
		token,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "webkit",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(common.GetConfig().Jwt.Secret))
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.GetConfig().Jwt.Secret), nil
	})
	if err == nil && tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// RefreshToken -
func RefreshToken(tokenString string) (string, error) {
	if claims, err := ParseToken(tokenString); err == nil {
		claims.ExpiresAt.Add(7 * 24 * time.Hour)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(common.GetConfig().Jwt.Secret))
	}
	return "", fmt.Errorf("couldn't handle this token")
}
