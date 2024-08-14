package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserJwt struct {
	jwt.RegisteredClaims
	UserId int `json:"userId"`
}

const (
	DefaultExpired = time.Hour * 1 // Default expiration time is 1 hour
)

func GenerateJWT(secretKey string, expiredTime int64, userId int) (string, error) {
	exp := DefaultExpired
	if expiredTime != 0 {
		exp = time.Hour * time.Duration(expiredTime)
	}

	claims := UserJwt{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)), // Expiration time
			IssuedAt:  jwt.NewNumericDate(time.Now()),          // Issue time
			NotBefore: jwt.NewNumericDate(time.Now()),          // Effective time
		},
	}

	// Use HS256 signature algorithm
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secretKey))
}

func ParseJWT(secretKey, token string) (*UserJwt, error) {
	t, err := jwt.ParseWithClaims(token, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*UserJwt); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
