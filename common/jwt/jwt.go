package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	defaultExpired = 1 // The unit is hours
)

func GenerateJWT(secretKey string, expiredTime int64, userId int) (string, error) {
	if expiredTime == 0 {
		expiredTime = defaultExpired
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix()
	claims["userId"] = userId

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseJWT(secretKey, tokenString string) (userId int, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unsupported signature method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["userId"].(int)
		if !ok {
			return 0, fmt.Errorf("incorrect userId value: %v", claims["userId"])
		}
		return userId, nil
	}
	return 0, fmt.Errorf("Invalid token")
}
