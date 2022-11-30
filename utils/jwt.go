package utils

import (
	"errors"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mozhu1024/virzz/common"
)

var jwtSecret = &common.Conf.Jwt.Secret

// Claims -
type Claims struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetHaderAuthorizationToken(authorization string) (token string, err error) {
	var (
		bearer string
		match  [][]string
	)
	match = regexp.MustCompile(`(?m)(Bearer)\s+([a-zA-Z0-9\_\-\.\=]*)`).
		FindAllStringSubmatch(authorization, -1)
	if len(match) > 0 && len(match[0]) > 2 {
		bearer = match[0][1]
		token = match[0][2]
		if bearer != "Bearer" {
			return "", errors.New("lack of Bearer")
		}
		return token, nil
	}
	return "", errors.New("authorization format error")
}

// GenerateToken generate tokens used for auth
func GenerateToken(token, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	claims := Claims{
		token,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "webkit",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(*jwtSecret))
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(*jwtSecret), nil
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
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(*jwtSecret))
	}
	return "", errors.New("couldn't handle this token")
}
