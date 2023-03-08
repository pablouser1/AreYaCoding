package tokens

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func extractToken(auth string) string {
	if len(strings.Split(auth, " ")) == 2 {
		return strings.Split(auth, " ")[1]
	}
	return ""
}

func Verify(auth string) (string, error) {
	tokenString := extractToken(auth)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("APP_TOKEN")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["user"].(string)
		return username, nil
	}
	return "", nil
}
